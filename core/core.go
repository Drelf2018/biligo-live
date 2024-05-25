package core

import (
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	msg "github.com/iyear/biligo-live/message"
	"github.com/iyear/biligo-live/utils"
)

type Core struct {
	utils.Conn
	interval time.Duration
	ch       chan msg.Msg
}

func (l *Core) Heartbeat() (err error) {
	err = l.SendHeartbeat()
	if err != nil {
		return
	}

	ticker := time.NewTicker(l.interval)
	defer ticker.Stop()

	for range ticker.C {
		err = l.SendHeartbeat()
		if err != nil {
			return
		}
	}

	return nil
}

func (l *Core) Run(data utils.VerifyData) {
	l.SendVerifyData(data)
	for {
		mt, b, err := l.Conn.ReadMessage()
		if err == nil && mt == websocket.BinaryMessage && len(b) > 16 {
			go l.Handle(b)
		}
	}
}

func (l *Core) Handle(b []byte) error {
	ver, op, body := utils.Decode(b)
	switch op {
	case utils.WsOpEnterRoomSuccess:
		return l.Heartbeat()
	case utils.WsOpMessage:
		switch ver {
		case utils.WsVerZlib:
			de, err := utils.ZlibDe(body)
			if err != nil {
				return err
			}
			return l.Handles(utils.Split(de))
		case utils.WsVerBrotli:
			de, err := utils.BrotliDe(body)
			if err != nil {
				return err
			}
			return l.Handles(utils.Split(de))
		case utils.WsVerPlain:
			msg, err := msg.Parse(body)
			if err != nil || msg == nil {
				return err
			}
			if l.ch != nil {
				l.ch <- msg
			}
		}
	}
	return nil
}

func (l *Core) Handles(bs [][]byte) error {
	errs := utils.NewSlice[error]()
	wg := new(sync.WaitGroup)
	wg.Add(len(bs))
	for _, b := range bs {
		b := b
		go func() {
			defer wg.Done()
			err := l.Handle(b)
			if err != nil {
				errs.Append(err)
			}
		}()
	}
	wg.Wait()
	return errors.Join(errs.Data()...)
}

func (l *Core) Channel() <-chan msg.Msg {
	return l.ch
}

func New(dialer *websocket.Dialer, host string, header http.Header, interval time.Duration, ch chan msg.Msg) (*Core, error) {
	conn, err := utils.NewConn(dialer, host, header)
	if err != nil {
		return nil, err
	}
	return &Core{conn, interval, ch}, nil
}

func Default() *Core {
	return &Core{utils.DefaultConn(), 30 * time.Second, make(chan msg.Msg, 1<<6)}
}
