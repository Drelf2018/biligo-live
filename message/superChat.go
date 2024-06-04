package message

/*
	{
		"background_bottom_color": "#427D9E",
		"token": "D22854F7",
		"background_color_end": "#29718B",
		"background_image": "https://i0.hdslb.com/bfs/live/a712efa5c6ebc67bafbe8352d3e74b820a00c13e.png",
		"background_icon": "",
		"background_price_color": "#7DA4BD",
		"dmscore": 128,
		"id": 2575658,
		"user_info": {
			"user_level": 20,
			"face_frame": "http://i0.hdslb.com/bfs/live/9b3cfee134611c61b71e38776c58ad67b253c40a.png",
			"guard_level": 2,
			"level_color": "#61c05a",
			"manager": 0,
			"uname": "卡纸哥我宣你",
			"title": "title-111-1",
			"face": "http://i0.hdslb.com/bfs/face/6badd87b9bf8c13c90fcb2c2b1b93b01e4b02664.jpg",
			"is_main_vip": 1,
			"is_svip": 0,
			"is_vip": 0,
			"name_color": "#E17AFF"
		},
		"is_send_audit": 1,
		"price": 50,
		"background_color": "#DBFFFD",
		"color_point": 0.7,
		"gift": {
			"gift_id": 12000,
			"gift_name": "醒目留言",
			"num": 1
		},
		"medal_info": {
			"target_id": 8739477,
			"anchor_roomid": 7777,
			"anchor_uname": "老实憨厚的笑笑",
			"guard_level": 2,
			"medal_color": "#6154c",
			"medal_color_end": 6850801,
			"medal_level": 27,
			"special": "",
			"icon_id": 0,
			"is_lighted": 1,
			"medal_color_border": 16771156,
			"medal_color_start": 398668,
			"medal_name": "德云色"
		},
		"trans_mark": 0,
		"ts": 1635749378,
		"background_color_start": "#4EA4C5",
		"end_time": 1635749498,
		"message_font_color": "#A3F6FF",
		"rate": 1000,
		"message_trans": "",
		"start_time": 1635749378,
		"is_ranked": 1,
		"message": "熊神可以打一拳旁边那个大胖子吗",
		"time": 120,
		"uid": 12777723
	}
*/
type SuperChatMessage struct {
	BackgroundBottomColor string `json:"background_bottom_color"`
	Token                 string `json:"token"`
	BackgroundColorEnd    string `json:"background_color_end"`
	BackgroundImage       string `json:"background_image"`
	BackgroundIcon        string `json:"background_icon"`
	BackgroundPriceColor  string `json:"background_price_color"`
	DmScore               int    `json:"dmscore"`
	ID                    int64  `json:"id"`
	UserInfo              struct {
		UserLevel  int    `json:"user_level"`
		FaceFrame  string `json:"face_frame"`
		GuardLevel int    `json:"guard_level"`
		LevelColor string `json:"level_color"`
		Manager    int    `json:"manager"`
		Uname      string `json:"uname"`
		Title      string `json:"title"`
		Face       string `json:"face"`
		IsMainVip  int    `json:"is_main_vip"`
		IsSvip     int    `json:"is_svip"`
		IsVip      int    `json:"is_vip"`
		NameColor  string `json:"name_color"`
	} `json:"user_info"`
	IsSendAudit     int     `json:"is_send_audit"`
	Price           int     `json:"price"`
	BackgroundColor string  `json:"background_color"`
	ColorPoint      float64 `json:"color_point"`
	Gift            struct {
		GiftID   int64  `json:"gift_id"`
		GiftName string `json:"gift_name"`
		Num      int    `json:"num"`
	} `json:"gift"`
	MedalInfo struct {
		TargetID         int64  `json:"target_id"`
		AnchorRoomID     int    `json:"anchor_roomid"`
		AnchorUname      string `json:"anchor_uname"`
		GuardLevel       int    `json:"guard_level"`
		MedalColor       string `json:"medal_color"`
		MedalColorEnd    int    `json:"medal_color_end"`
		MedalLevel       int    `json:"medal_level"`
		Special          string `json:"special"`
		IconID           int64  `json:"icon_id"`
		IsLighted        int    `json:"is_lighted"`
		MedalColorBorder int    `json:"medal_color_border"`
		MedalColorStart  int    `json:"medal_color_start"`
		MedalName        string `json:"medal_name"`
	} `json:"medal_info"`
	TransMark            int    `json:"trans_mark"`
	Ts                   int    `json:"ts"`
	BackgroundColorStart string `json:"background_color_start"`
	EndTime              int64  `json:"end_time"`
	MessageFontColor     string `json:"message_font_color"`
	Rate                 int    `json:"rate"`
	MessageTrans         string `json:"message_trans"`
	StartTime            int64  `json:"start_time"`
	IsRanked             int    `json:"is_ranked"`
	Message              string `json:"message"`
	Time                 int64  `json:"time"`
	UID                  int64  `json:"uid"`
}

func (SuperChatMessage) Cmd() string {
	return CmdSuperChatMessage
}

var _ = Register[SuperChatMessage]()

/*
	{
		"uid": "179810804",
		"is_ranked": 1,
		"medal_info": {
			"medal_color": "#1a544b",
			"icon_id": 0,
			"target_id": 419220,
			"special": "",
			"anchor_uname": "神奇陆夫人",
			"anchor_roomid": 115,
			"medal_level": 22,
			"medal_name": "陆夫人"
		},
		"user_info": {
			"user_level": 20,
			"level_color": "#61c05a",
			"is_vip": 0,
			"is_svip": 0,
			"is_main_vip": 1,
			"title": "0",
			"uname": "悲剧携带者",
			"face": "http://i2.hdslb.com/bfs/face/350aebb461ca8215b70c4cb4c1e8061ccb6d7db1.jpg",
			"manager": 0,
			"face_frame": "http://i0.hdslb.com/bfs/live/78e8a800e97403f1137c0c1b5029648c390be390.png",
			"guard_level": 3
		},
		"id": "2576342",
		"message_jpn": "",
		"time": 60,
		"rate": 1000,
		"background_image": "https://i0.hdslb.com/bfs/live/a712efa5c6ebc67bafbe8352d3e74b820a00c13e.png",
		"background_icon": "",
		"background_price_color": "#7497CD",
		"token": "1B6E22FC",
		"gift": {
			"num": 1,
			"gift_id": 12000,
			"gift_name": "醒目留言"
		},
		"price": 30,
		"message": "夫人，看你航天伴着好运来刷牛场，出了基德的财运，但是为什么mf才22（上限40啊",
		"background_color": "#EDF5FF",
		"background_bottom_color": "#2A60B2",
		"ts": 1635766505,
		"start_time": 1635766505,
		"end_time": 1635766565
	}
*/
type SuperChatMessageJPN struct {
	UID       string `json:"uid"`
	IsRanked  int    `json:"is_ranked"`
	MedalInfo struct {
		MedalColor   string `json:"medal_color"`
		IconID       int64  `json:"icon_id"`
		TargetID     int64  `json:"target_id"`
		Special      string `json:"special"`
		AnchorUname  string `json:"anchor_uname"`
		AnchorRoomID int    `json:"anchor_roomid"`
		MedalLevel   int    `json:"medal_level"`
		MedalName    string `json:"medal_name"`
	} `json:"medal_info"`
	UserInfo struct {
		UserLevel  int    `json:"user_level"`
		LevelColor string `json:"level_color"`
		IsVip      int    `json:"is_vip"`
		IsSvip     int    `json:"is_svip"`
		IsMainVip  int    `json:"is_main_vip"`
		Title      string `json:"title"`
		Uname      string `json:"uname"`
		Face       string `json:"face"`
		Manager    int    `json:"manager"`
		FaceFrame  string `json:"face_frame"`
		GuardLevel int    `json:"guard_level"`
	} `json:"user_info"`
	ID                   string `json:"id"`
	MessageJpn           string `json:"message_jpn"`
	Time                 int64  `json:"time"`
	Rate                 int    `json:"rate"`
	BackgroundImage      string `json:"background_image"`
	BackgroundIcon       string `json:"background_icon"`
	BackgroundPriceColor string `json:"background_price_color"`
	Token                string `json:"token"`
	Gift                 struct {
		Num      int    `json:"num"`
		GiftID   int64  `json:"gift_id"`
		GiftName string `json:"gift_name"`
	} `json:"gift"`
	Price                 int    `json:"price"`
	Message               string `json:"message"`
	BackgroundColor       string `json:"background_color"`
	BackgroundBottomColor string `json:"background_bottom_color"`
	TS                    int64  `json:"ts"`
	StartTime             int64  `json:"start_time"`
	EndTime               int64  `json:"end_time"`
}

func (SuperChatMessageJPN) Cmd() string {
	return CmdSuperChatMessageJPN
}

var _ = Register[SuperChatMessageJPN]()
