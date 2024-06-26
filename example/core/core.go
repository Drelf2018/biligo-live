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
	verify := utils.NewVerifyData(0, 21452505, GetBuvid3(), GetKey(21452505))

	c := core.Default(Show)
	c.Err = make(chan error)
	go c.Run(verify)

	for err := range c.Err {
		WriteError(err)
	}
}
