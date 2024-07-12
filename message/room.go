package message

/*
	{
		"uname": "白绫彡",
		"dmscore": 30,
		"operator": 1,
		"uid": 53342046
	}
*/
type RoomBlockMsg struct {
	Uname    string `json:"uname"`
	DmScore  int    `json:"dmscore"`
	Operator int    `json:"operator"`
	UID      int    `json:"uid"`
}

func (RoomBlockMsg) Cmd() string {
	return CmdRoomBlockMsg
}

var _ = Register[RoomBlockMsg]()

/*
	{
		"parent_area_id": 3,
		"area_name": "原神",
		"parent_area_name": "手游",
		"live_key": "181443822587250220",
		"sub_session_key": "181443822587250220sub_time:1635846195",
		"title": "快来直播间抽胡桃！托马！烈火拔刀！",
		"area_id": 321
	}
*/
type RoomChange struct {
	ParentAreaID   int    `json:"parent_area_id"`
	AreaName       string `json:"area_name"`
	ParentAreaName string `json:"parent_area_name"`
	LiveKey        string `json:"live_key"`
	SubSessionKey  string `json:"sub_session_key"`
	Title          string `json:"title"`
	AreaID         int    `json:"area_id"`
}

func (RoomChange) Cmd() string {
	return CmdRoomChange
}

var _ = Register[RoomChange]()

/*
	{
		"fans_club": 49182,
		"roomid": 545068,
		"fans": 1384297,
		"red_notice": -1
	}
*/
type RoomRealTimeMessageUpdate struct {
	FansClub  int   `json:"fans_club"`
	RoomID    int64 `json:"roomid"`
	Fans      int   `json:"fans"`
	RedNotice int   `json:"red_notice"`
}

func (RoomRealTimeMessageUpdate) Cmd() string {
	return CmdRoomRealTimeMessageUpdate
}

var _ = Register[RoomRealTimeMessageUpdate]()

/*
	{
		"room_id_list": [
			xxx
		]
	}
*/
type StopLiveRoomList struct {
	RoomIDList []int `json:"room_id_list"`
}

func (StopLiveRoomList) Cmd() string {
	return CmdStopLiveRoomList
}

var _ = Register[StopLiveRoomList]()

type EnterRoomSuccess struct {
	Code int `json:"code"`
}

func (EnterRoomSuccess) Cmd() string {
	return CmdEnterRoomSuccess
}
