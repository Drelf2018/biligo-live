package live

import (
	"bytes"
	"compress/zlib"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	api "github.com/Drelf2018/go-bilibili-api"
	"github.com/andybalholm/brotli"
	"github.com/gorilla/websocket"
	"github.com/iyear/biligo-live/core"
	"github.com/iyear/biligo-live/message"
)

type Room struct {
	// 处理消息 必须并发安全
	OnMessage func(*Room, message.Msg)

	// 处理错误 必须并发安全 可为空
	OnError func(*Room, []byte, error)

	// 连接弹幕服务器的实例 为空时使用默认地址连接
	Room *core.Room

	// 消息解析器 为空时使用内置解析器
	Parser message.Parser

	// 最后一次成功发送认证包后的真实直播间号
	roomID int

	// 是否发送认证包
	verified bool

	// 断开连接和心跳
	cancel context.CancelFunc
}

func (r Room) String() string {
	return fmt.Sprintf("Room[%d]", r.roomID)
}

func (r Room) RoomID() int {
	return r.roomID
}

func (r Room) Running() bool {
	return r.cancel != nil
}

// 处理错误
func (r *Room) Error(data []byte, err error) {
	if r.OnError != nil {
		r.OnError(r, data, err)
	}
}

// 处理一条消息
func (r *Room) Handle(data []byte) {
	if r.OnMessage == nil || r.Parser == nil {
		return
	}
	msg, err := r.Parser.UnmarshalMessage(data)
	if err != nil {
		r.Error(data, err)
		return
	}
	r.OnMessage(r, msg)
}

// 处理多条消息
func (r *Room) HandleReader(rd io.Reader) {
	if r.OnMessage == nil || r.Parser == nil {
		return
	}
	data, err := io.ReadAll(rd)
	if err != nil {
		r.Error(data, err)
		return
	}
	// 根据头部字段提供的长度分割多条消息并逐一并发处理
	for i, size := 0, 0; i < len(data); i += size {
		size = int(binary.BigEndian.Uint32(data[i : i+core.WsPackageLen]))
		go r.Handle(data[i+core.WsPackHeaderTotalLen : i+size])
	}
}

// 处理 zlib 压缩的多条消息
func (r *Room) HandleZlib(data []byte) {
	if r.OnMessage == nil || r.Parser == nil {
		return
	}
	rc, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		r.Error(data, err)
		return
	}
	defer rc.Close()
	r.HandleReader(rc)
}

func (r *Room) init() (err error) {
	if r.Room == nil {
		r.Room, err = core.DefaultRoom()
	}
	return
}

// 发送验证消息
func (r *Room) Verify(verify core.VerifyData) error {
	if r.verified {
		return nil
	}
	err := r.init()
	if err != nil {
		return err
	}
	err = r.Room.SendVerifyData(verify)
	if err == nil {
		r.verified = true
		r.roomID = verify.RoomID
	}
	return err
}

// 连接直播间
func (r *Room) Connect(roomid int, credential *api.Credential) error {
	verify, err := core.GetVerifyData(roomid, credential)
	if err != nil {
		return err
	}
	return r.Verify(verify)
}

func (r *Room) StartHeartbeat(ctx context.Context, interval time.Duration) {
	err := r.Room.SendHeartbeat()
	if err != nil {
		r.Error(nil, err)
	}

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err = r.Room.SendHeartbeat()
			if err != nil {
				r.Error(nil, err)
			}
		}
	}
}

var ErrNoVerification = errors.New("live: no verification")

// 携带上下文时启动监听
func (r *Room) ListenWithContext(ctx context.Context) {
	// 初始化
	if !r.verified {
		r.Error(nil, ErrNoVerification)
		return
	}
	// 使用默认解析器
	if r.Parser == nil {
		r.Parser = message.DefaultFilterParser
	}
	// 内部用于断开连接的上下文
	ctx, r.cancel = context.WithCancel(ctx)
	// 轮询接收消息
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// 获取消息
			op, ver, data, err := r.Room.Read()
			if err != nil {
				if e, ok := err.(core.ReadError); ok {
					if e.MessageType() == -1 {
						errShutdown := r.Shutdown()
						if errShutdown != nil {
							r.Error(nil, errShutdown)
						}
					}
				}
				go r.Error(data, err)
				continue
			}
			// 处理消息
			switch op {
			case core.WsOpEnterRoomSuccess:
				go r.StartHeartbeat(ctx, 30*time.Second)
				go r.Handle([]byte(`{"cmd":"ENTER","roomid":` + strconv.Itoa(r.roomID) + "}"))
			case core.WsOpMessage:
				switch ver {
				case core.WsVerPlain:
					go r.Handle(data)
				case core.WsVerZlib:
					go r.HandleZlib(data)
				case core.WsVerBrotli:
					go r.HandleReader(brotli.NewReader(bytes.NewReader(data)))
				}
			}
		}
	}
}

// 启动监听
func (r *Room) Listen() {
	r.ListenWithContext(context.Background())
}

// 断开直播间连接
//
// 此操作会关闭并清空 r.Room 字段
func (r *Room) Shutdown() error {
	r.verified = false
	if r.cancel != nil {
		r.cancel()
		r.cancel = nil
	}
	if r.Room != nil {
		err := (*websocket.Conn)(r.Room).Close()
		if err != nil {
			return err
		}
		r.Room = nil
	}
	return nil
}
