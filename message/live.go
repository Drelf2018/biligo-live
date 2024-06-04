package message

/*
	{
		"cmd": "LIVE",
		"live_key": "508011583643868085",
		"voice_background": "",
		"sub_session_key": "508011583643868085sub_time:1717090320",
		"live_platform": "pc",
		"live_model": 0,
		"roomid": 14703541,
		"live_time": 1717090319
	}
*/
type Live struct {
	LiveKey         string `json:"live_key"`
	VoiceBackground string `json:"voice_background"`
	SubSessionKey   string `json:"sub_session_key"`
	LivePlatform    string `json:"live_platform"`
	LiveModel       int    `json:"live_model"`
	RoomID          int    `json:"roomid"`
	LiveTime        int    `json:"live_time"`
}

func (Live) Cmd() string {
	return CmdLive
}

var _ = RegisterMsg[Live](true)
