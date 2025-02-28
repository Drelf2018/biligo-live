package message

import (
	"strconv"

	"github.com/tidwall/gjson"
)

// 原始信息
type Raw []byte

func (Raw) Cmd() string {
	return CmdRaw
}

func (r Raw) ParseCmd() string {
	return gjson.GetBytes(r, "cmd").Str
}

func (r Raw) Index(idx int) []byte {
	return []byte(gjson.GetBytes(r, strconv.Itoa(idx)).Raw)
}

func (r Raw) Int(idx int) int {
	return int(gjson.GetBytes(r, strconv.Itoa(idx)).Num)
}

func (r Raw) String() string {
	return "Raw" + string(r)
}

type EnterRoomSuccess struct {
	RoomID int `json:"roomid"`
}

func (EnterRoomSuccess) Cmd() string {
	return CmdEnterRoomSuccess
}
