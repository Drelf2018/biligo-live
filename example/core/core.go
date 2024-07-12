package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/iyear/biligo-live/core"
	"github.com/iyear/biligo-live/message"
)

func init() {
	os.Mkdir("messages", os.ModePerm)
	message.UseRaw = true
}

func Show(c *core.Core, m message.Msg) {
	switch m := m.(type) {
	case message.EnterRoomSuccess:
		fmt.Printf("直播间 %d 连接成功\n", c.RoomID())
	case message.Danmaku:
		fmt.Println(m)
	case message.Raw:
		file, err := os.OpenFile("./messages/"+m.ParseCmd()+".json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			t := time.Now().UnixMilli()
			file, _ = os.OpenFile("./messages/"+strconv.Itoa(int(t))+".json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		}
		defer file.Close()

		buf := &bytes.Buffer{}
		json.Indent(buf, m, "", "\t")
		buf.WriteTo(file)
	}
}

func WriteError(err error) {
	if err != nil {
		os.WriteFile("error.log", []byte(err.Error()), os.ModePerm)
		panic(err)
	}
}

func main() {
	room := 21452505
	verify := core.NewVerifyData(0, room, GetBuvid3(), GetKey(room))
	errCh := make(chan core.CoreError)

	go core.Default(Show, errCh).Run(verify)

	for err := range errCh {
		WriteError(err)
	}
}
