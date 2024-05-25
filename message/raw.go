package message

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

type Raw []byte

func (r Raw) Cmd() string {
	return gjson.GetBytes(r, "cmd").Str
}

func (Raw) Parse(b []byte) (Msg, error) {
	return Raw(b), nil
}

func GetData[T any](raw []byte) (T, error) {
	var d struct {
		Data T `json:"data"`
	}
	return d.Data, json.Unmarshal(raw, &d)
}
