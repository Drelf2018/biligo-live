package message

/*
	{
		"lot_status": 2,
		"url": "https://live.bilibili.com/p/html/live-lottery/anchor-join.html?is_live_half_webview=1&hybrid_biz=live-lottery-anchor&hybrid_half_ui=1,5,100p,100p,000000,0,30,0,0,1;2,5,100p,100p,000000,0,30,0,0,1;3,5,100p,100p,000000,0,30,0,0,1;4,5,100p,100p,000000,0,30,0,0,1;5,5,100p,100p,000000,0,30,0,0,1;6,5,100p,100p,000000,0,30,0,0,1;7,5,100p,100p,000000,0,30,0,0,1;8,5,100p,100p,000000,0,30,0,0,1",
		"web_url": "https://live.bilibili.com/p/html/live-lottery/anchor-join.html",
		"award_image": "",
		"award_name": "648元现金红包",
		"award_num": 1,
		"award_users": [{
			"uname": "咲友",
			"face": "http://i0.hdslb.com/bfs/face/09b505910990d61a0777984f8a142b8e70485987.jpg",
			"level": 21,
			"color": 5805790,
			"uid": 29038770
		}],
		"id": 1890708
	}
*/
type AnchorLotAward struct {
	LotStatus  int    `json:"lot_status"`
	Url        string `json:"url"`
	WebUrl     string `json:"web_url"`
	AwardImage string `json:"award_image"`
	AwardName  string `json:"award_name"`
	AwardNum   int    `json:"award_num"`
	AwardUsers []struct {
		Uname string `json:"uname"`
		Face  string `json:"face"`
		Level int    `json:"level"`
		Color int64  `json:"color"`
		UID   int64  `json:"uid"`
	} `json:"award_users"`
	ID int64 `json:"id"`
}

func (AnchorLotAward) Cmd() string {
	return CmdAnchorLotAward
}

var _ = Register[AnchorLotAward]()

/*
	{
		"id": 1890708,
		"reject_reason": "",
		"status": 4,
		"uid": 2920960
	}
*/
type AnchorLotCheckStatus struct {
	ID           int64  `json:"id"`
	RejectReason string `json:"reject_reason"`
	Status       int    `json:"status"`
	UID          int64  `json:"uid"`
}

func (AnchorLotCheckStatus) Cmd() string {
	return CmdAnchorLotCheckStatus
}

var _ = Register[AnchorLotCheckStatus]()

/*
	{
		"id": 6793387
	}
*/
type AnchorLotEnd struct {
	ID int `json:"id"`
}

func (AnchorLotEnd) Cmd() string {
	return CmdAnchorLotEnd
}

var _ = Register[AnchorLotEnd]()

/*
	{
		"max_time": 600,
		"danmu": "好耶！",
		"gift_num": 1,
		"join_type": 0,
		"award_image": "",
		"gift_price": 0,
		"gift_id": 0,
		"gift_name": "",
		"goods_id": -99998,
		"room_id": 732602,
		"time": 599,
		"url": "https://live.bilibili.com/p/html/live-lottery/anchor-join.html?is_live_half_webview=1&hybrid_biz=live-lottery-anchor&hybrid_half_ui=1,5,100p,100p,000000,0,30,0,0,1;2,5,100p,100p,000000,0,30,0,0,1;3,5,100p,100p,000000,0,30,0,0,1;4,5,100p,100p,000000,0,30,0,0,1;5,5,100p,100p,000000,0,30,0,0,1;6,5,100p,100p,000000,0,30,0,0,1;7,5,100p,100p,000000,0,30,0,0,1;8,5,100p,100p,000000,0,30,0,0,1",
		"cur_gift_num": 0,
		"current_time": 1635849356,
		"lot_status": 0,
		"require_type": 2,
		"web_url": "https://live.bilibili.com/p/html/live-lottery/anchor-join.html",
		"goaway_time": 180,
		"is_broadcast": 1,
		"require_value": 1,
		"show_panel": 1,
		"status": 1,
		"id": 1890708,
		"require_text": "当前主播粉丝勋章至少1级",
		"award_num": 1,
		"asset_icon": "https://i0.hdslb.com/bfs/live/627ee2d9e71c682810e7dc4400d5ae2713442c02.png",
		"award_name": "648元现金红包",
		"send_gift_ensure": 0
	}
*/
type AnchorLotStart struct {
	MaxTime        int    `json:"max_time"`
	Danmu          string `json:"danmu"`
	GiftNum        int    `json:"gift_num"`
	JoinType       int    `json:"join_type"`
	AwardImage     string `json:"award_image"`
	GiftPrice      int    `json:"gift_price"`
	GiftID         int64  `json:"gift_id"`
	GiftName       string `json:"gift_name"`
	GoodsID        int64  `json:"goods_id"`
	RoomID         int64  `json:"room_id"`
	Time           int64  `json:"time"`
	Url            string `json:"url"`
	CurGiftNum     int    `json:"cur_gift_num"`
	CurrentTime    int64  `json:"current_time"`
	LotStatus      int    `json:"lot_status"`
	RequireType    int    `json:"require_type"`
	WebUrl         string `json:"web_url"`
	GoawayTime     int    `json:"goaway_time"`
	IsBroadcast    int    `json:"is_broadcast"`
	RequireValue   int    `json:"require_value"`
	ShowPanel      int    `json:"show_panel"`
	Status         int    `json:"status"`
	ID             int64  `json:"id"`
	RequireText    string `json:"require_text"`
	AwardNum       int    `json:"award_num"`
	AssetIcon      string `json:"asset_icon"`
	AwardName      string `json:"award_name"`
	SendGiftEnsure int    `json:"send_gift_ensure"`
}

func (AnchorLotStart) Cmd() string {
	return CmdAnchorLotStart
}

var _ = Register[AnchorLotStart]()
