package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/iyear/biligo-live/core"
	"github.com/iyear/biligo-live/message"
	"github.com/iyear/biligo-live/utils"
)

func init() {
	os.Mkdir("messages", os.ModePerm)
	message.UseRaw = true
}

func Show(c *core.Core, m message.Msg) {
	switch m := m.(type) {
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

		_, err = buf.WriteTo(file)
		if err != nil {
			os.WriteFile("error.log", []byte(err.Error()), os.ModePerm)
		}
	}
}

func writeError(err error) {
	if err != nil {
		os.WriteFile("error.log", []byte(err.Error()), os.ModePerm)
	}
}

func Start(room int) {
	verify := utils.NewVerifyData(0, room, GetBuvid3(), GetKey(room))
	writeError(core.Default(Show).Run(verify))
}

func main() {
	room := flag.String("room", "", "roomid")
	flag.Parse()

	roomid, err := strconv.Atoi(*room)
	writeError(err)
	Start(roomid)
}
