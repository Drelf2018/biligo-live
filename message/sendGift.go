package message

// SendGift 部分interface{}抓取不到有效的信息，只能先留着
//
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

var _ = Register(SendGift{})

type SendGift struct {
	Action         string `json:"action"`
	BatchComboID   string `json:"batch_combo_id"`
	BatchComboSend struct {
		Action        string      `json:"action"`
		BatchComboID  string      `json:"batch_combo_id"`
		BatchComboNum int         `json:"batch_combo_num"`
		BlindGift     interface{} `json:"blind_gift"`
		GiftID        int64       `json:"gift_id"`
		GiftName      string      `json:"gift_name"`
		GiftNum       int         `json:"gift_num"`
		SendMaster    interface{} `json:"send_master"`
		Uid           int         `json:"uid"`
		Uname         string      `json:"uname"`
	} `json:"batch_combo_send"`
	BeatID           string      `json:"beatId"`
	BizSource        string      `json:"biz_source"`
	BlindGift        interface{} `json:"blind_gift"`
	BroadcastID      int64       `json:"broadcast_id"`
	CoinType         string      `json:"coin_type"`
	ComboResourcesID int64       `json:"combo_resources_id"`
	ComboSend        struct {
		Action     string      `json:"action"`
		ComboID    string      `json:"combo_id"`
		ComboNum   int         `json:"combo_num"`
		GiftID     int64       `json:"gift_id"`
		GiftName   string      `json:"gift_name"`
		GiftNum    int         `json:"gift_num"`
		SendMaster interface{} `json:"send_master"`
		UID        int64       `json:"uid"`
		Uname      string      `json:"uname"`
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
		AnchorRoomid     int    `json:"anchor_roomid"`
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
	NameColor         string      `json:"name_color"`
	Num               int         `json:"num"`
	OriginalGiftName  string      `json:"original_gift_name"`
	Price             int         `json:"price"`
	Rcost             int         `json:"rcost"`
	Remain            int         `json:"remain"`
	Rnd               string      `json:"rnd"`
	SendMaster        interface{} `json:"send_master"`
	Silver            int         `json:"silver"`
	Super             int         `json:"super"`
	SuperBatchGiftNum int         `json:"super_batch_gift_num"`
	SuperGiftNum      int         `json:"super_gift_num"`
	SvgaBlock         int         `json:"svga_block"`
	TagImage          string      `json:"tag_image"`
	TID               string      `json:"tid"`
	Timestamp         int64       `json:"timestamp"`
	TopList           interface{} `json:"top_list"`
	TotalCoin         int         `json:"total_coin"`
	UID               int64       `json:"uid"`
	Uname             string      `json:"uname"`
}

func (SendGift) Cmd() string {
	return CmdSendGift
}

func (SendGift) Parse(raw []byte) (Msg, error) {
	return GetData[SendGift](raw)
}
