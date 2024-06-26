package utils

const (
	WsOpHeartbeat        = 2 // 心跳
	WsOpHeartbeatReply   = 3 // 心跳回应
	WsOpMessage          = 5 // 弹幕消息等
	WsOpEnterRoom        = 7 // 请求进入房间
	WsOpEnterRoomSuccess = 8 // 进房回应
)

const (
	WsPackHeaderTotalLen = 16 // 头部字节大小
	WsPackageLen         = 4
	WsHeaderLen          = 2
	WsVerLen             = 2
	WsOpLen              = 4
	WsSequenceLen        = 4

	WsVerBegin = WsPackageLen + WsHeaderLen
	WsVerEnd   = WsVerBegin + WsVerLen

	WsOpBegin = WsPackageLen + WsHeaderLen + WsVerLen
	WsOpEnd   = WsOpBegin + WsOpLen
)

// ws header default
const (
	WsHeaderDefaultSequence = 1
)

// version protocol
const (
	WsVerPlain     = 0
	WsVerHeartbeat = 1
	WsVerZlib      = 2
	WsVerBrotli    = 3
)

const (
	WsDefaultHost = "wss://broadcastlv.chat.bilibili.com/sub"
)
