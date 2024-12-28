package message

const CmdUserToastMsgV2 = "USER_TOAST_MSG_V2"

type UserToastMsgV2 struct {
	SenderUinfo struct {
		UID  int `json:"uid"` // 23207109
		Base struct {
			Name string `json:"name"` // "干脆面猎人"
			Face string `json:"face"` // ""
		} `json:"base"`
	} `json:"sender_uinfo"`
	ReceiverUinfo struct {
		UID  int `json:"uid"` // 67141
		Base struct {
			Name string `json:"name"` // "傲慢的小肉包"
			Face string `json:"face"` // "https://i1.hdslb.com/bfs/face/361274190d08a36ed12c58b55dd4063cde3391eb.jpg"
		} `json:"base"`
	} `json:"receiver_uinfo"`
	GuardInfo struct {
		GuardLevel     int    `json:"guard_level"`      // 3
		RoleName       string `json:"role_name"`        // "舰长"
		RoomGuardCount int    `json:"room_guard_count"` // 708
		OpType         int    `json:"op_type"`          // 1
		StartTime      int    `json:"start_time"`       // 1735290731
		EndTime        int    `json:"end_time"`         // 1735290731
	} `json:"guard_info"`
	GroupGuardInfo any `json:"group_guard_info"`
	PayInfo        struct {
		PayflowID string `json:"payflow_id"` // "2412271711343342171092510"
		Price     int    `json:"price"`      // 138000
		Num       int    `json:"num"`        // 1
		Unit      string `json:"unit"`       // "月"
	} `json:"pay_info"`
	GiftInfo struct {
		GiftID int `json:"gift_id"` // 10003
	} `json:"gift_info"`
	EffectInfo struct {
		EffectID          int `json:"effect_id"`            // 397
		RoomEffectID      int `json:"room_effect_id"`       // 590
		FaceEffectID      int `json:"face_effect_id"`       // 44
		RoomGiftEffectID  int `json:"room_gift_effect_id"`  // 0
		RoomGroupEffectID int `json:"room_group_effect_id"` // 1337
	} `json:"effect_info"`
	ToastMsg string `json:"toast_msg"` // "<%干脆面猎人%> 在主播傲慢的小肉包的直播间开通了舰长，今天是TA陪伴主播的第1天"
	Option   struct {
		AnchorShow bool   `json:"anchor_show"` // true
		UserShow   bool   `json:"user_show"`   // true
		IsGroup    int    `json:"is_group"`    // 0
		IsShow     int    `json:"is_show"`     // 0
		Source     int    `json:"source"`      // 0
		SvgaBlock  int    `json:"svga_block"`  // 0
		Color      string `json:"color"`       // "#00D1F1"
	} `json:"option"`
}

func (UserToastMsgV2) Cmd() string {
	return CmdUserToastMsgV2
}
