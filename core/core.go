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
)

var ErrConnNotExist = errors.New("core: Conn does not exist")

// 核心错误
type CoreError struct {
	err error

	// 产生该错误的核心
	core *Core

	// 发生错误区域的相关参数
	// 方便记录和排查错误
	args []any
}

func (c CoreError) Unwrap() error {
	return c.err
}

func (c CoreError) Core() *Core {
	return c.core
}

func (c CoreError) Args() []any {
	return c.args
}

func (c CoreError) Error() string {
	return fmt.Sprintf("core: the %s occured an error: %v", c.core, c.err)
}

// 用于连接直播间的核心
type Core struct {
	// 对 *websocket.Conn 进行了封装
	// 使用时一般不需要手动操作
	*websocket.Conn

	// 用来处理接收到的消息 必须是并发安全的
	// 考虑到可能多个直播间共用一个处理函数
	// 这里还会提供发出消息的核心作为参数
	Handler func(*Core, []byte)

	// 用来接收核心运行时产生的错误
	Err chan error

	// 直播间号
	roomID int

	// 断开连接和心跳
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

// 向 Err 通道发送一个经过封装的错误
func (c *Core) Error(err error, args ...any) {
	if err == nil {
		return
	}
	if c.Err != nil {
		c.Err <- CoreError{err: err, core: c, args: args}
	}
}

// 处理一条消息
func (c *Core) Handle(body []byte) {
	if c.Handler != nil {
		c.Handler(c, body)
	}
}

// 处理读入器中的多条消息
func (c *Core) HandleReader(r io.Reader) {
	body, err := io.ReadAll(r)
	if err != nil {
		c.Error(err, body)
		return
	}
	// 根据头部字段提供的长度分割多条消息
	// 并逐一并发处理
	for i, size := 0, 0; i < len(body); i += size {
		size = int(binary.BigEndian.Uint32(body[i : i+WsPackageLen]))
		go c.Handle(body[i+WsPackHeaderTotalLen : i+size])
	}
}

// 处理 zlib 压缩的多条消息
func (c *Core) HandleZlib(body []byte) {
	rc, err := zlib.NewReader(bytes.NewReader(body))
	if err != nil {
		c.Error(err, body)
		return
	}
	defer rc.Close()
	c.HandleReader(rc)
}

// 处理 brotli 压缩的多条消息
func (c *Core) HandleBrotli(body []byte) {
	c.HandleReader(brotli.NewReader(bytes.NewReader(body)))
}

// 发送认证包
func (c *Core) SendVerifyData(data VerifyData) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.Conn.WriteMessage(websocket.BinaryMessage, Encode(WsVerHeartbeat, WsOpEnterRoom, body))
}

// 发送心跳
func (c *Core) SendHeartbeat() error {
	return c.Conn.WriteMessage(websocket.BinaryMessage, HeartbeatBody)
}

// 启动核心
func (c *Core) Run(data VerifyData) {
	c.RunWithContext(context.Background(), data)
}

// 携带上下文时启动核心
func (c *Core) RunWithContext(ctx context.Context, data VerifyData) {
	if c.Conn == nil {
		c.Error(ErrConnNotExist)
		return
	}

	// 发送认证消息
	if err := c.SendVerifyData(data); err != nil {
		c.Error(err)
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
			messageType, body, err := c.Conn.ReadMessage()
			if err != nil {
				c.Error(err, body)
				continue
			}
			if messageType != websocket.BinaryMessage {
				continue
			}
			switch binary.BigEndian.Uint32(body[WsOpBegin:WsOpEnd]) {
			case WsOpEnterRoomSuccess:
				// 成功进入直播间后开启心跳
				// 并通过 c.Handler 通知调用方
				go func() {
					c.Error(c.SendHeartbeat())

					ticker := time.NewTicker(30 * time.Second)
					defer ticker.Stop()

					for {
						select {
						case <-ctx.Done():
							return
						case <-ticker.C:
							c.Error(c.SendHeartbeat())
						}
					}
				}()
				if c.Handler != nil {
					enter := fmt.Sprintf(`{"cmd":"ENTER","data":%s}`, body[WsPHTL:])
					go c.Handler(c, []byte(enter))
				}
			case WsOpMessage:
				switch binary.BigEndian.Uint16(body[WsVerBegin:WsVerEnd]) {
				case WsVerPlain:
					go c.Handle(body[WsPHTL:])
				case WsVerZlib:
					go c.HandleZlib(body[WsPHTL:])
				case WsVerBrotli:
					go c.HandleBrotli(body[WsPHTL:])
				}
			}
		}
	}
}

// 断开连接
//
// 此操作会关闭并清空 *websocket.Conn 字段
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

// 新建 *websocket.Conn
func (c *Core) New(dialer *websocket.Dialer, host string, header http.Header) (err error) {
	c.Conn, _, err = dialer.Dial(host, header)
	return
}

// 新建核心
func New(dialer *websocket.Dialer, host string, header http.Header, handler func(*Core, []byte), err chan error) (*Core, error) {
	c := &Core{Handler: handler, Err: err}
	return c, c.New(dialer, host, header)
}

// 默认核心
func Default(handler func(*Core, []byte), err chan error) (*Core, error) {
	return New(websocket.DefaultDialer, WsDefaultHost, DefaultHeader, handler, err)
}
