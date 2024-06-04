package utils

import (
	"compress/zlib"
	"encoding/binary"
	"io"

	"github.com/andybalholm/brotli"
)

func AppendData(header []byte, data ...int) []byte {
	for i := 0; i < len(data)-1; i += 2 {
		switch data[i] {
		case 2:
			header = binary.BigEndian.AppendUint16(header, uint16(data[i+1]))
		case 4:
			header = binary.BigEndian.AppendUint32(header, uint32(data[i+1]))
		case 8:
			header = binary.BigEndian.AppendUint64(header, uint64(data[i+1]))
		}
	}
	return header
}

func Encode(ver, op int, body []byte) []byte {
	l := WsPackHeaderTotalLen + len(body)
	header := AppendData(
		make([]byte, 0, l),
		WsPackageLen, l,
		WsHeaderLen, WsPackHeaderTotalLen,
		WsVerLen, ver,
		WsOpLen, op,
		WsSequenceLen, WsHeaderDefaultSequence,
	)
	return append(header, body...)
}

func ZlibReader(r io.Reader) ([]byte, error) {
	rc, err := zlib.NewReader(r)
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	return io.ReadAll(rc)
}

func BrotliReader(r io.Reader) ([]byte, error) {
	return io.ReadAll(brotli.NewReader(r))
}
