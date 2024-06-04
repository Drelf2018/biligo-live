package message

/*
	{
		"tail_icon": 0,
		"uid": 13729609,
		"uname": "阿羽AYu-",
		"uname_color": "",
		"dmscore": 12,
		"score": 1635849011802,
		"spread_desc": "",
		"timestamp": 1635849011,
		"identities": [
			1
		],
		"is_spread": 0,
		"roomid": 732602,
		"trigger_time": 1635849010751414020,
		"contribution": {
			"grade": 0
		},
		"fans_medal": {
			"medal_color": 12632256,
			"medal_color_start": 12632256,
			"medal_level": 10,
			"score": 10103,
			"target_id": 67141,
			"guard_level": 0,
			"icon_id": 0,
			"is_lighted": 0,
			"medal_name": "C酱",
			"special": "",
			"anchor_roomid": 47867,
			"medal_color_border": 12632256,
			"medal_color_end": 12632256
		},
		"msg_type": 1,
		"spread_info": ""
	}
*/
type InteractWord struct {
	TailIcon     int    `json:"tail_icon"`
	UID          int64  `json:"uid"`
	Uname        string `json:"uname"`
	UnameColor   string `json:"uname_color"`
	Dmscore      int    `json:"dmscore"`
	Score        int64  `json:"score"`
	SpreadDesc   string `json:"spread_desc"`
	Timestamp    int64  `json:"timestamp"`
	Identities   []int  `json:"identities"`
	IsSpread     int    `json:"is_spread"`
	RoomID       int    `json:"roomid"`
	TriggerTime  int64  `json:"trigger_time"`
	Contribution struct {
		Grade int `json:"grade"`
	} `json:"contribution"`
	FansMedal struct {
		MedalColor       int64  `json:"medal_color"`
		MedalColorStart  int64  `json:"medal_color_start"`
		MedalLevel       int    `json:"medal_level"`
		Score            int    `json:"score"`
		TargetID         int    `json:"target_id"`
		GuardLevel       int    `json:"guard_level"`
		IconID           int    `json:"icon_id"`
		IsLighted        int    `json:"is_lighted"`
		MedalName        string `json:"medal_name"`
		Special          string `json:"special"`
		AnchorRoomID     int64  `json:"anchor_roomid"`
		MedalColorBorder int64  `json:"medal_color_border"`
		MedalColorEnd    int64  `json:"medal_color_end"`
	} `json:"fans_medal"`
	MsgType    int    `json:"msg_type"`
	SpreadInfo string `json:"spread_info"`
}

func (InteractWord) Cmd() string {
	return CmdInteractWord
}

var _ = Register[InteractWord]()
