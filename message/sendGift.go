package message

// SendGift 部分any抓取不到有效的信息，只能先留着
/*
	{
		"discount_price": 100,
		"giftName": "牛哇牛哇",
		"gold": 0,
		"guard_level": 0,
		"remain": 0,
		"silver": 0,
		"super_gift_num": 4,
		"top_list": null,
		"biz_source": "xlottery-anchor",
		"combo_total_coin": 400,
		"giftType": 0,
		"magnification": 1,
		"medal_info": {
			"medal_name": "吉祥草",
			"special": "",
			"anchor_roomid": 0,
			"anchor_uname": "",
			"medal_color_border": 6067854,
			"medal_color_end": 6067854,
			"medal_color_start": 6067854,
			"medal_level": 4,
			"target_id": 2920960,
			"guard_level": 0,
			"icon_id": 0,
			"is_lighted": 1,
			"medal_color": 6067854
		},
		"name_color": "",
		"price": 100,
		"super": 0,
		"tag_image": "",
		"total_coin": 100,
		"uname": "余烬的圆舞曲",
		"blind_gift": null,
		"rnd": "1635849011111500002",
		"action": "投喂",
		"broadcast_id": 0,
		"effect": 0,
		"giftId": 31039,
		"is_special_batch": 0,
		"tid": "1635849011111500002",
		"batch_combo_id": "batch:gift:combo_id:9184735:2920960:31039:1635849007.7560",
		"float_sc_resource_id": 0,
		"original_gift_name": "",
		"batch_combo_send": null,
		"is_first": false,
		"num": 1,
		"rcost": 189509940,
		"uid": 9184735,
		"beatId": "",
		"combo_send": null,
		"combo_stay_time": 3,
		"dmscore": 32,
		"svga_block": 0,
		"timestamp": 1635849011,
		"coin_type": "gold",
		"combo_resources_id": 1,
		"crit_prob": 0,
		"demarcation": 1,
		"draw": 0,
		"effect_block": 0,
		"face": "http://i1.hdslb.com/bfs/face/80cd97607e8ab30acc768047db37a17c9270ec76.jpg",
		"send_master": null,
		"super_batch_gift_num": 4
	}
*/
type SendGift struct {
	Action         string `json:"action"`
	BatchComboID   string `json:"batch_combo_id"`
	BatchComboSend struct {
		Action        string `json:"action"`
		BatchComboID  string `json:"batch_combo_id"`
		BatchComboNum int    `json:"batch_combo_num"`
		BlindGift     any    `json:"blind_gift"`
		GiftID        int64  `json:"gift_id"`
		GiftName      string `json:"gift_name"`
		GiftNum       int    `json:"gift_num"`
		SendMaster    any    `json:"send_master"`
		UID           int    `json:"uid"`
		Uname         string `json:"uname"`
	} `json:"batch_combo_send"`
	BeatID           string `json:"beatId"`
	BizSource        string `json:"biz_source"`
	BlindGift        any    `json:"blind_gift"`
	BroadcastID      int64  `json:"broadcast_id"`
	CoinType         string `json:"coin_type"`
	ComboResourcesID int64  `json:"combo_resources_id"`
	ComboSend        struct {
		Action     string `json:"action"`
		ComboID    string `json:"combo_id"`
		ComboNum   int    `json:"combo_num"`
		GiftID     int64  `json:"gift_id"`
		GiftName   string `json:"gift_name"`
		GiftNum    int    `json:"gift_num"`
		SendMaster any    `json:"send_master"`
		UID        int64  `json:"uid"`
		Uname      string `json:"uname"`
	} `json:"combo_send"`
	ComboStayTime     int64   `json:"combo_stay_time"`
	ComboTotalCoin    int     `json:"combo_total_coin"`
	CritProb          int     `json:"crit_prob"`
	Demarcation       int     `json:"demarcation"`
	DiscountPrice     int     `json:"discount_price"`
	Dmscore           int     `json:"dmscore"`
	Draw              int     `json:"draw"`
	Effect            int     `json:"effect"`
	EffectBlock       int     `json:"effect_block"`
	Face              string  `json:"face"`
	FloatScResourceID int64   `json:"float_sc_resource_id"`
	GiftID            int64   `json:"giftId"`
	GiftName          string  `json:"giftName"`
	GiftType          int     `json:"giftType"`
	Gold              int     `json:"gold"`
	GuardLevel        int     `json:"guard_level"`
	IsFirst           bool    `json:"is_first"`
	IsSpecialBatch    int     `json:"is_special_batch"`
	Magnification     float64 `json:"magnification"`
	MedalInfo         struct {
		AnchorRoomID     int    `json:"anchor_roomid"`
		AnchorUname      string `json:"anchor_uname"`
		GuardLevel       int    `json:"guard_level"`
		IconID           int64  `json:"icon_id"`
		IsLighted        int    `json:"is_lighted"`
		MedalColor       int    `json:"medal_color"`
		MedalColorBorder int64  `json:"medal_color_border"`
		MedalColorEnd    int64  `json:"medal_color_end"`
		MedalColorStart  int64  `json:"medal_color_start"`
		MedalLevel       int    `json:"medal_level"`
		MedalName        string `json:"medal_name"`
		Special          string `json:"special"`
		TargetID         int    `json:"target_id"`
	} `json:"medal_info"`
	NameColor         string `json:"name_color"`
	Num               int    `json:"num"`
	OriginalGiftName  string `json:"original_gift_name"`
	Price             int    `json:"price"`
	Rcost             int    `json:"rcost"`
	Remain            int    `json:"remain"`
	Rnd               string `json:"rnd"`
	SendMaster        any    `json:"send_master"`
	Silver            int    `json:"silver"`
	Super             int    `json:"super"`
	SuperBatchGiftNum int    `json:"super_batch_gift_num"`
	SuperGiftNum      int    `json:"super_gift_num"`
	SvgaBlock         int    `json:"svga_block"`
	TagImage          string `json:"tag_image"`
	TID               string `json:"tid"`
	Timestamp         int64  `json:"timestamp"`
	TopList           any    `json:"top_list"`
	TotalCoin         int    `json:"total_coin"`
	UID               int64  `json:"uid"`
	Uname             string `json:"uname"`
}

func (SendGift) Cmd() string {
	return CmdSendGift
}

var _ = Register[SendGift]()

/*
	{
		"action": "投喂",
		"batch_combo_id": "batch:gift:combo_id:4109296:67141:31036:1717323308.7334",
		"batch_combo_num": 2,
		"coin_type": "gold",
		"combo_id": "gift:combo_id:4109296:67141:31036:1717323308.7321",
		"combo_num": 2,
		"combo_total_coin": 200,
		"dmscore": 714,
		"gift_id": 31036,
		"gift_name": "小花花",
		"gift_num": 0,
		"group_medal": null,
		"is_join_receiver": false,
		"is_naming": false,
		"is_show": 1,
		"medal_info": {
			"anchor_roomid": 0,
			"anchor_uname": "",
			"guard_level": 3,
			"icon_id": 0,
			"is_lighted": 1,
			"medal_color": 1725515,
			"medal_color_border": 6809855,
			"medal_color_end": 5414290,
			"medal_color_start": 1725515,
			"medal_level": 23,
			"medal_name": "C酱",
			"special": "",
			"target_id": 67141
		},
		"name_color": "#00D1F1",
		"r_uname": "C酱です",
		"receive_user_info": {
			"uid": 67141,
			"uname": "C酱です"
		},
		"receiver_uinfo": {
			"base": {
				"face": "https://i1.hdslb.com/bfs/face/361274190d08a36ed12c58b55dd4063cde3391eb.jpg",
				"is_mystery": false,
				"name": "C酱です",
				"name_color": 0,
				"name_color_str": "",
				"official_info": {
					"desc": "代表作《【生化危机7DLC】伊森必须死让伊森哭着找妈妈的通关攻略视频》",
					"role": 1,
					"title": "2021年度巅峰主播、bilibili 知名游戏UP主、直播高能主播",
					"type": 0
				},
				"origin_info": {
					"face": "https://i1.hdslb.com/bfs/face/361274190d08a36ed12c58b55dd4063cde3391eb.jpg",
					"name": "C酱です"
				},
				"risk_ctrl_info": {
					"face": "https://i1.hdslb.com/bfs/face/361274190d08a36ed12c58b55dd4063cde3391eb.jpg",
					"name": "C酱です"
				}
			},
			"guard": null,
			"guard_leader": null,
			"medal": null,
			"title": null,
			"uhead_frame": null,
			"uid": 67141,
			"wealth": null
		},
		"ruid": 67141,
		"send_master": null,
		"sender_uinfo": {
			"base": {
				"face": "https://i1.hdslb.com/bfs/face/d59a78f41c20a9efec74e5b108642651397e786f.jpg",
				"is_mystery": false,
				"name": "小鳄鱼鱼鱼c",
				"name_color": 0,
				"name_color_str": "",
				"official_info": {
					"desc": "",
					"role": 0,
					"title": "",
					"type": -1
				},
				"origin_info": {
					"face": "https://i1.hdslb.com/bfs/face/d59a78f41c20a9efec74e5b108642651397e786f.jpg",
					"name": "小鳄鱼鱼鱼c"
				},
				"risk_ctrl_info": {
					"face": "https://i1.hdslb.com/bfs/face/d59a78f41c20a9efec74e5b108642651397e786f.jpg",
					"name": "小鳄鱼鱼鱼c"
				}
			},
			"guard": null,
			"guard_leader": null,
			"medal": null,
			"title": null,
			"uhead_frame": null,
			"uid": 4109296,
			"wealth": null
		},
		"total_num": 2,
		"uid": 4109296,
		"uname": "小鳄鱼鱼鱼c",
		"wealth_level": 26
	}
*/
type ComboSend struct {
	Action         string `json:"action"`
	BatchComboID   string `json:"batch_combo_id"`
	BatchComboNum  int    `json:"batch_combo_num"`
	CoinType       string `json:"coin_type"`
	ComboID        string `json:"combo_id"`
	ComboNum       int    `json:"combo_num"`
	ComboTotalCoin int    `json:"combo_total_coin"`
	Dmscore        int    `json:"dmscore"`
	GiftID         int    `json:"gift_id"`
	GiftName       string `json:"gift_name"`
	GiftNum        int    `json:"gift_num"`
	GroupMedal     any    `json:"group_medal"`
	IsJoinReceiver bool   `json:"is_join_receiver"`
	IsNaming       bool   `json:"is_naming"`
	IsShow         int    `json:"is_show"`
	MedalInfo      struct {
		AnchorRoomid     int    `json:"anchor_roomid"`
		AnchorUname      string `json:"anchor_uname"`
		GuardLevel       int    `json:"guard_level"`
		IconID           int    `json:"icon_id"`
		IsLighted        int    `json:"is_lighted"`
		MedalColor       int    `json:"medal_color"`
		MedalColorBorder int    `json:"medal_color_border"`
		MedalColorEnd    int    `json:"medal_color_end"`
		MedalColorStart  int    `json:"medal_color_start"`
		MedalLevel       int    `json:"medal_level"`
		MedalName        string `json:"medal_name"`
		Special          string `json:"special"`
		TargetID         int    `json:"target_id"`
	} `json:"medal_info"`
	NameColor       string `json:"name_color"`
	RUname          string `json:"r_uname"`
	ReceiveUserInfo struct {
		UID   int    `json:"uid"`
		Uname string `json:"uname"`
	} `json:"receive_user_info"`
	ReceiverUinfo struct {
		Base struct {
			Face         string `json:"face"`
			IsMystery    bool   `json:"is_mystery"`
			Name         string `json:"name"`
			NameColor    int    `json:"name_color"`
			NameColorStr string `json:"name_color_str"`
			OfficialInfo struct {
				Desc  string `json:"desc"`
				Role  int    `json:"role"`
				Title string `json:"title"`
				Type  int    `json:"type"`
			} `json:"official_info"`
			OriginInfo struct {
				Face string `json:"face"`
				Name string `json:"name"`
			} `json:"origin_info"`
			RiskCtrlInfo struct {
				Face string `json:"face"`
				Name string `json:"name"`
			} `json:"risk_ctrl_info"`
		} `json:"base"`
		Guard       any `json:"guard"`
		GuardLeader any `json:"guard_leader"`
		Medal       any `json:"medal"`
		Title       any `json:"title"`
		UheadFrame  any `json:"uhead_frame"`
		UID         int `json:"uid"`
		Wealth      any `json:"wealth"`
	} `json:"receiver_uinfo"`
	Ruid        int `json:"ruid"`
	SendMaster  any `json:"send_master"`
	SenderUinfo struct {
		Base struct {
			Face         string `json:"face"`
			IsMystery    bool   `json:"is_mystery"`
			Name         string `json:"name"`
			NameColor    int    `json:"name_color"`
			NameColorStr string `json:"name_color_str"`
			OfficialInfo struct {
				Desc  string `json:"desc"`
				Role  int    `json:"role"`
				Title string `json:"title"`
				Type  int    `json:"type"`
			} `json:"official_info"`
			OriginInfo struct {
				Face string `json:"face"`
				Name string `json:"name"`
			} `json:"origin_info"`
			RiskCtrlInfo struct {
				Face string `json:"face"`
				Name string `json:"name"`
			} `json:"risk_ctrl_info"`
		} `json:"base"`
		Guard       any `json:"guard"`
		GuardLeader any `json:"guard_leader"`
		Medal       any `json:"medal"`
		Title       any `json:"title"`
		UheadFrame  any `json:"uhead_frame"`
		UID         int `json:"uid"`
		Wealth      any `json:"wealth"`
	} `json:"sender_uinfo"`
	TotalNum    int    `json:"total_num"`
	UID         int    `json:"uid"`
	Uname       string `json:"uname"`
	WealthLevel int    `json:"wealth_level"`
}

func (ComboSend) Cmd() string {
	return CmdComboSend
}

var _ = Register[ComboSend]()
