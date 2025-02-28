package message

// 消息内容不是 data 字段的值
// 因此不需要被包装 例如：
//
//	{
//		"cmd": "LIVE",
//		"live_key": "508011583643868085",
//		"voice_background": "",
//		"sub_session_key": "508011583643868085sub_time:1717090320",
//		"live_platform": "pc",
//		"live_model": 0,
//		"roomid": 14703541,
//		"live_time": 1717090319
//	}
type NoWrapper interface {
	NoWrapper()
}

// 接口 NoWrapper 的实现 直接嵌入结构体使用
type NoWrapperImpl struct{}

func (NoWrapperImpl) NoWrapper() {}

// 消息的命令别名
//
// 例如 CmdDanmakuWithCode 是 Danmaku 的命令别名
//
//	const CmdDanmaku         = "DANMU_MSG"             // 弹幕消息
//	const CmdDanmakuWithCode = "DANMU_MSG:4:0:2:2:2:0" // 特殊情况弹幕消息
type CmdAlias interface {
	CmdAlias() []string
}

// 消息接口
type Msg interface {
	Cmd() string
}

// 消息解析器
//
// 将单条消息解析成符合 Msg 接口，允许为空
type Parser interface {
	UnmarshalMessage([]byte) (Msg, error)
}
