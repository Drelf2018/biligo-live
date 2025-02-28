package core

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// 读取消息时的错误
type ReadError struct {
	err         error
	messageType int
}

func (r ReadError) Unwrap() error {
	return r.err
}

func (r ReadError) Error() string {
	return r.err.Error()
}

func (r ReadError) MessageType() int {
	return r.messageType
}

// 最小直播间连接实例
//
// 连接流程：先按照 websocket.Conn 进行连接，接着发送认证包。
// 此时读取消息会收到一个进房成功事件，接着每一段时间发送心跳就可以保持连接。
type Room websocket.Conn

// 发送认证包
func (r *Room) SendVerifyData(verify VerifyData) error {
	body, err := json.Marshal(verify)
	if err != nil {
		return err
	}
	return (*websocket.Conn)(r).WriteMessage(websocket.BinaryMessage, Encode(WsVerHeartbeat, WsOpEnterRoom, body))
}

// 发送一次心跳
func (r *Room) SendHeartbeat() error {
	return (*websocket.Conn)(r).WriteMessage(websocket.BinaryMessage, HeartbeatBody)
}

// 读取一条消息
func (r *Room) Read() (op uint32, ver uint16, data []byte, err error) {
	mt, data, err := (*websocket.Conn)(r).ReadMessage()
	if err != nil || mt != websocket.BinaryMessage {
		err = ReadError{err, mt}
		return
	}

	op = binary.BigEndian.Uint32(data[WsOpBegin:WsOpEnd])
	switch op {
	case WsOpHeartbeatReply, WsOpEnterRoomSuccess:
		return
	case WsOpMessage:
		ver = binary.BigEndian.Uint16(data[WsVerBegin:WsVerEnd])
		data = data[WsPackHeaderTotalLen:]
		return
	default:
		err = fmt.Errorf("live/core.Read: invalid operation: %d", op)
		return
	}
}

func NewRoom(dialer *websocket.Dialer, host string, header http.Header) (*Room, error) {
	conn, _, err := dialer.Dial(host, header)
	return (*Room)(conn), err
}

func WithHost(host string) (*Room, error) {
	return NewRoom(websocket.DefaultDialer, host, DefaultHeader)
}

func DefaultRoom() (*Room, error) {
	return NewRoom(websocket.DefaultDialer, WsDefaultHost, DefaultHeader)
}
