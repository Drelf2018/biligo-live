package core

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/iyear/biligo-live/message"
	"github.com/iyear/biligo-live/utils"
)

type Core struct {
	utils.Conn
	Handler func(*Core, message.Msg)

	roomID int
	err    chan error
	ctx    context.Context
	cancel func()
}

func (c Core) String() string {
	return fmt.Sprintf("Room<%d>", c.roomID)
}

func (c Core) RoomID() int {
	return c.roomID
}

func (c *Core) Running() bool {
	return c.ctx.Err() == nil
}

func (c *Core) Dispatch(body []byte) error {
	if c.Handler == nil {
		return nil
	}
	msg, err := message.Parse(body)
	if err != nil || msg == nil {
		return err
	}
	c.Handler(c, msg)
	return nil
}

var ErrInvalidVer = errors.New("core: invalid ver")

func (c *Core) ParseBytes(ver uint16, body []byte) (err error) {
	switch ver {
	case utils.WsVerPlain:
		err = c.Dispatch(body)
	case utils.WsVerZlib, utils.WsVerBrotli:
		l := len(body)
		for i, size := 0, 0; i < l; i += size {
			size = int(binary.BigEndian.Uint32(body[i : i+utils.WsPackageLen]))
			err = c.ParseBytes(
				binary.BigEndian.Uint16(body[i+utils.WsVerBegin:i+utils.WsVerEnd]),
				body[i+utils.WsPackHeaderTotalLen:i+size],
			)
			if err != nil {
				break
			}
		}
	default:
		err = ErrInvalidVer
	}
	if c.err != nil && err != nil {
		c.err <- err
	}
	return
}

func (c *Core) Run(data utils.VerifyData) error {
	return c.RunWithContext(context.Background(), data)
}

func (c *Core) RunWithContext(ctx context.Context, data utils.VerifyData) (err error) {
	c.roomID = data.RoomID
	c.ctx, c.cancel = context.WithCancel(ctx)

	// 发送认证消息
	err = c.SendVerifyData(data)
	if err != nil {
		return
	}

	var (
		msg = make(utils.Message, utils.WsPackHeaderTotalLen)
		mt  int
		r   io.Reader
	)

	for {
		select {
		case <-c.ctx.Done():
			return
		case err = <-c.err:
			c.Close()
		default:
			// 获取新消息的 io.Reader
			mt, r, err = c.Conn.NextReader()
			if err != nil {
				return
			}
			if mt != websocket.BinaryMessage {
				continue
			}
			// 仅读取消息头 用于判断这条消息的类型 决定怎么处理
			_, err = r.Read(msg)
			if err != nil {
				return
			}
			switch msg.Op() {
			case utils.WsOpEnterRoomSuccess:
				go c.Heartbeat(c.ctx)
			case utils.WsOpMessage:
				// 根据 ver 选择解码方式
				// 这里必须先解码成 body([]byte) 后才能放进异步
				// 如果把 r(io.Reader) 放进异步去解码可能会因为接收到新消息把当前消息覆写了
				var (
					ver  = msg.Ver()
					body []byte
				)
				switch ver {
				case utils.WsVerPlain:
					body, err = io.ReadAll(r)
				case utils.WsVerZlib:
					body, err = utils.ZlibReader(r)
				case utils.WsVerBrotli:
					body, err = utils.BrotliReader(r)
				}
				if err != nil {
					return
				}
				go c.ParseBytes(ver, body)
			}
		}
	}
}

func (c *Core) Close() error {
	if c.cancel != nil {
		c.cancel()
	}
	if c.err != nil {
		close(c.err)
		c.err = nil
	}
	return c.Conn.Close()
}

func New(dialer *websocket.Dialer, host string, header http.Header, handler func(*Core, message.Msg)) (core *Core, err error) {
	core = &Core{
		Handler: handler,
		err:     make(chan error, 16),
	}
	core.Conn, err = utils.NewConn(dialer, host, header)
	return
}

func Default(handler func(*Core, message.Msg)) *Core {
	return &Core{
		Conn:    utils.DefaultConn(),
		Handler: handler,
		err:     make(chan error, 16),
	}
}
