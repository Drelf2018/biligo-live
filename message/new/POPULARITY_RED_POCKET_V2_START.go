package message

const CmdPopularityRedPocketV2Start = "POPULARITY_RED_POCKET_V2_START"

type PopularityRedPocketV2Start struct {
	LotID           int    `json:"lot_id"`           // 25109135
	SenderUID       int    `json:"sender_uid"`       // 26738680
	SenderName      string `json:"sender_name"`      // "aimeAC布鲁贝里"
	SenderFace      string `json:"sender_face"`      // "https://i1.hdslb.com/bfs/face/f6c63e4d1d2e66c00ea7f61ac169aad667366a9c.jpg"
	JoinRequirement int    `json:"join_requirement"` // 1
	Danmu           string `json:"danmu"`            // "老板大气！点点红包抽礼物"
	CurrentTime     int    `json:"current_time"`     // 1735301880
	StartTime       int    `json:"start_time"`       // 1735301879
	EndTime         int    `json:"end_time"`         // 1735302059
	LastTime        int    `json:"last_time"`        // 180
	RemoveTime      int    `json:"remove_time"`      // 1735302074
	ReplaceTime     int    `json:"replace_time"`     // 1735302069
	LotStatus       int    `json:"lot_status"`       // 1
	H5URL           string `json:"h5_url"`           // "https://live.bilibili.com/p/html/live-app-red-envelope/popularity.html?is_live_half_webview=1&hybrid_half_ui=1,5,100p,100p,000000,0,50,0,0,1;2,5,100p,100p,000000,0,50,0,0,1;3,5,100p,100p,000000,0,50,0,0,1;4,5,100p,100p,000000,0,50,0,0,1;5,5,100p,100p,000000,0,50,0,0,1;6,5,100p,100p,000000,0,50,0,0,1;7,5,100p,100p,000000,0,50,0,0,1;8,5,100p,100p,000000,0,50,0,0,1&hybrid_rotate_d=1&hybrid_biz=popularityRedPacket&lotteryId=25109135"
	UserStatus      int    `json:"user_status"`      // 2
	Awards          []struct {
		GiftID   int    `json:"gift_id"`   // 31213
		GiftName string `json:"gift_name"` // "这个好诶"
		GiftPic  string `json:"gift_pic"`  // "https://s1.hdslb.com/bfs/live/9260c680959428c45b3a2742e42ea7ae75e457ef.png"
		Num      int    `json:"num"`       // 3
	} `json:"awards"`
	LotConfigID int  `json:"lot_config_id"` // 4
	TotalPrice  int  `json:"total_price"`   // 8000
	WaitNum     int  `json:"wait_num"`      // 0
	WaitNumV2   int  `json:"wait_num_v2"`   // 0
	IsMystery   bool `json:"is_mystery"`    // false
	RpType      int  `json:"rp_type"`       // 0
	SenderUinfo struct {
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
		Medal       any `json:"medal"`
		Wealth      any `json:"wealth"`
		Title       any `json:"title"`
		Guard       any `json:"guard"`
		UheadFrame  any `json:"uhead_frame"`
		GuardLeader any `json:"guard_leader"`
	} `json:"sender_uinfo"`
	IconURL          string `json:"icon_url"`           // ""
	AnimationIconURL string `json:"animation_icon_url"` // ""
	RpGuardInfo      any    `json:"rp_guard_info"`
}

func (PopularityRedPocketV2Start) Cmd() string {
	return CmdPopularityRedPocketV2Start
}
