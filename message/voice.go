package message

/*
	{
		"room_id": 34348,
		"category": 1,
		"apply_count": 1,
		"red_point": 1,
		"refresh": 1
	}
*/
type VoiceJoinList struct {
	RoomID     int64 `json:"room_id"`
	Category   int   `json:"category"`
	ApplyCount int   `json:"apply_count"`
	RedPoint   int   `json:"red_point"`
	Refresh    int   `json:"refresh"`
}

func (VoiceJoinList) Cmd() string {
	return CmdVoiceJoinList
}

var _ = Register[VoiceJoinList]()

/*
	{
		"apply_count": 1,
		"notify_count": 0,
		"red_point": 0,
		"room_id": 34348,
		"root_status": 1,
		"room_status": 1
	}
*/
type VoiceJoinRoomCountInfo struct {
	ApplyCount  int   `json:"apply_count"`
	NotifyCount int   `json:"notify_count"`
	RedPoint    int   `json:"red_point"`
	RoomID      int64 `json:"room_id"`
	RootStatus  int   `json:"root_status"`
	RoomStatus  int   `json:"room_status"`
}

func (VoiceJoinRoomCountInfo) Cmd() string {
	return CmdVoiceJoinRoomCountInfo
}

var _ = Register[VoiceJoinRoomCountInfo]()
