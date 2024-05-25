package utils

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"io"

	"github.com/andybalholm/brotli"
)

func Write(l, n int) []byte {
	b := make([]byte, l)
	switch l {
	case 2:
		binary.BigEndian.PutUint16(b, uint16(n))
	case 4:
		binary.BigEndian.PutUint32(b, uint32(n))
	case 8:
		binary.BigEndian.PutUint64(b, uint64(n))
	}
	return b
}

func Encode(ver, op uint8, body []byte) []byte {
	header := make([]byte, 0, WsPackHeaderTotalLen)
	header = append(header, Write(WsPackageLen, len(body)+WsPackHeaderTotalLen)...)
	header = append(header, Write(WsHeaderLen, WsPackHeaderTotalLen)...)
	header = append(header, Write(WsVerLen, int(ver))...)
	header = append(header, Write(WsOpLen, int(op))...)
	header = append(header, Write(WsSequenceLen, WsHeaderDefaultSequence)...)

	return append(header, body...)
}

// Decode 必须确保 len(b)>16
func Decode(b []byte) (ver uint16, op uint32, body []byte) {
	ver = binary.BigEndian.Uint16(b[WsPackageLen+WsHeaderLen : WsPackageLen+WsHeaderLen+WsVerLen])
	op = binary.BigEndian.Uint32(b[WsPackageLen+WsHeaderLen+WsVerLen : WsPackageLen+WsHeaderLen+WsVerLen+WsOpLen])
	body = b[WsPackHeaderTotalLen:]
	return
}

func ZlibDe(src []byte) ([]byte, error) {
	var (
		r   io.ReadCloser
		o   bytes.Buffer
		err error
	)
	if r, err = zlib.NewReader(bytes.NewReader(src)); err != nil {
		return nil, err
	}
	if _, err = io.Copy(&o, r); err != nil {
		return nil, err
	}
	return o.Bytes(), nil
}

func BrotliDe(src []byte) ([]byte, error) {
	o := new(bytes.Buffer)
	r := brotli.NewReader(bytes.NewReader(src))
	if _, err := io.Copy(o, r); err != nil {
		return nil, err
	}
	return o.Bytes(), nil
}

func BrotliEn(src []byte) ([]byte, error) {
	b := new(bytes.Buffer)
	w := brotli.NewWriter(b)

	if _, err := w.Write(src); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func Split(b []byte) [][]byte {
	var packs [][]byte
	for i, size := uint32(0), uint32(0); i < uint32(len(b)); i += size {
		size = binary.BigEndian.Uint32(b[i : i+4])
		packs = append(packs, b[i:i+size])
	}
	return packs
}
