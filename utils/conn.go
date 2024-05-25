package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Conn struct {
	*websocket.Conn
}

// Encode 函数为加密函数 不重要

func (c *Conn) SendBytes(ver, op uint8, body []byte) error {
	return c.Conn.WriteMessage(websocket.BinaryMessage, Encode(ver, op, body))
}

func (c *Conn) SendControl(ver, op uint8, body []byte, deadline time.Time) error {
	return c.Conn.WriteControl(websocket.BinaryMessage, Encode(ver, op, body), deadline)
}

func (c *Conn) Send(ver, op uint8, body any) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	return c.SendBytes(ver, op, b)
}

type VerifyData struct {
	Uid      int    `json:"uid"`
	Roomid   int    `json:"roomid"`
	Protover int    `json:"protover"`
	Buvid    string `json:"buvid"`
	Platform string `json:"platform"`
	Type     int    `json:"type"`
	Key      string `json:"key"`
}

func (c *Conn) SendVerifyData(data VerifyData) error {
	return c.Send(WsVerInt, WsOpEnterRoom, data)
}

var HeartbeatBody = []byte("[object Object]")

func (c *Conn) SendHeartbeat() error {
	return c.SendBytes(WsVerPlain, WsOpHeartbeat, HeartbeatBody)
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
