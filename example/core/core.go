package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"strings"

	api "github.com/Drelf2018/go-bilibili-api"
	"github.com/Drelf2018/req"
	"github.com/iyear/biligo-live/core"
	"github.com/iyear/biligo-live/message"
)

func init() {
	os.Mkdir("message", os.ModePerm)
}

func Show(c *core.Core, m message.Msg) {
	switch m := m.(type) {
	case message.EnterRoomSuccess:
		fmt.Printf("直播间 %d 连接成功\n", c.RoomID())
	case message.Danmaku:
		fmt.Println(m)
	case message.Raw:
		cmd := m.ParseCmd()
		name := req.NormalizeName(strings.ToLower(cmd))
		b, err := req.NewConverter(true, true).JSONToStruct(m, strings.ToLower(cmd))
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		buf := &bytes.Buffer{}
		buf.WriteString("package message\n\nconst Cmd")
		buf.WriteString(name)
		buf.WriteString(" = \"")
		buf.WriteString(cmd)
		buf.WriteString("\"\n\n")
		buf.Write(b)
		buf.WriteString("\n\nfunc (")
		buf.WriteString(name)
		buf.WriteString(") Cmd() string {\n\treturn Cmd")
		buf.WriteString(name)
		buf.WriteString("\n}")
		os.WriteFile("./message/"+cmd+".go", buf.Bytes(), os.ModePerm)
	}
}

func WriteError(err error) {
	if err != nil {
		os.WriteFile("error.log", []byte(err.Error()), os.ModePerm)
		panic(err)
	}
}

var cred = &api.Credential{
	SESSDATA:   "444e8f79%2C1745860278%2C1f8f3%2Aa1CjCOU8Do60BfJ-Bzf9_Nc4JQHnUTpkuAEr5tuQz9hNhTKE3yEv_aoNaQ3yZgf4josZYSVm1LUTJScC1aTENCNGl5R09CS3NKMjFXalhraElWR0U0ZlBtcW1JLWptdjN6eEtsS0E5VkFEbUJ2ZEZEYU9adFBLdklQVlY0UzNfRjRRVDNqZUx4SVF3IIEC",
	BiliJct:    "fae10d2598dac3a0a5365cd5ba7a7528",
	DedeUserID: "188888131",
}

func main() {
	verify, err := core.GetVerifyData(21452505, cred)
	if err != nil {
		panic(err)
	}
	ch := make(chan error)

	room, err := core.Default(message.ParseFunc(message.DefaultFilterParser, Show), ch)
	if err != nil {
		panic(err)
	}
	go room.Run(verify)

	for err := range ch {
		WriteError(err)
	}
}
