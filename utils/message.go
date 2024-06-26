package utils

import (
	"encoding/binary"
	"io"
)

type Message []byte

func (m Message) Ver() uint16 {
	return binary.BigEndian.Uint16(m[WsVerBegin:WsVerEnd])
}

func (m Message) Op() uint32 {
	return binary.BigEndian.Uint32(m[WsOpBegin:WsOpEnd])
}

func (m Message) ReadFrom(r io.Reader) (int64, error) {
	n, err := r.Read(m)
	return int64(n), err
}

func (m Message) Write(p []byte) (n int, err error) {
	return copy(p, m), nil
}

var _ io.ReaderFrom = &Message{}
