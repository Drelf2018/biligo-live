package utils

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Conn struct {
	*websocket.Conn
}

func (c *Conn) SendBytes(ver, op int, body []byte) error {
	return c.Conn.WriteMessage(websocket.BinaryMessage, Encode(ver, op, body))
}

func (c *Conn) SendControl(ver, op int, body []byte, deadline time.Time) error {
	return c.Conn.WriteControl(websocket.BinaryMessage, Encode(ver, op, body), deadline)
}

func (c *Conn) Send(ver, op int, v any) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return c.SendBytes(ver, op, body)
}

type VerifyData struct {
	// 自己的 UID
	UID int `json:"uid"`
	// 台主房间号
	RoomID int `json:"roomid"`
	// 协议版本 通常用 3
	Protover int `json:"protover"`
	// 用于区分用户的字符串
	Buvid string `json:"buvid"`
	// 平台 通常用 web
	Platform string `json:"platform"`
	// 类型 通常用 2
	Type int `json:"type"`
	// 用户私钥 与 Buvid 对应
	Key string `json:"key"`
}

func NewVerifyData(uid, room int, buvid, key string) VerifyData {
	return VerifyData{
		UID:      uid,
		RoomID:   room,
		Protover: 3,
		Buvid:    buvid,
		Platform: "web",
		Type:     2,
		Key:      key,
	}
}

func (c *Conn) SendVerifyData(data VerifyData) error {
	return c.Send(WsVerInt, WsOpEnterRoom, data)
}

var HeartbeatBody = Encode(WsVerPlain, WsOpHeartbeat, []byte("[object Object]"))

func (c *Conn) SendHeartbeat() error {
	return c.Conn.WriteMessage(websocket.BinaryMessage, HeartbeatBody)
}

func (c *Conn) Heartbeat(ctx context.Context) (err error) {
	err = c.SendHeartbeat()
	if err != nil {
		return
	}

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err = c.SendHeartbeat()
			if err != nil {
				return
			}
		}
	}
}

var DefaultHeader = http.Header{
	"Origin":     {"https://live.bilibili.com"},
	"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"},
}

func NewConn(dialer *websocket.Dialer, host string, header http.Header) (conn Conn, err error) {
	conn.Conn, _, err = dialer.Dial(host, header)
	return
}

func DefaultConn() (conn Conn) {
	conn, _ = NewConn(websocket.DefaultDialer, WsDefaultHost, DefaultHeader)
	return
}
