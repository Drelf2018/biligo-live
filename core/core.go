package core

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/gorilla/websocket"
	"github.com/iyear/biligo-live/message"
	"github.com/iyear/biligo-live/utils"
)

var ErrNoConn = errors.New("core: no existing conn")

type CoreError struct {
	error
	core *Core
}

func (c CoreError) Core() *Core {
	return c.core
}

type Handler func(*Core, message.Msg)

type Core struct {
	*websocket.Conn
	Handler Handler
	Err     chan CoreError

	roomID int
	cancel context.CancelFunc
}

func (c Core) String() string {
	return fmt.Sprintf("Core<%d>", c.roomID)
}

func (c Core) RoomID() int {
	return c.roomID
}

func (c *Core) Running() bool {
	return c.cancel != nil
}

func (c *Core) Push(err error) bool {
	if err == nil {
		return false
	}
	if c.Err != nil {
		c.Err <- CoreError{error: err, core: c}
	}
	return true
}

// 处理一条消息
func (c *Core) Handle(body []byte) error {
	if c.Handler == nil {
		return nil
	}
	msg, err := message.Parse(body)
	if c.Push(err) {
		return err
	}
	if msg != nil {
		c.Handler(c, msg)
	}
	return nil
}

func (c *Core) HandleReader(r io.Reader) {
	body, err := io.ReadAll(r)
	if c.Push(err) {
		return
	}
	for i, size := 0, 0; i < len(body); i += size {
		size = int(binary.BigEndian.Uint32(body[i : i+utils.WsPackageLen]))
		go c.Handle(body[i+utils.WsPackHeaderTotalLen : i+size])
	}
}

func (c *Core) HandleZlib(body []byte) {
	rc, err := zlib.NewReader(bytes.NewBuffer(body))
	if c.Push(err) {
		return
	}
	defer rc.Close()
	c.HandleReader(rc)
}

func (c *Core) HandleBrotli(body []byte) {
	c.HandleReader(brotli.NewReader(bytes.NewBuffer(body)))
}

func (c *Core) SendBytes(ver, op int, body []byte) error {
	return c.Conn.WriteMessage(websocket.BinaryMessage, utils.Encode(ver, op, body))
}

func (c *Core) SendControl(ver, op int, body []byte, deadline time.Time) error {
	return c.Conn.WriteControl(websocket.BinaryMessage, utils.Encode(ver, op, body), deadline)
}

func (c *Core) Send(ver, op int, v any) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return c.SendBytes(ver, op, body)
}

func (c *Core) SendVerifyData(data VerifyData) error {
	return c.Send(utils.WsVerHeartbeat, utils.WsOpEnterRoom, data)
}

func (c *Core) SendHeartbeat() error {
	return c.Conn.WriteMessage(websocket.BinaryMessage, utils.HeartbeatBody)
}

func (c *Core) Heartbeat(ctx context.Context) {
	c.Push(c.SendHeartbeat())

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.Push(c.SendHeartbeat())
		}
	}
}

func (c *Core) Run(data VerifyData) {
	c.RunWithContext(context.Background(), data)
}

func (c *Core) RunWithContext(ctx context.Context, data VerifyData) {
	if c.Conn == nil {
		c.Push(ErrNoConn)
		return
	}

	// 发送认证消息
	if c.Push(c.SendVerifyData(data)) {
		return
	}

	c.roomID = data.RoomID
	ctx, c.cancel = context.WithCancel(ctx)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// 获取新消息
			mt, body, err := c.Conn.ReadMessage()
			if mt != websocket.BinaryMessage || c.Push(err) {
				continue
			}
			switch binary.BigEndian.Uint32(body[utils.WsOpBegin:utils.WsOpEnd]) {
			case utils.WsOpEnterRoomSuccess:
				// 成功进入直播间后开启心跳
				// 并通过 c.Handler 通知调用方
				go c.Heartbeat(ctx)
				if c.Handler != nil {
					var enter message.EnterRoomSuccess
					json.Unmarshal(body[utils.WsPHTL:], &enter)
					go c.Handler(c, enter)
				}
			case utils.WsOpMessage:
				switch binary.BigEndian.Uint16(body[utils.WsVerBegin:utils.WsVerEnd]) {
				case utils.WsVerPlain:
					go c.Handle(body[utils.WsPHTL:])
				case utils.WsVerZlib:
					go c.HandleZlib(body[utils.WsPHTL:])
				case utils.WsVerBrotli:
					go c.HandleBrotli(body[utils.WsPHTL:])
				}
			}
		}
	}
}

func (c *Core) Close() error {
	c.roomID = 0
	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
	}
	if c.Conn != nil {
		err := c.Conn.Close()
		if err != nil {
			return err
		}
		c.Conn = nil
	}
	return nil
}

func (c *Core) New(dialer *websocket.Dialer, host string, header http.Header) (err error) {
	c.Conn, _, err = dialer.Dial(host, header)
	return
}

func New(dialer *websocket.Dialer, host string, header http.Header, handler func(*Core, message.Msg), err chan CoreError) (*Core, error) {
	c := &Core{Handler: handler, Err: err}
	return c, c.New(dialer, host, header)
}

func Default(handler func(*Core, message.Msg), err chan CoreError) *Core {
	c, _ := New(websocket.DefaultDialer, utils.WsDefaultHost, utils.DefaultHeader, handler, err)
	return c
}
