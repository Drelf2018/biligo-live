package message

const CmdEntryEffectMustReceive = "ENTRY_EFFECT_MUST_RECEIVE"

type EntryEffectMustReceive struct {
	ID                   int    `json:"id"`                 // 4
	UID                  int    `json:"uid"`                // 188888131
	TargetID             int    `json:"target_id"`          // 434334701
	MockEffect           int    `json:"mock_effect"`        // 0
	Face                 string `json:"face"`               // "https://i1.hdslb.com/bfs/face/86faab4844dd2c45870fdafa8f2c9ce7be3e999f.jpg"
	PrivilegeType        int    `json:"privilege_type"`     // 3
	CopyWriting          string `json:"copy_writing"`       // "<%脆鲨12138%> 来了"
	CopyColor            string `json:"copy_color"`         // "#ffffff"
	HighlightColor       string `json:"highlight_color"`    // "#E6FF00"
	Priority             int    `json:"priority"`           // 1
	BasemapURL           string `json:"basemap_url"`        // ""
	ShowAvatar           int    `json:"show_avatar"`        // 0
	EffectiveTime        int    `json:"effective_time"`     // 0
	WebBasemapURL        string `json:"web_basemap_url"`    // "https://i0.hdslb.com/bfs/live/mlive/75e5cc5dc03c3ea8f2fd40d6833f9f11a8d26b88.png"
	WebEffectiveTime     int    `json:"web_effective_time"` // 2
	WebEffectClose       int    `json:"web_effect_close"`   // 0
	WebCloseTime         int    `json:"web_close_time"`     // 0
	Business             int    `json:"business"`           // 1
	CopyWritingV2        string `json:"copy_writing_v2"`    // "<%脆鲨12138%> 来了"
	IconList             []any  `json:"icon_list"`
	MaxDelayTime         int    `json:"max_delay_time"`          // 7
	TriggerTime          int64  `json:"trigger_time"`            // 1735301622220277011
	Identities           int    `json:"identities"`              // 6
	EffectSilentTime     int    `json:"effect_silent_time"`      // 0
	EffectiveTimeNew     int    `json:"effective_time_new"`      // 0
	WebDynamicURLWebp    string `json:"web_dynamic_url_webp"`    // ""
	WebDynamicURLApng    string `json:"web_dynamic_url_apng"`    // ""
	MobileDynamicURLWebp string `json:"mobile_dynamic_url_webp"` // ""
	WealthyInfo          any    `json:"wealthy_info"`
	NewStyle             int    `json:"new_style"`  // 0
	IsMystery            bool   `json:"is_mystery"` // false
	Uinfo                struct {
		UID  int `json:"uid"` // 188888131
		Base struct {
			Name         string `json:"name"`       // "脆鲨12138"
			Face         string `json:"face"`       // "https://i1.hdslb.com/bfs/face/86faab4844dd2c45870fdafa8f2c9ce7be3e999f.jpg"
			NameColor    int    `json:"name_color"` // 0
			IsMystery    bool   `json:"is_mystery"` // false
			RiskCtrlInfo any    `json:"risk_ctrl_info"`
			OriginInfo   any    `json:"origin_info"`
			OfficialInfo any    `json:"official_info"`
			NameColorStr string `json:"name_color_str"` // "#00D1F1"
		} `json:"base"`
		Medal struct {
			Name               string `json:"name"`                  // "脆鲨"
			Level              int    `json:"level"`                 // 27
			ColorStart         int    `json:"color_start"`           // 398668
			ColorEnd           int    `json:"color_end"`             // 6850801
			ColorBorder        int    `json:"color_border"`          // 6809855
			Color              int    `json:"color"`                 // 398668
			ID                 int    `json:"id"`                    // 193893
			Typ                int    `json:"typ"`                   // 0
			IsLight            int    `json:"is_light"`              // 1
			Ruid               int    `json:"ruid"`                  // 434334701
			GuardLevel         int    `json:"guard_level"`           // 3
			Score              int    `json:"score"`                 // 50131762
			GuardIcon          string `json:"guard_icon"`            // "https://i0.hdslb.com/bfs/live/143f5ec3003b4080d1b5f817a9efdca46d631945.png"
			HonorIcon          string `json:"honor_icon"`            // ""
			V2MedalColorStart  string `json:"v2_medal_color_start"`  // "#4775EFCC"
			V2MedalColorEnd    string `json:"v2_medal_color_end"`    // "#4775EFCC"
			V2MedalColorBorder string `json:"v2_medal_color_border"` // "#58A1F8FF"
			V2MedalColorText   string `json:"v2_medal_color_text"`   // "#FFFFFFFF"
			V2MedalColorLevel  string `json:"v2_medal_color_level"`  // "#000B7099"
			UserReceiveCount   int    `json:"user_receive_count"`    // 0
		} `json:"medal"`
		Wealth struct {
			Level     int    `json:"level"`       // 34
			DmIconKey string `json:"dm_icon_key"` // ""
		} `json:"wealth"`
		Title any `json:"title"`
		Guard struct {
			Level      int    `json:"level"`       // 3
			ExpiredStr string `json:"expired_str"` // "2025-01-24 23:59:59"
		} `json:"guard"`
		UheadFrame  any `json:"uhead_frame"`
		GuardLeader any `json:"guard_leader"`
	} `json:"uinfo"`
	FullCartoonID   int `json:"full_cartoon_id"` // 1805
	PriorityLevel   int `json:"priority_level"`  // 0
	WealthStyleInfo any `json:"wealth_style_info"`
}

func (EntryEffectMustReceive) Cmd() string {
	return CmdEntryEffectMustReceive
}
