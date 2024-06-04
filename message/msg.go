package message

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

type Msg interface {
	Cmd() string
}

type Handler func([]byte) (Msg, error)

type base[T Msg] struct {
	Data T `json:"data"`
}

func Unmarshal[T Msg](raw []byte) (Msg, error) {
	var b base[T]
	return b.Data, json.Unmarshal(raw, &b)
}

func UnmarshalDirect[T Msg](raw []byte) (Msg, error) {
	var zero T
	return zero, json.Unmarshal(raw, &zero)
}

var messages = make(map[string]Handler)

func RegisterCmd(cmd string, handler Handler) bool {
	_, exists := messages[cmd]
	if !exists {
		messages[cmd] = handler
	}
	return !exists
}

func RegisterMsg[T Msg](direct bool) bool {
	var msg T
	_, exists := messages[msg.Cmd()]
	if exists {
		return false
	}
	if direct {
		messages[msg.Cmd()] = UnmarshalDirect[T]
	} else {
		messages[msg.Cmd()] = Unmarshal[T]
	}
	return true
}

func Register[T Msg]() bool {
	return RegisterMsg[T](false)
}

func Delete(cmd string) {
	delete(messages, cmd)
}

func Clear() {
	for k := range messages {
		delete(messages, k)
	}
}

var UseRaw = false

func Parse(raw []byte) (msg Msg, err error) {
	if handler, ok := messages[gjson.GetBytes(raw, "cmd").Str]; ok {
		return handler(raw)
	}
	if UseRaw {
		msg = Raw(raw)
	}
	return
}
