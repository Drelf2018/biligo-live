package utils

// ops
const (
	WsOpHeartbeat        = 2 // 心跳
	WsOpHeartbeatReply   = 3 // 心跳回应
	WsOpMessage          = 5 // 弹幕消息等
	WsOpEnterRoom        = 7 // 请求进入房间
	WsOpEnterRoomSuccess = 8 // 进房回应
)

// Header
const (
	WsPackHeaderTotalLen = 16 // 头部字节大小
	WsPackageLen         = 4
	WsHeaderLen          = 2
	WsVerLen             = 2
	WsOpLen              = 4
	WsSequenceLen        = 4
)

// ws header default
const (
	WsHeaderDefaultSequence = 1
)

// version protocol
const (
	WsVerPlain  = 0
	WsVerInt    = 1
	WsVerZlib   = 2
	WsVerBrotli = 3
)

const (
	WsDefaultHost = "wss://broadcastlv.chat.bilibili.com/sub"
)
