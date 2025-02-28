package main

import (
	_ "embed"
	"errors"
	"os"

	api "github.com/Drelf2018/go-bilibili-api"
	live "github.com/iyear/biligo-live"
	"github.com/iyear/biligo-live/message"
	"github.com/sirupsen/logrus"

	nested "github.com/antonfisher/nested-logrus-formatter"
)

var log = &logrus.Logger{
	Out:   os.Stdout,
	Hooks: make(logrus.LevelHooks),
	Formatter: &nested.Formatter{
		HideKeys:        true,
		NoColors:        true,
		TimestampFormat: "15:04:05",
		ShowFullLevel:   true,
	},
	Level: logrus.DebugLevel,
}

func Show(r *live.Room, m message.Msg) {
	switch m := m.(type) {
	default:
		log.Debug(m.Cmd(), m)
	case message.EnterRoomSuccess:
		log.Infof("直播间 %d 连接成功", m.RoomID)
	case message.Danmaku:
		log.Info(m)
	}
}

var cred = &api.Credential{DedeUserID: "0"}

// 重新拉起
func WithReload(fn func(*live.Room, []byte, error)) func(*live.Room, []byte, error) {
	return func(room *live.Room, data []byte, err error) {
		if room.Running() || errors.Is(err, live.ErrNoVerification) {
			fn(room, data, err)
			return
		}
		err = room.Connect(room.RoomID(), cred)
		if err != nil {
			go room.Error(nil, err)
			return
		}
		go room.Listen()
	}
}

func main() {
	room := &live.Room{OnMessage: Show, OnError: WithReload(func(r *live.Room, b []byte, err error) {
		log.Errorf("%v: %v: %v", r, err, string(b))
	})}
	err := room.Connect(22907643, cred)
	if err != nil {
		log.Panic(err)
	}
	room.Listen()

	ch := make(chan struct{})
	<-ch
}
