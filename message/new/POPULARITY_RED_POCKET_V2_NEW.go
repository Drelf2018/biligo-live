package message

const CmdPopularityRedPocketV2New = "POPULARITY_RED_POCKET_V2_NEW"

type PopularityRedPocketV2New struct {
	LotID       int    `json:"lot_id"`       // 25109135
	StartTime   int    `json:"start_time"`   // 1735301879
	CurrentTime int    `json:"current_time"` // 1735301879
	WaitNum     int    `json:"wait_num"`     // 0
	WaitNumV2   int    `json:"wait_num_v2"`  // 0
	Uname       string `json:"uname"`        // "aimeAC布鲁贝里"
	UID         int    `json:"uid"`          // 26738680
	Action      string `json:"action"`       // "送出"
	Num         int    `json:"num"`          // 1
	GiftName    string `json:"gift_name"`    // "红包"
	GiftID      int    `json:"gift_id"`      // 13000
	Price       int    `json:"price"`        // 100
	NameColor   string `json:"name_color"`   // ""
	MedalInfo   struct {
		TargetID         int    `json:"target_id"`          // 434334701
		Special          string `json:"special"`            // ""
		IconID           int    `json:"icon_id"`            // 0
		AnchorUname      string `json:"anchor_uname"`       // ""
		AnchorRoomid     int    `json:"anchor_roomid"`      // 0
		MedalLevel       int    `json:"medal_level"`        // 27
		MedalName        string `json:"medal_name"`         // "脆鲨"
		MedalColor       int    `json:"medal_color"`        // 398668
		MedalColorStart  int    `json:"medal_color_start"`  // 398668
		MedalColorEnd    int    `json:"medal_color_end"`    // 6850801
		MedalColorBorder int    `json:"medal_color_border"` // 398668
		IsLighted        int    `json:"is_lighted"`         // 1
		GuardLevel       int    `json:"guard_level"`        // 0
	} `json:"medal_info"`
	WealthLevel int  `json:"wealth_level"` // 35
	GroupMedal  any  `json:"group_medal"`
	IsMystery   bool `json:"is_mystery"` // false
	SenderInfo  struct {
		UID  int `json:"uid"` // 26738680
		Base struct {
			Name         string `json:"name"`       // "aimeAC布鲁贝里"
			Face         string `json:"face"`       // "https://i1.hdslb.com/bfs/face/f6c63e4d1d2e66c00ea7f61ac169aad667366a9c.jpg"
			NameColor    int    `json:"name_color"` // 0
			IsMystery    bool   `json:"is_mystery"` // false
			RiskCtrlInfo any    `json:"risk_ctrl_info"`
			OriginInfo   struct {
				Name string `json:"name"` // "aimeAC布鲁贝里"
				Face string `json:"face"` // "https://i1.hdslb.com/bfs/face/f6c63e4d1d2e66c00ea7f61ac169aad667366a9c.jpg"
			} `json:"origin_info"`
			OfficialInfo struct {
				Role  int    `json:"role"`  // 0
				Title string `json:"title"` // ""
				Desc  string `json:"desc"`  // ""
				Type  int    `json:"type"`  // -1
			} `json:"official_info"`
			NameColorStr string `json:"name_color_str"` // ""
		} `json:"base"`
		Medal struct {
			Name               string `json:"name"`                  // "脆鲨"
			Level              int    `json:"level"`                 // 27
			ColorStart         int    `json:"color_start"`           // 398668
			ColorEnd           int    `json:"color_end"`             // 6850801
			ColorBorder        int    `json:"color_border"`          // 398668
			Color              int    `json:"color"`                 // 398668
			ID                 int    `json:"id"`                    // 0
			Typ                int    `json:"typ"`                   // 0
			IsLight            int    `json:"is_light"`              // 1
			Ruid               int    `json:"ruid"`                  // 434334701
			GuardLevel         int    `json:"guard_level"`           // 0
			Score              int    `json:"score"`                 // 50117604
			GuardIcon          string `json:"guard_icon"`            // ""
			HonorIcon          string `json:"honor_icon"`            // ""
			V2MedalColorStart  string `json:"v2_medal_color_start"`  // "#4775EFCC"
			V2MedalColorEnd    string `json:"v2_medal_color_end"`    // "#4775EFCC"
			V2MedalColorBorder string `json:"v2_medal_color_border"` // "#58A1F8FF"
			V2MedalColorText   string `json:"v2_medal_color_text"`   // "#FFFFFFFF"
			V2MedalColorLevel  string `json:"v2_medal_color_level"`  // "#000B7099"
			UserReceiveCount   int    `json:"user_receive_count"`    // 0
		} `json:"medal"`
		Wealth struct {
			Level     int    `json:"level"`       // 35
			DmIconKey string `json:"dm_icon_key"` // ""
		} `json:"wealth"`
		Title any `json:"title"`
		Guard struct {
			Level      int    `json:"level"`       // 0
			ExpiredStr string `json:"expired_str"` // ""
		} `json:"guard"`
		UheadFrame  any `json:"uhead_frame"`
		GuardLeader any `json:"guard_leader"`
	} `json:"sender_info"`
	GiftIcon string `json:"gift_icon"` // ""
	RpType   int    `json:"rp_type"`   // 0
}

func (PopularityRedPocketV2New) Cmd() string {
	return CmdPopularityRedPocketV2New
}
