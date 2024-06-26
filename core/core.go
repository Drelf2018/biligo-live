package core

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"

	"github.com/gorilla/websocket"
	"github.com/iyear/biligo-live/message"
	"github.com/iyear/biligo-live/utils"
)

var ErrNoConn = errors.New("core: no existing conn")

type Core struct {
	utils.Conn
	Handler func(*Core, message.Msg)
	Err     chan error

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
	fmt.Println(runtime.Caller(1))
	if c.Err != nil {
		c.Err <- err
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

// 将压缩后的多条消息分割后逐一处理
func (c *Core) HandleAll(body []byte) error {
	l := len(body)
	errs := make([]error, 0, 8)
	for i, size := 0, 0; i < l; i += size {
		size = int(binary.BigEndian.Uint32(body[i : i+utils.WsPackageLen]))
		errs = append(errs, c.Handle(body[i+utils.WsPackHeaderTotalLen:i+size]))
	}
	return errors.Join(errs...)
}

func (c *Core) Run(data utils.VerifyData) error {
	return c.RunWithContext(context.Background(), data)
}

func (c *Core) RunWithContext(ctx context.Context, data utils.VerifyData) error {
	if c.Conn.Conn == nil {
		c.Push(ErrNoConn)
		return ErrNoConn
	}

	// 发送认证消息
	err := c.SendVerifyData(data)
	if c.Push(err) {
		return err
	}

	c.roomID = data.RoomID
	ctx, c.cancel = context.WithCancel(ctx)

	var (
		msg = make(utils.Message, utils.WsPackHeaderTotalLen)
		mt  int
		r   io.Reader
	)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			// 获取新消息的 io.Reader
			mt, r, err = c.Conn.NextReader()
			if mt == -1 || c.Push(err) || mt != websocket.BinaryMessage {
				continue
			}
			// 仅读取消息头 用于判断这条消息的类型 决定怎么处理
			_, err = r.Read(msg)
			if c.Push(err) {
				continue
			}
			// 分析消息
			switch msg.Op() {
			case utils.WsOpEnterRoomSuccess:
				go c.Heartbeat(ctx)
			case utils.WsOpMessage:
				// 必须先解码出 body([]byte) 后才能继续异步处理
				// 如果把 r(io.Reader) 放进异步会因为接收到新消息把当前消息覆写了
				var body []byte

				switch msg.Ver() {
				case utils.WsVerPlain:
					body, err = io.ReadAll(r)
					if !c.Push(err) {
						go c.Handle(body)
					}
					continue
				case utils.WsVerZlib:
					body, err = utils.ZlibReader(r)
				case utils.WsVerBrotli:
					body, err = utils.BrotliReader(r)
				}
				if !c.Push(err) {
					go c.HandleAll(body)
				}
			}
		}
	}
}

func (c *Core) Close() (err error) {
	c.roomID = 0
	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
	}
	if c.Conn.Conn != nil {
		err = c.Conn.Close()
		c.Conn.Conn = nil
	}
	return
}

func (c *Core) New(dialer *websocket.Dialer, host string, header http.Header) (err error) {
	c.Conn, err = utils.NewConn(dialer, host, header)
	return
}

func New(dialer *websocket.Dialer, host string, header http.Header, handler func(*Core, message.Msg)) (core *Core, err error) {
	core = &Core{Handler: handler}
	core.Conn, err = utils.NewConn(dialer, host, header)
	return
}

func Default(handler func(*Core, message.Msg)) *Core {
	return &Core{
		Conn:    utils.DefaultConn(),
		Handler: handler,
	}
}
