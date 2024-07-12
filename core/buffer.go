package core

import (
	"bytes"
	"sync"
)

var bufferPool = &sync.Pool{
	New: func() any {
		return &bytes.Buffer{}
	},
}

func GetBuffer(body []byte) (buf *bytes.Buffer) {
	buf = bufferPool.Get().(*bytes.Buffer)
	buf.Write(body)
	return
}

func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}
