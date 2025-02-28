package live

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/gorilla/websocket"
	"github.com/iyear/biligo-live/core"
	"github.com/iyear/biligo-live/message"
)

type Transport struct {
	Msg   message.Msg
	Error error
}

type Live struct {
	ws      *websocket.Conn
	debug   bool
	logger  *log.Logger
	entered chan struct{}
	hb      time.Duration
	recover func(error)
	Rev     chan *Transport
}

// NewLive 创建一个新的直播连接
func NewLive(debug bool, heartbeat time.Duration, cache int, recover func(error)) *Live {
	return &Live{
		debug:   debug,
		logger:  log.New(os.Stdout, "Live ", log.LstdFlags|log.Lshortfile),
		hb:      heartbeat,
		entered: make(chan struct{}),
		recover: recover,
		Rev:     make(chan *Transport, cache),
	}
}

// Conn ws连接bilibili弹幕服务器
func (l *Live) Conn(dialer *websocket.Dialer, host string) error {
	return l.ConnWithHeader(dialer, host, nil)
}

// ConnWithHeader ws连接bilibili弹幕服务器 (带header)
func (l *Live) ConnWithHeader(dialer *websocket.Dialer, host string, header http.Header) (err error) {
	w, _, err := dialer.Dial(host, header)
	if err != nil {
		return err
	}
	l.ws = w
	return
}

// TODO 错误重试机制

// Enter 进入房间。 Conn 后五秒内必须进入房间，否则服务器主动断开连接
func (l *Live) Enter(ctx context.Context, uid, room int, buvid, key string) error {
	enter := core.VerifyData{
		UID:      uid,
		RoomID:   room,
		Protover: 3,
		Buvid:    buvid,
		Platform: "web",
		Type:     2,
		Key:      key,
	}
	body, err := json.Marshal(enter)
	if err != nil {
		return err
	}
	if err = l.ws.WriteMessage(websocket.BinaryMessage, core.Encode(core.WsVerPlain, core.WsOpEnterRoom, body)); err != nil {
		return err
	}

	hbCtx, hbCancel := context.WithCancel(ctx)
	revCtx, revCancel := context.WithCancel(ctx)
	ifError := make(chan error, 1)
	go l.revWithError(revCtx, ifError)

	go func() {
		select {
		case <-l.entered:
		case <-hbCtx.Done():
			return
		}
		l.heartbeat(hbCtx, l.hb)
	}()

	defer func() {
		hbCancel()
		revCancel()
		err = l.ws.Close()
	}()

	select {
	// 外部停止ws
	case <-ctx.Done():
		l.info("websocket conn stopped")
		break
	// 內部接收 Websocket 訊息錯誤
	case err = <-ifError:
		l.error("websocket conn stopped with an error: %s", err)
		break
	}

	return err
}
func (l *Live) report() {
	if r := recover(); r != nil {
		var e error
		switch r := r.(type) {
		case error:
			e = r
		case string:
			e = fmt.Errorf("%s", r)
		}

		if l.recover != nil {
			l.recover(e)
		}
		l.error("panic: %s", e)
	}
}

func (l *Live) heartbeat(ctx context.Context, t time.Duration) {
	hb := func(live *Live) {
		err := live.ws.WriteMessage(websocket.BinaryMessage, core.Encode(core.WsVerPlain, core.WsOpHeartbeat, nil))
		if err != nil {
			live.push(ctx, nil, fmt.Errorf("failed to send hearbeat: %s", err))
		}
	}

	// 开头先执行一次
	hb(l)
	ticker := time.NewTicker(t)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			l.info("heartbeat stopped")
			return
		case <-ticker.C:
			hb(l)
		}
	}
}

// revWithError 接收訊息並捕捉錯誤
func (l *Live) revWithError(ctx context.Context, ifError chan<- error) {
	msgCtx, msgCancel := context.WithCancel(ctx)
	defer l.info("receiving stopped")
	defer msgCancel()
	defer close(ifError)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if t, msg, err := l.ws.ReadMessage(); t == websocket.BinaryMessage && err == nil && len(msg) > 16 {
				go l.handle(msgCtx, msg)
			} else if err != nil {
				ifError <- err
				return
			}
		}
	}
}

func (l *Live) handle(ctx context.Context, b []byte) {
	defer l.report()
	body := b[16:]
	switch binary.BigEndian.Uint32(b[core.WsOpBegin:core.WsOpEnd]) {
	case core.WsOpEnterRoomSuccess:
		l.info("enter room success: %s", string(body))
		l.entered <- struct{}{}
	case core.WsOpHeartbeatReply:
		l.info("heartbeat reply: %d", binary.BigEndian.Uint32(body))
		// l.push(ctx, &message.HeartbeatReply{Raw: body}, nil)
	case core.WsOpMessage:
		// 压缩版本重新解包再调用，直到 ver==0
		switch binary.BigEndian.Uint16(b[core.WsVerBegin:core.WsVerEnd]) {
		case core.WsVerZlib:
			r, err := zlib.NewReader(bytes.NewReader(body))
			var de []byte
			if err == nil {
				de, err = io.ReadAll(r)
			}
			if err != nil {
				l.push(ctx, nil, fmt.Errorf("failed to decode zlib msg: %s", err))
				return
			}
			l.handles(ctx, l.split(de))
		case core.WsVerBrotli:
			de, err := io.ReadAll(brotli.NewReader(bytes.NewReader(body)))
			if err != nil {
				l.push(ctx, nil, fmt.Errorf("failed to decode brotli msg: %s", err))
				return
			}
			l.handles(ctx, l.split(de))
		case core.WsVerPlain:
			l.handlePlain(ctx, body)
		}
	}
}

// split 压缩过的body需要拆包
func (l *Live) split(b []byte) [][]byte {
	var packs [][]byte
	for i, size := uint32(0), uint32(0); i < uint32(len(b)); i += size {
		size = binary.BigEndian.Uint32(b[i : i+4])
		packs = append(packs, b[i:i+size])
	}
	return packs
}

func (l *Live) handles(ctx context.Context, bs [][]byte) {
	for _, b := range bs {
		go l.handle(ctx, b)
	}
}

func (l *Live) handlePlain(ctx context.Context, body []byte) {
	m, err := message.DefaultFilterParser.UnmarshalMessage(body)
	if err != nil {
		l.push(ctx, nil, fmt.Errorf("failed to unmarshal plain msg: %s", err))
		return
	}
	l.push(ctx, m, nil)
}

func (l *Live) push(ctx context.Context, msg message.Msg, err error) {
	go func(c context.Context, m message.Msg, e error) {
		// 五秒超时
		after := time.NewTimer(5 * time.Second)
		defer after.Stop()

		select {
		case <-c.Done():
			l.info("push stopped")
			return
		case <-after.C:
			return
		case l.Rev <- &Transport{Msg: m, Error: e}:
			return
		}
	}(ctx, msg, err)
}

func (l *Live) logf(format string, v ...interface{}) {
	if l.debug {
		l.logger.Printf(format, v...)
	}
}

func (l *Live) info(format string, v ...interface{}) {
	l.logf("[INFO] "+format, v...)
}

func (l *Live) error(format string, v ...interface{}) {
	l.logf("[ERROR] "+format, v)
}
