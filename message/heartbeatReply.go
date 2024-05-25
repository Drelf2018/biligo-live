package message

import "encoding/binary"

var _ = Register(HeartbeatReply{})

type HeartbeatReply struct {
	Raw
}

func (m *HeartbeatReply) GetHot() int {
	return int(binary.BigEndian.Uint32(m.Raw))
}
