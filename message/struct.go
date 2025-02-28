package message

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tidwall/gjson"
)

type AnchorLotAward struct {
	LotStatus  int    `json:"lot_status"`  // 2
	URL        string `json:"url"`         // "https://live.bilibili.com/p/html/live-lottery/anchor-join.html?is_live_half_webview=1&hybrid_biz=live-lottery-anchor&hybrid_half_ui=1,5,100p,100p,000000,0,30,0,0,1;2,5,100p,100p,000000,0,30,0,0,1;3,5,100p,100p,000000,0,30,0,0,1;4,5,100p,100p,000000,0,30,0,0,1;5,5,100p,100p,000000,0,30,0,0,1;6,5,100p,100p,000000,0,30,0,0,1;7,5,100p,100p,000000,0,30,0,0,1;8,5,100p,100p,000000,0,30,0,0,1"
	WebURL     string `json:"web_url"`     // "https://live.bilibili.com/p/html/live-lottery/anchor-join.html"
	AwardImage string `json:"award_image"` // ""
	AwardName  string `json:"award_name"`  // "648元现金红包"
	AwardNum   int    `json:"award_num"`   // 1
	AwardUsers []struct {
		Uname string `json:"uname"` // "咲友"
		Face  string `json:"face"`  // "http://i0.hdslb.com/bfs/face/09b505910990d61a0777984f8a142b8e70485987.jpg"
		Level int    `json:"level"` // 21
		Color int    `json:"color"` // 5805790
		UID   int    `json:"uid"`   // 29038770
	} `json:"award_users"`
	ID int `json:"id"` // 1890708
}

func (AnchorLotAward) Cmd() string {
	return CmdAnchorLotAward
}

type AnchorLotCheckStatus struct {
	ID           int    `json:"id"`            // 1890708
	RejectReason string `json:"reject_reason"` // ""
	Status       int    `json:"status"`        // 4
	UID          int    `json:"uid"`           // 2920960
}

func (AnchorLotCheckStatus) Cmd() string {
	return CmdAnchorLotCheckStatus
}

type AnchorLotEnd struct {
	ID int `json:"id"` // 6793387
}

func (AnchorLotEnd) Cmd() string {
	return CmdAnchorLotEnd
}

type AnchorLotStart struct {
	MaxTime        int    `json:"max_time"`         // 600
	Danmu          string `json:"danmu"`            // "好耶！"
	GiftNum        int    `json:"gift_num"`         // 1
	JoinType       int    `json:"join_type"`        // 0
	AwardImage     string `json:"award_image"`      // ""
	GiftPrice      int    `json:"gift_price"`       // 0
	GiftID         int    `json:"gift_id"`          // 0
	GiftName       string `json:"gift_name"`        // ""
	GoodsID        int    `json:"goods_id"`         // -99998
	RoomID         int    `json:"room_id"`          // 732602
	Time           int    `json:"time"`             // 599
	URL            string `json:"url"`              // "https://live.bilibili.com/p/html/live-lottery/anchor-join.html?is_live_half_webview=1&hybrid_biz=live-lottery-anchor&hybrid_half_ui=1,5,100p,100p,000000,0,30,0,0,1;2,5,100p,100p,000000,0,30,0,0,1;3,5,100p,100p,000000,0,30,0,0,1;4,5,100p,100p,000000,0,30,0,0,1;5,5,100p,100p,000000,0,30,0,0,1;6,5,100p,100p,000000,0,30,0,0,1;7,5,100p,100p,000000,0,30,0,0,1;8,5,100p,100p,000000,0,30,0,0,1"
	CurGiftNum     int    `json:"cur_gift_num"`     // 0
	CurrentTime    int    `json:"current_time"`     // 1635849356
	LotStatus      int    `json:"lot_status"`       // 0
	RequireType    int    `json:"require_type"`     // 2
	WebURL         string `json:"web_url"`          // "https://live.bilibili.com/p/html/live-lottery/anchor-join.html"
	GoawayTime     int    `json:"goaway_time"`      // 180
	IsBroadcast    int    `json:"is_broadcast"`     // 1
	RequireValue   int    `json:"require_value"`    // 1
	ShowPanel      int    `json:"show_panel"`       // 1
	Status         int    `json:"status"`           // 1
	ID             int    `json:"id"`               // 1890708
	RequireText    string `json:"require_text"`     // "当前主播粉丝勋章至少1级"
	AwardNum       int    `json:"award_num"`        // 1
	AssetIcon      string `json:"asset_icon"`       // "https://i0.hdslb.com/bfs/live/627ee2d9e71c682810e7dc4400d5ae2713442c02.png"
	AwardName      string `json:"award_name"`       // "648元现金红包"
	SendGiftEnsure int    `json:"send_gift_ensure"` // 0
}

func (AnchorLotStart) Cmd() string {
	return CmdAnchorLotStart
}

type AreaRankChanged struct {
	ConfID      int    `json:"conf_id"`       // 22
	RankName    string `json:"rank_name"`     // "网游航海"
	UID         int    `json:"uid"`           // 193584
	Rank        int    `json:"rank"`          // 21
	IconURLBlue string `json:"icon_url_blue"` // "https://i0.hdslb.com/bfs/live/18e2990a546d33368200f9058f3d9dbc4038eb5c.png"
	IconURLPink string `json:"icon_url_pink"` // "https://i0.hdslb.com/bfs/live/a6c490c36e88c7b191a04883a5ec15aed187a8f7.png"
	IconURLGrey string `json:"icon_url_grey"` // "https://i0.hdslb.com/bfs/live/cb7444b1faf1d785df6265bfdc1fcfc993419b76.png"
	ActionType  int    `json:"action_type"`   // 1
	Timestamp   int    `json:"timestamp"`     // 1717322899
	MsgID       string `json:"msg_id"`        // "008d250d-4d19-40d3-be51-f7d1f01be005"
	JumpURLLink string `json:"jump_url_link"` // "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=3&ruid=193584&conf_id=22&is_live_half_webview=1&hybrid_rotate_d=1&is_cling_player=1&hybrid_half_ui=1,3,100p,70p,f4eefa,0,30,100,0,0;2,2,375,100p,f4eefa,0,30,100,0,0;3,3,100p,70p,f4eefa,0,30,100,0,0;4,2,375,100p,f4eefa,0,30,100,0,0;5,3,100p,70p,f4eefa,0,30,100,0,0;6,3,100p,70p,f4eefa,0,30,100,0,0;7,3,100p,70p,f4eefa,0,30,100,0,0;8,3,100p,70p,f4eefa,0,30,100,0,0#/area-rank"
	JumpURLPc   string `json:"jump_url_pc"`   // "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=4&ruid=193584&conf_id=22&pc_ui=338,465,f4eefa,0#/area-rank"
	JumpURLPink string `json:"jump_url_pink"` // "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=1&ruid=193584&conf_id=22&is_live_half_webview=1&hybrid_rotate_d=1&hybrid_half_ui=1,3,100p,70p,ffffff,0,30,100,12,0;2,2,375,100p,ffffff,0,30,100,0,0;3,3,100p,70p,ffffff,0,30,100,12,0;4,2,375,100p,ffffff,0,30,100,0,0;5,3,100p,70p,ffffff,0,30,100,0,0;6,3,100p,70p,ffffff,0,30,100,0,0;7,3,100p,70p,ffffff,0,30,100,0,0;8,3,100p,70p,ffffff,0,30,100,0,0#/area-rank"
	JumpURLWeb  string `json:"jump_url_web"`  // "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=2&ruid=193584&conf_id=22#/area-rank"
}

func (AreaRankChanged) Cmd() string {
	return CmdAreaRankChanged
}

type ComboSend struct {
	Action         string `json:"action"`           // "投喂"
	BatchComboID   string `json:"batch_combo_id"`   // "batch:gift:combo_id:4109296:67141:31036:1717323308.7334"
	BatchComboNum  int    `json:"batch_combo_num"`  // 2
	CoinType       string `json:"coin_type"`        // "gold"
	ComboID        string `json:"combo_id"`         // "gift:combo_id:4109296:67141:31036:1717323308.7321"
	ComboNum       int    `json:"combo_num"`        // 2
	ComboTotalCoin int    `json:"combo_total_coin"` // 200
	Dmscore        int    `json:"dmscore"`          // 714
	GiftID         int    `json:"gift_id"`          // 31036
	GiftName       string `json:"gift_name"`        // "小花花"
	GiftNum        int    `json:"gift_num"`         // 0
	GroupMedal     any    `json:"group_medal"`
	IsJoinReceiver bool   `json:"is_join_receiver"` // false
	IsNaming       bool   `json:"is_naming"`        // false
	IsShow         int    `json:"is_show"`          // 1
	MedalInfo      struct {
		AnchorRoomID     int    `json:"anchor_roomid"`      // 0
		AnchorUname      string `json:"anchor_uname"`       // ""
		GuardLevel       int    `json:"guard_level"`        // 3
		IconID           int    `json:"icon_id"`            // 0
		IsLighted        int    `json:"is_lighted"`         // 1
		MedalColor       int    `json:"medal_color"`        // 1725515
		MedalColorBorder int    `json:"medal_color_border"` // 6809855
		MedalColorEnd    int    `json:"medal_color_end"`    // 5414290
		MedalColorStart  int    `json:"medal_color_start"`  // 1725515
		MedalLevel       int    `json:"medal_level"`        // 23
		MedalName        string `json:"medal_name"`         // "C酱"
		Special          string `json:"special"`            // ""
		TargetID         int    `json:"target_id"`          // 67141
	} `json:"medal_info"`
	NameColor       string `json:"name_color"` // "#00D1F1"
	RUname          string `json:"r_uname"`    // "C酱です"
	ReceiveUserInfo struct {
		UID   int    `json:"uid"`   // 67141
		Uname string `json:"uname"` // "C酱です"
	} `json:"receive_user_info"`
	ReceiverUinfo struct {
		Base struct {
			Face         string `json:"face"`           // "https://i1.hdslb.com/bfs/face/361274190d08a36ed12c58b55dd4063cde3391eb.jpg"
			IsMystery    bool   `json:"is_mystery"`     // false
			Name         string `json:"name"`           // "C酱です"
			NameColor    int    `json:"name_color"`     // 0
			NameColorStr string `json:"name_color_str"` // ""
			OfficialInfo struct {
				Desc  string `json:"desc"`  // "代表作《【生化危机7DLC】伊森必须死让伊森哭着找妈妈的通关攻略视频》"
				Role  int    `json:"role"`  // 1
				Title string `json:"title"` // "2021年度巅峰主播、bilibili 知名游戏UP主、直播高能主播"
				Type  int    `json:"type"`  // 0
			} `json:"official_info"`
			OriginInfo struct {
				Face string `json:"face"` // "https://i1.hdslb.com/bfs/face/361274190d08a36ed12c58b55dd4063cde3391eb.jpg"
				Name string `json:"name"` // "C酱です"
			} `json:"origin_info"`
			RiskCtrlInfo struct {
				Face string `json:"face"` // "https://i1.hdslb.com/bfs/face/361274190d08a36ed12c58b55dd4063cde3391eb.jpg"
				Name string `json:"name"` // "C酱です"
			} `json:"risk_ctrl_info"`
		} `json:"base"`
		Guard       any `json:"guard"`
		GuardLeader any `json:"guard_leader"`
		Medal       any `json:"medal"`
		Title       any `json:"title"`
		UheadFrame  any `json:"uhead_frame"`
		UID         int `json:"uid"` // 67141
		Wealth      any `json:"wealth"`
	} `json:"receiver_uinfo"`
	Ruid        int `json:"ruid"` // 67141
	SendMaster  any `json:"send_master"`
	SenderUinfo struct {
		Base struct {
			Face         string `json:"face"`           // "https://i1.hdslb.com/bfs/face/d59a78f41c20a9efec74e5b108642651397e786f.jpg"
			IsMystery    bool   `json:"is_mystery"`     // false
			Name         string `json:"name"`           // "小鳄鱼鱼鱼c"
			NameColor    int    `json:"name_color"`     // 0
			NameColorStr string `json:"name_color_str"` // ""
			OfficialInfo struct {
				Desc  string `json:"desc"`  // ""
				Role  int    `json:"role"`  // 0
				Title string `json:"title"` // ""
				Type  int    `json:"type"`  // -1
			} `json:"official_info"`
			OriginInfo struct {
				Face string `json:"face"` // "https://i1.hdslb.com/bfs/face/d59a78f41c20a9efec74e5b108642651397e786f.jpg"
				Name string `json:"name"` // "小鳄鱼鱼鱼c"
			} `json:"origin_info"`
			RiskCtrlInfo struct {
				Face string `json:"face"` // "https://i1.hdslb.com/bfs/face/d59a78f41c20a9efec74e5b108642651397e786f.jpg"
				Name string `json:"name"` // "小鳄鱼鱼鱼c"
			} `json:"risk_ctrl_info"`
		} `json:"base"`
		Guard       any `json:"guard"`
		GuardLeader any `json:"guard_leader"`
		Medal       any `json:"medal"`
		Title       any `json:"title"`
		UheadFrame  any `json:"uhead_frame"`
		UID         int `json:"uid"` // 4109296
		Wealth      any `json:"wealth"`
	} `json:"sender_uinfo"`
	TotalNum    int    `json:"total_num"`    // 2
	UID         int    `json:"uid"`          // 4109296
	Uname       string `json:"uname"`        // "小鳄鱼鱼鱼c"
	WealthLevel int    `json:"wealth_level"` // 26
}

func (ComboSend) Cmd() string {
	return CmdComboSend
}

type DanmakuMedal struct {
	Raw    Raw
	Level  int
	Name   string
	UpName string
	RoomID int
	UpUID  int
}

func (d *DanmakuMedal) UnmarshalJSON(data []byte) error {
	d.Raw = data

	var r []any
	err := json.Unmarshal(data, &r)
	if err != nil || len(r) == 0 {
		return err
	}

	d.Level = int(r[0].(float64))
	d.Name = r[1].(string)
	d.UpName = r[2].(string)
	d.RoomID = int(r[3].(float64))
	d.UpUID = int(r[12].(float64))

	return nil
}

func (d DanmakuMedal) String() string {
	if d.Name == "" {
		return ""
	}
	return fmt.Sprintf("【%s|%d】", d.Name, d.Level)
}

type DanmakuUser struct {
	UID          int
	Uname        string
	RoomAdmin    int
	Vip          int
	SVip         int
	Rank         int
	MobileVerify int
	UnameColor   string
}

func (d *DanmakuUser) UnmarshalJSON(data []byte) error {
	var r []any
	err := json.Unmarshal(data, &r)
	if err != nil {
		return err
	}

	d.UID = int(r[0].(float64))
	d.Uname = r[1].(string)
	d.RoomAdmin = int(r[2].(float64))
	d.Vip = int(r[3].(float64))
	d.SVip = int(r[4].(float64))
	d.Rank = int(r[5].(float64))
	d.MobileVerify = int(r[6].(float64))
	d.UnameColor = r[7].(string)

	return nil
}

func (d DanmakuUser) MarshalJSON() ([]byte, error) {
	return json.Marshal([]any{
		d.UID,
		d.Uname,
		d.RoomAdmin,
		d.Vip,
		d.SVip,
		d.Rank,
		d.MobileVerify,
		d.UnameColor,
	})
}

type DanmakuEmot struct {
	Count          int    `json:"count"`
	Descript       string `json:"descript"`
	Emoji          string `json:"emoji"`
	EmoticonID     int    `json:"emoticon_id"`
	EmoticonUnique string `json:"emoticon_unique"`
	Height         int    `json:"height"`
	URL            string `json:"url"`
	Width          int    `json:"width"`
}

type DanmakuExtra struct {
	SendFromMe            bool                   `json:"send_from_me"`
	Mode                  int                    `json:"mode"`
	Color                 int                    `json:"color"`
	DmType                int                    `json:"dm_type"`
	FontSize              int                    `json:"font_size"`
	PlayerMode            int                    `json:"player_mode"`
	ShowPlayerType        int                    `json:"show_player_type"`
	Content               string                 `json:"content"`
	UserHash              string                 `json:"user_hash"`
	EmoticonUnique        string                 `json:"emoticon_unique"`
	BulgeDisplay          int                    `json:"bulge_display"`
	RecommendScore        int                    `json:"recommend_score"`
	MainStateDmColor      string                 `json:"main_state_dm_color"`
	ObjectiveStateDmColor string                 `json:"objective_state_dm_color"`
	Direction             int                    `json:"direction"`
	PkDirection           int                    `json:"pk_direction"`
	QuartetDirection      int                    `json:"quartet_direction"`
	AnniversaryCrowd      int                    `json:"anniversary_crowd"`
	YeahSpaceType         string                 `json:"yeah_space_type"`
	YeahSpaceURL          string                 `json:"yeah_space_url"`
	JumpToURL             string                 `json:"jump_to_url"`
	SpaceType             string                 `json:"space_type"`
	SpaceURL              string                 `json:"space_url"`
	Animation             any                    `json:"animation"`
	Emots                 map[string]DanmakuEmot `json:"emots"`
	IsAudited             bool                   `json:"is_audited"`
	IDStr                 string                 `json:"id_str"`
	Icon                  any                    `json:"icon"`
	ShowReply             bool                   `json:"show_reply"`
	ReplyMID              int                    `json:"reply_mid"`
	ReplyUname            string                 `json:"reply_uname"`
	ReplyUnameColor       string                 `json:"reply_uname_color"`
	ReplyIsMystery        bool                   `json:"reply_is_mystery"`
	HitCombo              int                    `json:"hit_combo"`
}

type DanmakuDetail struct {
	Mode           int    `json:"mode"`
	ShowPlayerType int    `json:"show_player_type"`
	ExtraRaw       string `json:"extra"`
	User           struct {
		UID  int `json:"uid"`
		Base struct {
			Name         string `json:"name"`
			Face         string `json:"face"`
			NameColor    int    `json:"name_color"`
			IsMystery    bool   `json:"is_mystery"`
			RiskCtrlInfo any    `json:"risk_ctrl_info"`
			OriginInfo   struct {
				Name string `json:"name"`
				Face string `json:"face"`
			} `json:"origin_info"`
			OfficialInfo struct {
				Role  int    `json:"role"`
				Title string `json:"title"`
				Desc  string `json:"desc"`
				Type  int    `json:"type"`
			} `json:"official_info"`
			NameColorStr string `json:"name_color_str"`
		} `json:"base"`
		Medal  any `json:"medal"`
		Wealth any `json:"wealth"`
		Title  struct {
			OldTitleCSSID string `json:"old_title_css_id"`
			TitleCSSID    string `json:"title_css_id"`
		} `json:"title"`
		Guard       any `json:"guard"`
		UheadFrame  any `json:"uhead_frame"`
		GuardLeader struct {
			IsGuardLeader bool `json:"is_guard_leader"`
		} `json:"guard_leader"`
	} `json:"user"`
}

func (d *DanmakuDetail) Extra() (extra DanmakuExtra) {
	json.Unmarshal([]byte(d.ExtraRaw), &extra)
	return
}

type DanmakuData struct {
	Raw          Raw
	SendMode     int
	SendFontSize int
	DanmakuColor int
	Time         int64
	DMID         int
	MsgType      int
	Bubble       string
}

func (d *DanmakuData) UnmarshalJSON(data []byte) error {
	d.Raw = data
	d.SendMode = d.Raw.Int(1)
	d.SendFontSize = d.Raw.Int(2)
	d.DanmakuColor = d.Raw.Int(3)
	d.Time = int64(gjson.GetBytes(d.Raw, "4").Num)
	d.DMID = d.Raw.Int(5)
	d.MsgType = d.Raw.Int(10)
	d.Bubble = gjson.GetBytes(data, "11").Str
	return nil
}

func (d *DanmakuData) Detail() (detail DanmakuDetail) {
	json.Unmarshal(d.Raw.Index(15), &detail)
	return
}

type DanmakuInfo struct {
	Raw     Raw
	Data    DanmakuData
	Content string
	User    DanmakuUser
	Medal   DanmakuMedal
}

func (d *DanmakuInfo) UnmarshalJSON(data []byte) (err error) {
	d.Raw = data

	err = d.Data.UnmarshalJSON(d.Raw.Index(0))
	if err != nil {
		return
	}

	d.Content = gjson.GetBytes(data, "1").Str

	err = d.User.UnmarshalJSON(d.Raw.Index(2))
	if err != nil {
		return
	}

	err = d.Medal.UnmarshalJSON(d.Raw.Index(3))
	if err != nil {
		return
	}

	return
}

type Danmaku struct {
	RawCmd string      `json:"cmd"`
	Info   DanmakuInfo `json:"info"`
	DmV2   string      `json:"dm_v2"`
}

func (Danmaku) NoWrapper() {}

func (Danmaku) Cmd() string {
	return CmdDanmaku
}

func (Danmaku) CmdAlias() []string {
	return []string{CmdDanmakuWithCode}
}

func (d *Danmaku) Content() string {
	return d.Info.Content
}

func (d *Danmaku) UID() int {
	return d.Info.User.UID
}

func (d *Danmaku) Uname() string {
	return d.Info.User.Uname
}

func (d *Danmaku) Medal() string {
	return d.Info.Medal.Name
}

func (d *Danmaku) Time() int64 {
	return d.Info.Data.Time
}

func (d *Danmaku) Face() string {
	return d.Info.Data.Detail().User.Base.Face
}

func (d Danmaku) String() string {
	return fmt.Sprintf("%s%s(%d): %s (%s)", d.Info.Medal, d.Info.User.Uname, d.Info.User.UID, d.Info.Content, time.UnixMilli(d.Info.Data.Time).Format(time.TimeOnly))
}

type DmInteraction struct {
	DataRaw string `json:"data"`    // "{"combo":[{"id":37699755023360,"status":4,"content":"人妻","cnt":25,"guide":"他们都在说:","left_duration":6000,"fade_duration":10000,"prefix_icon":""},{"id":37699834614784,"status":4,"content":"？","cnt":5,"guide":"他们都在说:","left_duration":19000,"fade_duration":10000,"prefix_icon":""}],"merge_interval":1000,"card_appear_interval":1000,"send_interval":1000,"reset_cnt":1,"display_flag":0}"
	Dmscore int    `json:"dmscore"` // 36
	ID      int64  `json:"id"`      // 37699755023360
	Status  int    `json:"status"`  // 4
	Type    int    `json:"type"`    // 102
}

func (DmInteraction) Cmd() string {
	return CmdDmInteraction
}

// 这个解出来不是一个固定的结构 但是一定是键值对类型
func (d *DmInteraction) Data() map[string]any {
	m := make(map[string]any)
	json.Unmarshal([]byte(d.DataRaw), &m)
	return m
}

type EntryEffect struct {
	ID                   int    `json:"id"`                 // 380
	UID                  int    `json:"uid"`                // 3454913
	TargetID             int    `json:"target_id"`          // 508963009
	MockEffect           int    `json:"mock_effect"`        // 0
	Face                 string `json:"face"`               // "https://i2.hdslb.com/bfs/face/17c1855bfc0ef24a7b50fbecc8b10531cb2bb0f2.jpg"
	PrivilegeType        int    `json:"privilege_type"`     // 0
	CopyWriting          string `json:"copy_writing"`       // "<%爱染已%> 来了"
	CopyColor            string `json:"copy_color"`         // "#F7F7F7"
	HighlightColor       string `json:"highlight_color"`    // "#FFFFFF"
	Priority             int    `json:"priority"`           // 1
	BasemapURL           string `json:"basemap_url"`        // ""
	ShowAvatar           int    `json:"show_avatar"`        // 0
	EffectiveTime        int    `json:"effective_time"`     // 0
	WebBasemapURL        string `json:"web_basemap_url"`    // "https://i0.hdslb.com/bfs/live/mlive/19e7564ed9d466b02f341abfa979c6e38c2ffffb.png"
	WebEffectiveTime     int    `json:"web_effective_time"` // 4
	WebEffectClose       int    `json:"web_effect_close"`   // 1
	WebCloseTime         int    `json:"web_close_time"`     // 900
	Business             int    `json:"business"`           // 6
	CopyWritingV2        string `json:"copy_writing_v2"`    // "<%爱染已%> 来了"
	IconList             []any  `json:"icon_list"`
	MaxDelayTime         int    `json:"max_delay_time"`          // 7
	TriggerTime          int64  `json:"trigger_time"`            // 1735284739438265222
	Identities           int    `json:"identities"`              // 1
	EffectSilentTime     int    `json:"effect_silent_time"`      // 0
	EffectiveTimeNew     int    `json:"effective_time_new"`      // 0
	WebDynamicURLWebp    string `json:"web_dynamic_url_webp"`    // ""
	WebDynamicURLApng    string `json:"web_dynamic_url_apng"`    // ""
	MobileDynamicURLWebp string `json:"mobile_dynamic_url_webp"` // ""
	WealthyInfo          struct {
		UID              int    `json:"uid"`                // 0
		Level            int    `json:"level"`              // 16
		LevelTotalScore  int    `json:"level_total_score"`  // 0
		CurScore         int    `json:"cur_score"`          // 0
		UpgradeNeedScore int    `json:"upgrade_need_score"` // 0
		Status           int    `json:"status"`             // 0
		DmIconKey        string `json:"dm_icon_key"`        // ""
	} `json:"wealthy_info"`
	NewStyle  int  `json:"new_style"`  // 1
	IsMystery bool `json:"is_mystery"` // false
	Uinfo     struct {
		UID  int `json:"uid"` // 3454913
		Base struct {
			Name         string `json:"name"`       // "爱染已"
			Face         string `json:"face"`       // "http://i2.hdslb.com/bfs/face/17c1855bfc0ef24a7b50fbecc8b10531cb2bb0f2.jpg"
			NameColor    int    `json:"name_color"` // 0
			IsMystery    bool   `json:"is_mystery"` // false
			RiskCtrlInfo any    `json:"risk_ctrl_info"`
			OriginInfo   any    `json:"origin_info"`
			OfficialInfo any    `json:"official_info"`
			NameColorStr string `json:"name_color_str"` // ""
		} `json:"base"`
		Medal struct {
			Name               string `json:"name"`                  // "帅Pi"
			Level              int    `json:"level"`                 // 18
			ColorStart         int    `json:"color_start"`           // 13081892
			ColorEnd           int    `json:"color_end"`             // 13081892
			ColorBorder        int    `json:"color_border"`          // 13081892
			Color              int    `json:"color"`                 // 13081892
			ID                 int    `json:"id"`                    // 1786
			Typ                int    `json:"typ"`                   // 0
			IsLight            int    `json:"is_light"`              // 1
			Ruid               int    `json:"ruid"`                  // 13046
			GuardLevel         int    `json:"guard_level"`           // 0
			Score              int    `json:"score"`                 // 429660
			GuardIcon          string `json:"guard_icon"`            // ""
			HonorIcon          string `json:"honor_icon"`            // ""
			V2MedalColorStart  string `json:"v2_medal_color_start"`  // "#DC6B6B99"
			V2MedalColorEnd    string `json:"v2_medal_color_end"`    // "#DC6B6B99"
			V2MedalColorBorder string `json:"v2_medal_color_border"` // "#DC6B6B99"
			V2MedalColorText   string `json:"v2_medal_color_text"`   // "#FFFFFFFF"
			V2MedalColorLevel  string `json:"v2_medal_color_level"`  // "#81001F99"
			UserReceiveCount   int    `json:"user_receive_count"`    // 0
		} `json:"medal"`
		Wealth struct {
			Level     int    `json:"level"`       // 16
			DmIconKey string `json:"dm_icon_key"` // ""
		} `json:"wealth"`
		Title any `json:"title"`
		Guard struct {
			Level      int    `json:"level"`       // 0
			ExpiredStr string `json:"expired_str"` // ""
		} `json:"guard"`
		UheadFrame  any `json:"uhead_frame"`
		GuardLeader any `json:"guard_leader"`
	} `json:"uinfo"`
	FullCartoonID   int `json:"full_cartoon_id"` // 1802
	PriorityLevel   int `json:"priority_level"`  // 0
	WealthStyleInfo struct {
		URL string `json:"url"` // "https://i0.hdslb.com/bfs/live/ea32c9db0a5402580541ec974f336ad55b473869.png"
	} `json:"wealth_style_info"`
}

func (EntryEffect) Cmd() string {
	return CmdEntryEffect
}

type GuardBuy struct {
	GuardLevel int    `json:"guard_level"` // 2
	Price      int    `json:"price"`       // 1998000
	UID        int    `json:"uid"`         // 1751924
	Num        int    `json:"num"`         // 1
	GiftID     int    `json:"gift_id"`     // 10002
	GiftName   string `json:"gift_name"`   // "提督"
	StartTime  int    `json:"start_time"`  // 1635766940
	EndTime    int    `json:"end_time"`    // 1635766940
	Username   string `json:"username"`    // "ppmann"
}

func (GuardBuy) Cmd() string {
	return CmdGuardBuy
}

type HotRankChanged struct {
	Rank        int    `json:"rank"`
	Timestamp   int64  `json:"timestamp"`
	WebUrl      string `json:"web_url"`
	LiveUrl     string `json:"live_url"`
	LiveLinkUrl string `json:"live_link_url"`
	AreaName    string `json:"area_name"`
	Trend       int    `json:"trend"`
	Countdown   int    `json:"countdown"`
	BlinkUrl    string `json:"blink_url"`
	PCLinkUrl   string `json:"pc_link_url"`
	Icon        string `json:"icon"`
}

func (HotRankChanged) Cmd() string {
	return CmdHotRankChanged
}

type HotRankSettlement struct {
	DmMsg     string `json:"dm_msg"`    // "恭喜主播 <% 老实憨厚的笑笑 %> 荣登限时热门榜网游榜top3! 即将获得热门流量推荐哦！"
	Dmscore   int    `json:"dmscore"`   // 144
	Timestamp int    `json:"timestamp"` // 1635748200
	Uname     string `json:"uname"`     // "老实憨厚的笑笑"
	URL       string `json:"url"`       // "https://live.bilibili.com/p/html/live-app-hotrank/result.html?is_live_half_webview=1&hybrid_half_ui=1,5,250,200,f4eefa,0,30,0,0,0;2,5,250,200,f4eefa,0,30,0,0,0;3,5,250,200,f4eefa,0,30,0,0,0;4,5,250,200,f4eefa,0,30,0,0,0;5,5,250,200,f4eefa,0,30,0,0,0;6,5,250,200,f4eefa,0,30,0,0,0;7,5,250,200,f4eefa,0,30,0,0,0;8,5,250,200,f4eefa,0,30,0,0,0&areaId=2&cache_key=d04fe4e6d0f2bc24c1a5acc63f574b68"
	AreaName  string `json:"area_name"` // "网游"
	CacheKey  string `json:"cache_key"` // "d04fe4e6d0f2bc24c1a5acc63f574b68"
	Rank      int    `json:"rank"`      // 3
	Face      string `json:"face"`      // "http://i0.hdslb.com/bfs/face/d92282ac664afffd0ef38c275f3fca17a9567d5a.jpg"
	Icon      string `json:"icon"`      // "https://i0.hdslb.com/bfs/live/65dbe013f7379c78fc50dfb2fd38d67f5e4895f9.png"
}

func (HotRankSettlement) Cmd() string {
	return CmdHotRankSettlement
}

type InteractWord struct {
	TailIcon     int    `json:"tail_icon"`   // 0
	UID          int    `json:"uid"`         // 13729609
	Uname        string `json:"uname"`       // "阿羽AYu-"
	UnameColor   string `json:"uname_color"` // ""
	Dmscore      int    `json:"dmscore"`     // 12
	Score        int64  `json:"score"`       // 1635849011802
	SpreadDesc   string `json:"spread_desc"` // ""
	Timestamp    int    `json:"timestamp"`   // 1635849011
	Identities   []int  `json:"identities"`
	IsSpread     int    `json:"is_spread"`    // 0
	RoomID       int    `json:"roomid"`       // 732602
	TriggerTime  int64  `json:"trigger_time"` // 1635849010751414020
	Contribution struct {
		Grade int `json:"grade"` // 0
	} `json:"contribution"`
	FansMedal struct {
		MedalColor       int    `json:"medal_color"`        // 12632256
		MedalColorStart  int    `json:"medal_color_start"`  // 12632256
		MedalLevel       int    `json:"medal_level"`        // 10
		Score            int    `json:"score"`              // 10103
		TargetID         int    `json:"target_id"`          // 67141
		GuardLevel       int    `json:"guard_level"`        // 0
		IconID           int    `json:"icon_id"`            // 0
		IsLighted        int    `json:"is_lighted"`         // 0
		MedalName        string `json:"medal_name"`         // "C酱"
		Special          string `json:"special"`            // ""
		AnchorRoomID     int    `json:"anchor_roomid"`      // 47867
		MedalColorBorder int    `json:"medal_color_border"` // 12632256
		MedalColorEnd    int    `json:"medal_color_end"`    // 12632256
	} `json:"fans_medal"`
	MsgType    int    `json:"msg_type"`    // 1
	SpreadInfo string `json:"spread_info"` // ""
}

func (InteractWord) Cmd() string {
	return CmdInteractWord
}

type LikeInfoV3Click struct {
	ContributionInfo struct {
		Grade int `json:"grade"` // 0
	} `json:"contribution_info"`
	Dmscore   int `json:"dmscore"` // 18
	FansMedal struct {
		AnchorRoomID     int    `json:"anchor_roomid"`      // 0
		GuardLevel       int    `json:"guard_level"`        // 0
		IconID           int    `json:"icon_id"`            // 0
		IsLighted        int    `json:"is_lighted"`         // 0
		MedalColor       int    `json:"medal_color"`        // 0
		MedalColorBorder int    `json:"medal_color_border"` // 0
		MedalColorEnd    int    `json:"medal_color_end"`    // 0
		MedalColorStart  int    `json:"medal_color_start"`  // 0
		MedalLevel       int    `json:"medal_level"`        // 0
		MedalName        string `json:"medal_name"`         // ""
		Score            int    `json:"score"`              // 0
		Special          string `json:"special"`            // ""
		TargetID         int    `json:"target_id"`          // 0
	} `json:"fans_medal"`
	GroupMedal any    `json:"group_medal"`
	Identities []int  `json:"identities"`
	IsMystery  bool   `json:"is_mystery"` // false
	LikeIcon   string `json:"like_icon"`  // "https://i0.hdslb.com/bfs/live/23678e3d90402bea6a65251b3e728044c21b1f0f.png"
	LikeText   string `json:"like_text"`  // "为主播点赞了"
	MsgType    int    `json:"msg_type"`   // 6
	ShowArea   int    `json:"show_area"`  // 0
	UID        int    `json:"uid"`        // 2621417
	Uinfo      struct {
		Base struct {
			Face         string `json:"face"`           // "http://i2.hdslb.com/bfs/face/63dbba1bf6f4e897668a2e252250812a5150741f.jpg"
			IsMystery    bool   `json:"is_mystery"`     // false
			Name         string `json:"name"`           // "面面面面面条丶"
			NameColor    int    `json:"name_color"`     // 0
			NameColorStr string `json:"name_color_str"` // ""
			OfficialInfo struct {
				Desc  string `json:"desc"`  // ""
				Role  int    `json:"role"`  // 0
				Title string `json:"title"` // ""
				Type  int    `json:"type"`  // -1
			} `json:"official_info"`
			OriginInfo struct {
				Face string `json:"face"` // "http://i2.hdslb.com/bfs/face/63dbba1bf6f4e897668a2e252250812a5150741f.jpg"
				Name string `json:"name"` // "面面面面面条丶"
			} `json:"origin_info"`
			RiskCtrlInfo any `json:"risk_ctrl_info"`
		} `json:"base"`
		Guard struct {
			ExpiredStr string `json:"expired_str"` // ""
			Level      int    `json:"level"`       // 0
		} `json:"guard"`
		GuardLeader any `json:"guard_leader"`
		Medal       any `json:"medal"`
		Title       any `json:"title"`
		UheadFrame  any `json:"uhead_frame"`
		UID         int `json:"uid"` // 2621417
		Wealth      any `json:"wealth"`
	} `json:"uinfo"`
	Uname      string `json:"uname"`       // "面面面面面条丶"
	UnameColor string `json:"uname_color"` // ""
}

func (LikeInfoV3Click) Cmd() string {
	return CmdLikeInfoV3Click
}

type LikeInfoV3Update struct {
	ClickCount int `json:"click_count"` // 6567
}

func (LikeInfoV3Update) Cmd() string {
	return CmdLikeInfoV3Update
}

type LittleMessageBox struct {
	From     string `json:"from"` // "fans_medal"
	Platform struct {
		Android bool `json:"android"` // true
		Ios     bool `json:"ios"`     // true
		Web     bool `json:"web"`     // true
	} `json:"platform"`
	Msg    string `json:"msg"`     // "今日首条弹幕发送成功~亲密度+100"
	RoomID int    `json:"room_id"` // 8432038
	Type   int    `json:"type"`    // 1
}

func (LittleMessageBox) Cmd() string {
	return CmdLittleMessageBox
}

type Live struct {
	LiveKey         string `json:"live_key"`         // "508011583643868085"
	VoiceBackground string `json:"voice_background"` // ""
	SubSessionKey   string `json:"sub_session_key"`  // "508011583643868085sub_time:1717090320"
	LivePlatform    string `json:"live_platform"`    // "pc"
	LiveModel       int    `json:"live_model"`       // 0
	RoomID          int    `json:"roomid"`           // 14703541
	LiveTime        int    `json:"live_time"`        // 1717090319
}

func (Live) NoWrapper() {}

func (Live) Cmd() string {
	return CmdLive
}

type LogInNotice struct {
	NoticeMsg string `json:"notice_msg"` // "为保护用户隐私，未登录无法查看他人昵称"
	ImageWeb  string `json:"image_web"`  // "http://i0.hdslb.com/bfs/dm/75e7c16b99208df259fe0a93354fd3440cbab412.png"
	ImageApp  string `json:"image_app"`  // "http://i0.hdslb.com/bfs/dm/b632f7dcd3acf47deffb5f9ccc9546ae97a3415b.png"
}

func (LogInNotice) Cmd() string {
	return CmdLogInNotice
}

type MessageboxUserMedalChange struct {
	Type              int    `json:"type"`                // 2
	UID               int    `json:"uid"`                 // 188888131
	UpUID             int    `json:"up_uid"`              // 267766441
	MedalLevel        int    `json:"medal_level"`         // 15
	MedalName         string `json:"medal_name"`          // "全撤了"
	MedalColorStart   int    `json:"medal_color_start"`   // 12478086
	MedalColorEnd     int    `json:"medal_color_end"`     // 12478086
	MedalColorBorder  int    `json:"medal_color_border"`  // 12478086
	IsLighted         int    `json:"is_lighted"`          // 1
	IsLightedV2       bool   `json:"is_lighted_v2"`       // true
	GuardLevel        int    `json:"guard_level"`         // 0
	Unlock            int    `json:"unlock"`              // 0
	UnlockLevel       int    `json:"unlock_level"`        // 0
	MultiUnlockLevel  string `json:"multi_unlock_level"`  // ""
	UpperBoundContent string `json:"upper_bound_content"` // ""
}

func (MessageboxUserMedalChange) Cmd() string {
	return CmdMessageboxUserMedalChange
}

type NoticeMsg struct {
	BusinessID string `json:"business_id"` // "31087"
	Full       struct {
		HeadIcon    string `json:"head_icon"`     // "http://i0.hdslb.com/bfs/live/00f26756182b2e9d06c00af23001bc8e10da67d0.webp"
		TailIcon    string `json:"tail_icon"`     // "http://i0.hdslb.com/bfs/live/822da481fdaba986d738db5d8fd469ffa95a8fa1.webp"
		HeadIconFa  string `json:"head_icon_fa"`  // "http://i0.hdslb.com/bfs/live/77983005023dc3f31cd599b637c83a764c842f87.png"
		TailIconFa  string `json:"tail_icon_fa"`  // "http://i0.hdslb.com/bfs/live/38cb2a9f1209b16c0f15162b0b553e3b28d9f16f.png"
		Background  string `json:"background"`    // "#6098FFFF"
		Highlight   string `json:"highlight"`     // "#FDFF2FFF"
		HeadIconFan int    `json:"head_icon_fan"` // 36
		TailIconFan int    `json:"tail_icon_fan"` // 4
		Color       string `json:"color"`         // "#FFFFFFFF"
		Time        int    `json:"time"`          // 20
	} `json:"full"`
	Half struct {
		Time       int    `json:"time"`       // 15
		HeadIcon   string `json:"head_icon"`  // "http://i0.hdslb.com/bfs/live/358cc52e974b315e83eee429858de4fee97a1ef5.png"
		TailIcon   string `json:"tail_icon"`  // ""
		Background string `json:"background"` // "#7BB6F2FF"
		Color      string `json:"color"`      // "#FFFFFFFF"
		Highlight  string `json:"highlight"`  // "#FDFF2FFF"
	} `json:"half"`
	ID         int    `json:"id"`          // 2
	LinkURL    string `json:"link_url"`    // "https://live.bilibili.com/5655865?accept_quality=%5B10000%2C150%5D&broadcast_type=0&current_qn=150&current_quality=150&is_room_feed=1&live_play_network=other&p2p_type=-2&playurl_h264=http%3A%2F%2Fd1--cn-gotcha03.bilivideo.com%2Flive-bvc%2F429443%2Flive_2257663_5953069_1500.flv%3Fexpires%3D1635753433%26len%3D0%26oi%3D0%26pt%3D%26qn%3D150%26trid%3D10004aaecf5169e74b51b5932933468e0364%26sigparams%3Dcdn%2Cexpires%2Clen%2Coi%2Cpt%2Cqn%2Ctrid%26cdn%3Dcn-gotcha03%26sign%3De0b8728896efe026833d99655b05c084%26p2p_type%3D4294967294%26src%3D5%26sl%3D1%26flowtype%3D1%26source%3Dbatch%26order%3D1%26machinezone%3Dylf%26sk%3D2935686d6cb9146c7a6a6a0b4e120e2594e074fa0760377f1a7a2b2fa0ee6443&playurl_h265=&quality_description=%5B%7B%22qn%22%3A10000%2C%22desc%22%3A%22%E5%8E%9F%E7%94%BB%22%7D%2C%7B%22qn%22%3A150%2C%22desc%22%3A%22%E9%AB%98%E6%B8%85%22%7D%5D&from=28003&extra_jump_from=28003&live_lottery_type=1"
	MsgCommon  string `json:"msg_common"`  // "<%JamesTuT%>投喂:<%木之本切%>1个次元之城，点击前往TA的房间吧！"
	MsgSelf    string `json:"msg_self"`    // "<%JamesTuT%>投喂:<%木之本切%>1个次元之城，快来围观吧！"
	MsgType    int    `json:"msg_type"`    // 2
	Name       string `json:"name"`        // "分区道具抽奖广播样式"
	RealRoomID int    `json:"real_roomid"` // 5655865
	RoomID     int    `json:"roomid"`      // 5655865
	Scatter    struct {
		Min int `json:"min"` // 0
		Max int `json:"max"` // 0
	} `json:"scatter"`
	ShieldUID int `json:"shield_uid"` // -1
	Side      struct {
		HeadIcon   string `json:"head_icon"`  // ""
		Background string `json:"background"` // ""
		Color      string `json:"color"`      // ""
		Highlight  string `json:"highlight"`  // ""
		Border     string `json:"border"`     // ""
	} `json:"side"`
}

func (NoticeMsg) NoWrapper() {}

func (NoticeMsg) Cmd() string {
	return CmdNoticeMsg
}

type OnlineRankCount struct {
	Count           int    `json:"count"`             // 5679
	CountText       string `json:"count_text"`        // "5679"
	OnlineCount     int    `json:"online_count"`      // 6455
	OnlineCountText string `json:"online_count_text"` // "6455"
}

func (OnlineRankCount) Cmd() string {
	return CmdOnlineRankCount
}

type OnlineRankTop3 struct {
	Dmscore int `json:"dmscore"` // 112
	List    []struct {
		Msg  string `json:"msg"`  // "恭喜 <%你们有多腐%> 成为高能榜"
		Rank int    `json:"rank"` // 2
	} `json:"list"`
}

func (OnlineRankTop3) Cmd() string {
	return CmdOnlineRankTop3
}

type OnlineRankV2 struct {
	List []struct {
		GuardLevel int    `json:"guard_level"` // 2 (3)舰长 (2)提督 (1)总督?
		UID        int    `json:"uid"`         // 277278853
		Face       string `json:"face"`        // "http://i1.hdslb.com/bfs/face/59839130848b8f8d99f8c649f7897ac7f406a052.jpg"
		Score      string `json:"score"`       // "15980"
		Uname      string `json:"uname"`       // "勤俭持家的席撒"
		Rank       int    `json:"rank"`        // 1
	} `json:"list"`
	RankType string `json:"rank_type"` // "gold-rank"
}

func (OnlineRankV2) Cmd() string {
	return CmdOnlineRankV2
}

type Preparing struct {
	RoomID string `json:"roomid"` // "14703541"
	Round  int    `json:"round"`  // 1
}

func (Preparing) NoWrapper() {}

func (Preparing) Cmd() string {
	return CmdPreparing
}

type RankChanged struct {
	UID              int    `json:"uid"`                  // 508963009
	Rank             int    `json:"rank"`                 // 27
	Countdown        int    `json:"countdown"`            // 0
	Timestamp        int    `json:"timestamp"`            // 1735284374
	OnRankNameByType string `json:"on_rank_name_by_type"` // "全站热门"
	RankNameByType   string `json:"rank_name_by_type"`    // "热门榜"
	URLByType        string `json:"url_by_type"`          // "https://live.bilibili.com/p/html/live-app-hotrank/index.html?is_live_half_webview=1&hybrid_rotate_d=1&hybrid_half_ui=1,3,100p,70p,0,0,30,100,12;2,2,375,100p,0,0,30,100,0;3,3,100p,70p,0,0,30,100,12;4,2,375,100p,0,0,30,100,0;5,3,100p,70p,0,0,30,100,0;6,3,100p,70p,0,0,30,100,0;7,3,100p,70p,0,0,30,100,0;8,3,100p,70p,0,0,30,100,0&pc_ui=338,465,f4eefa,0&redirect=v2&rank=hot&anchorId=508963009&rank_type=1"
	RankByType       int    `json:"rank_by_type"`         // 27
	RankType         int    `json:"rank_type"`            // 3
}

func (RankChanged) Cmd() string {
	return CmdRankChanged
}

type RevenueRankChanged struct {
	ConfID      int    `json:"conf_id"`       // 11
	RankName    string `json:"rank_name"`     // "虚拟航海"
	UID         int    `json:"uid"`           // 508963009
	Rank        int    `json:"rank"`          // 13
	IconURLBlue string `json:"icon_url_blue"` // "https://i0.hdslb.com/bfs/live/18e2990a546d33368200f9058f3d9dbc4038eb5c.png"
	IconURLPink string `json:"icon_url_pink"` // "https://i0.hdslb.com/bfs/live/a6c490c36e88c7b191a04883a5ec15aed187a8f7.png"
	IconURLGrey string `json:"icon_url_grey"` // "https://i0.hdslb.com/bfs/live/cb7444b1faf1d785df6265bfdc1fcfc993419b76.png"
	ActionType  int    `json:"action_type"`   // 1
	Timestamp   int    `json:"timestamp"`     // 1735285539
	MsgID       string `json:"msg_id"`        // "a9b0ffe4-71be-463b-8415-db8c7835f493"
	JumpURLLink string `json:"jump_url_link"` // "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=3&ruid=508963009&conf_id=11&is_live_half_webview=1&hybrid_rotate_d=1&is_cling_player=1&hybrid_half_ui=1,3,100p,70p,f4eefa,0,30,100,0,0;2,2,375,100p,f4eefa,0,30,100,0,0;3,3,100p,70p,f4eefa,0,30,100,0,0;4,2,375,100p,f4eefa,0,30,100,0,0;5,3,100p,70p,f4eefa,0,30,100,0,0;6,3,100p,70p,f4eefa,0,30,100,0,0;7,3,100p,70p,f4eefa,0,30,100,0,0;8,3,100p,70p,f4eefa,0,30,100,0,0#/area-rank"
	JumpURLPc   string `json:"jump_url_pc"`   // "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=4&ruid=508963009&conf_id=11&pc_ui=338,465,f4eefa,0#/area-rank"
	JumpURLPink string `json:"jump_url_pink"` // "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=1&ruid=508963009&conf_id=11&is_live_half_webview=1&hybrid_rotate_d=1&hybrid_half_ui=1,3,100p,70p,ffffff,0,30,100,12,0;2,2,375,100p,ffffff,0,30,100,0,0;3,3,100p,70p,ffffff,0,30,100,12,0;4,2,375,100p,ffffff,0,30,100,0,0;5,3,100p,70p,ffffff,0,30,100,0,0;6,3,100p,70p,ffffff,0,30,100,0,0;7,3,100p,70p,ffffff,0,30,100,0,0;8,3,100p,70p,ffffff,0,30,100,0,0#/area-rank"
	JumpURLWeb  string `json:"jump_url_web"`  // "https://live.bilibili.com/p/html/live-app-hotrank/index.html?clientType=2&ruid=508963009&conf_id=11#/area-rank"
}

func (RevenueRankChanged) Cmd() string {
	return CmdRevenueRankChanged
}

type RoomBlockMsg struct {
	Uname    string `json:"uname"`    // "白绫彡"
	Dmscore  int    `json:"dmscore"`  // 30
	Operator int    `json:"operator"` // 1
	UID      int    `json:"uid"`      // 53342046
}

func (RoomBlockMsg) Cmd() string {
	return CmdRoomBlockMsg
}

type RoomChange struct {
	ParentAreaID   int    `json:"parent_area_id"`   // 3
	AreaName       string `json:"area_name"`        // "原神"
	ParentAreaName string `json:"parent_area_name"` // "手游"
	LiveKey        string `json:"live_key"`         // "181443822587250220"
	SubSessionKey  string `json:"sub_session_key"`  // "181443822587250220sub_time:1635846195"
	Title          string `json:"title"`            // "快来直播间抽胡桃！托马！烈火拔刀！"
	AreaID         int    `json:"area_id"`          // 321
}

func (RoomChange) Cmd() string {
	return CmdRoomChange
}

type RoomRealTimeMessageUpdate struct {
	FansClub  int `json:"fans_club"`  // 49182
	RoomID    int `json:"roomid"`     // 545068
	Fans      int `json:"fans"`       // 1384297
	RedNotice int `json:"red_notice"` // -1
}

func (RoomRealTimeMessageUpdate) Cmd() string {
	return CmdRoomRealTimeMessageUpdate
}

// SendGift 部分 any 抓取不到有效的信息，只能先留着
type SendGift struct {
	DiscountPrice  int    `json:"discount_price"` // 100
	GiftName       string `json:"giftName"`       // "牛哇牛哇"
	Gold           int    `json:"gold"`           // 0
	GuardLevel     int    `json:"guard_level"`    // 0
	Remain         int    `json:"remain"`         // 0
	Silver         int    `json:"silver"`         // 0
	SuperGiftNum   int    `json:"super_gift_num"` // 4
	TopList        any    `json:"top_list"`
	BizSource      string `json:"biz_source"`       // "xlottery-anchor"
	ComboTotalCoin int    `json:"combo_total_coin"` // 400
	GiftType       int    `json:"giftType"`         // 0
	Magnification  int    `json:"magnification"`    // 1
	MedalInfo      struct {
		MedalName        string `json:"medal_name"`         // "吉祥草"
		Special          string `json:"special"`            // ""
		AnchorRoomID     int    `json:"anchor_roomid"`      // 0
		AnchorUname      string `json:"anchor_uname"`       // ""
		MedalColorBorder int    `json:"medal_color_border"` // 6067854
		MedalColorEnd    int    `json:"medal_color_end"`    // 6067854
		MedalColorStart  int    `json:"medal_color_start"`  // 6067854
		MedalLevel       int    `json:"medal_level"`        // 4
		TargetID         int    `json:"target_id"`          // 2920960
		GuardLevel       int    `json:"guard_level"`        // 0
		IconID           int    `json:"icon_id"`            // 0
		IsLighted        int    `json:"is_lighted"`         // 1
		MedalColor       int    `json:"medal_color"`        // 6067854
	} `json:"medal_info"`
	NameColor         string `json:"name_color"` // ""
	Price             int    `json:"price"`      // 100
	Super             int    `json:"super"`      // 0
	TagImage          string `json:"tag_image"`  // ""
	TotalCoin         int    `json:"total_coin"` // 100
	Uname             string `json:"uname"`      // "余烬的圆舞曲"
	BlindGift         any    `json:"blind_gift"`
	Rnd               string `json:"rnd"`                  // "1635849011111500002"
	Action            string `json:"action"`               // "投喂"
	BroadcastID       int    `json:"broadcast_id"`         // 0
	Effect            int    `json:"effect"`               // 0
	GiftID            int    `json:"giftId"`               // 31039
	IsSpecialBatch    int    `json:"is_special_batch"`     // 0
	TID               string `json:"tid"`                  // "1635849011111500002"
	BatchComboID      string `json:"batch_combo_id"`       // "batch:gift:combo_id:9184735:2920960:31039:1635849007.7560"
	FloatScResourceID int    `json:"float_sc_resource_id"` // 0
	OriginalGiftName  string `json:"original_gift_name"`   // ""
	BatchComboSend    any    `json:"batch_combo_send"`
	IsFirst           bool   `json:"is_first"` // false
	Num               int    `json:"num"`      // 1
	Rcost             int    `json:"rcost"`    // 189509940
	UID               int    `json:"uid"`      // 9184735
	BeatID            string `json:"beatId"`   // ""
	ComboSend         any    `json:"combo_send"`
	ComboStayTime     int    `json:"combo_stay_time"`    // 3
	Dmscore           int    `json:"dmscore"`            // 32
	SvgaBlock         int    `json:"svga_block"`         // 0
	Timestamp         int    `json:"timestamp"`          // 1635849011
	CoinType          string `json:"coin_type"`          // "gold"
	ComboResourcesID  int    `json:"combo_resources_id"` // 1
	CritProb          int    `json:"crit_prob"`          // 0
	Demarcation       int    `json:"demarcation"`        // 1
	Draw              int    `json:"draw"`               // 0
	EffectBlock       int    `json:"effect_block"`       // 0
	Face              string `json:"face"`               // "http://i1.hdslb.com/bfs/face/80cd97607e8ab30acc768047db37a17c9270ec76.jpg"
	SendMaster        any    `json:"send_master"`
	SuperBatchGiftNum int    `json:"super_batch_gift_num"` // 4
}

func (SendGift) Cmd() string {
	return CmdSendGift
}

type StopLiveRoomList struct {
	RoomIDList []int `json:"room_id_list"`
}

func (StopLiveRoomList) Cmd() string {
	return CmdStopLiveRoomList
}

type SuperChatMessage struct {
	BackgroundBottomColor string `json:"background_bottom_color"` // "#427D9E"
	Token                 string `json:"token"`                   // "D22854F7"
	BackgroundColorEnd    string `json:"background_color_end"`    // "#29718B"
	BackgroundImage       string `json:"background_image"`        // "https://i0.hdslb.com/bfs/live/a712efa5c6ebc67bafbe8352d3e74b820a00c13e.png"
	BackgroundIcon        string `json:"background_icon"`         // ""
	BackgroundPriceColor  string `json:"background_price_color"`  // "#7DA4BD"
	Dmscore               int    `json:"dmscore"`                 // 128
	ID                    int    `json:"id"`                      // 2575658
	UserInfo              struct {
		UserLevel  int    `json:"user_level"`  // 20
		FaceFrame  string `json:"face_frame"`  // "http://i0.hdslb.com/bfs/live/9b3cfee134611c61b71e38776c58ad67b253c40a.png"
		GuardLevel int    `json:"guard_level"` // 2
		LevelColor string `json:"level_color"` // "#61c05a"
		Manager    int    `json:"manager"`     // 0
		Uname      string `json:"uname"`       // "卡纸哥我宣你"
		Title      string `json:"title"`       // "title-111-1"
		Face       string `json:"face"`        // "http://i0.hdslb.com/bfs/face/6badd87b9bf8c13c90fcb2c2b1b93b01e4b02664.jpg"
		IsMainVip  int    `json:"is_main_vip"` // 1
		IsSvip     int    `json:"is_svip"`     // 0
		IsVip      int    `json:"is_vip"`      // 0
		NameColor  string `json:"name_color"`  // "#E17AFF"
	} `json:"user_info"`
	IsSendAudit     int     `json:"is_send_audit"`    // 1
	Price           int     `json:"price"`            // 50
	BackgroundColor string  `json:"background_color"` // "#DBFFFD"
	ColorPoint      float64 `json:"color_point"`      // 0.7
	Gift            struct {
		GiftID   int    `json:"gift_id"`   // 12000
		GiftName string `json:"gift_name"` // "醒目留言"
		Num      int    `json:"num"`       // 1
	} `json:"gift"`
	MedalInfo struct {
		TargetID         int    `json:"target_id"`          // 8739477
		AnchorRoomID     int    `json:"anchor_roomid"`      // 7777
		AnchorUname      string `json:"anchor_uname"`       // "老实憨厚的笑笑"
		GuardLevel       int    `json:"guard_level"`        // 2
		MedalColor       string `json:"medal_color"`        // "#6154c"
		MedalColorEnd    int    `json:"medal_color_end"`    // 6850801
		MedalLevel       int    `json:"medal_level"`        // 27
		Special          string `json:"special"`            // ""
		IconID           int    `json:"icon_id"`            // 0
		IsLighted        int    `json:"is_lighted"`         // 1
		MedalColorBorder int    `json:"medal_color_border"` // 16771156
		MedalColorStart  int    `json:"medal_color_start"`  // 398668
		MedalName        string `json:"medal_name"`         // "德云色"
	} `json:"medal_info"`
	TransMark            int    `json:"trans_mark"`             // 0
	Ts                   int    `json:"ts"`                     // 1635749378
	BackgroundColorStart string `json:"background_color_start"` // "#4EA4C5"
	EndTime              int    `json:"end_time"`               // 1635749498
	MessageFontColor     string `json:"message_font_color"`     // "#A3F6FF"
	Rate                 int    `json:"rate"`                   // 1000
	MessageTrans         string `json:"message_trans"`          // ""
	StartTime            int    `json:"start_time"`             // 1635749378
	IsRanked             int    `json:"is_ranked"`              // 1
	Message              string `json:"message"`                // "熊神可以打一拳旁边那个大胖子吗"
	Time                 int    `json:"time"`                   // 120
	UID                  int    `json:"uid"`                    // 12777723
}

func (SuperChatMessage) Cmd() string {
	return CmdSuperChatMessage
}

type SuperChatMessageJPN struct {
	UID       string `json:"uid"`       // "179810804"
	IsRanked  int    `json:"is_ranked"` // 1
	MedalInfo struct {
		MedalColor   string `json:"medal_color"`   // "#1a544b"
		IconID       int    `json:"icon_id"`       // 0
		TargetID     int    `json:"target_id"`     // 419220
		Special      string `json:"special"`       // ""
		AnchorUname  string `json:"anchor_uname"`  // "神奇陆夫人"
		AnchorRoomID int    `json:"anchor_roomid"` // 115
		MedalLevel   int    `json:"medal_level"`   // 22
		MedalName    string `json:"medal_name"`    // "陆夫人"
	} `json:"medal_info"`
	UserInfo struct {
		UserLevel  int    `json:"user_level"`  // 20
		LevelColor string `json:"level_color"` // "#61c05a"
		IsVip      int    `json:"is_vip"`      // 0
		IsSvip     int    `json:"is_svip"`     // 0
		IsMainVip  int    `json:"is_main_vip"` // 1
		Title      string `json:"title"`       // "0"
		Uname      string `json:"uname"`       // "悲剧携带者"
		Face       string `json:"face"`        // "http://i2.hdslb.com/bfs/face/350aebb461ca8215b70c4cb4c1e8061ccb6d7db1.jpg"
		Manager    int    `json:"manager"`     // 0
		FaceFrame  string `json:"face_frame"`  // "http://i0.hdslb.com/bfs/live/78e8a800e97403f1137c0c1b5029648c390be390.png"
		GuardLevel int    `json:"guard_level"` // 3
	} `json:"user_info"`
	ID                   string `json:"id"`                     // "2576342"
	MessageJpn           string `json:"message_jpn"`            // ""
	Time                 int    `json:"time"`                   // 60
	Rate                 int    `json:"rate"`                   // 1000
	BackgroundImage      string `json:"background_image"`       // "https://i0.hdslb.com/bfs/live/a712efa5c6ebc67bafbe8352d3e74b820a00c13e.png"
	BackgroundIcon       string `json:"background_icon"`        // ""
	BackgroundPriceColor string `json:"background_price_color"` // "#7497CD"
	Token                string `json:"token"`                  // "1B6E22FC"
	Gift                 struct {
		Num      int    `json:"num"`       // 1
		GiftID   int    `json:"gift_id"`   // 12000
		GiftName string `json:"gift_name"` // "醒目留言"
	} `json:"gift"`
	Price                 int    `json:"price"`                   // 30
	Message               string `json:"message"`                 // "夫人，看你航天伴着好运来刷牛场，出了基德的财运，但是为什么mf才22（上限40啊"
	BackgroundColor       string `json:"background_color"`        // "#EDF5FF"
	BackgroundBottomColor string `json:"background_bottom_color"` // "#2A60B2"
	Ts                    int    `json:"ts"`                      // 1635766505
	StartTime             int    `json:"start_time"`              // 1635766505
	EndTime               int    `json:"end_time"`                // 1635766565
}

func (SuperChatMessageJPN) Cmd() string {
	return CmdSuperChatMessageJPN
}

type SysMsg struct {
	Msg string `json:"msg"` // "争夺开启，时间周五20点至周日20点，逾期不候哟！"
	URL string `json:"url"` // ""
}

func (SysMsg) NoWrapper() {}

func (SysMsg) Cmd() string {
	return CmdSysMsg
}

type UserToastMsg struct {
	GuardLevel       int    `json:"guard_level"`        // 2
	OpType           int    `json:"op_type"`            // 1
	PayflowID        string `json:"payflow_id"`         // "2111011646287352179657661"
	Unit             string `json:"unit"`               // "月"
	IsShow           int    `json:"is_show"`            // 0
	Num              int    `json:"num"`                // 1
	Price            int    `json:"price"`              // 1998000
	StartTime        int    `json:"start_time"`         // 1635756388
	SvgaBlock        int    `json:"svga_block"`         // 0
	UserShow         bool   `json:"user_show"`          // true
	Color            string `json:"color"`              // "#E17AFF"
	EndTime          int    `json:"end_time"`           // 1635756388
	RoleName         string `json:"role_name"`          // "提督"
	ToastMsg         string `json:"toast_msg"`          // "<%何图灵%> 开通了提督"
	UID              int    `json:"uid"`                // 4497965
	AnchorShow       bool   `json:"anchor_show"`        // true
	Dmscore          int    `json:"dmscore"`            // 96
	TargetGuardCount int    `json:"target_guard_count"` // 12089
	Username         string `json:"username"`           // "何图灵"
}

func (UserToastMsg) Cmd() string {
	return CmdUserToastMsg
}

type VoiceJoinList struct {
	RoomID     int64 `json:"room_id"`     // 34348
	Category   int   `json:"category"`    // 1
	ApplyCount int   `json:"apply_count"` // 1
	RedPoint   int   `json:"red_point"`   // 1
	Refresh    int   `json:"refresh"`     // 1
}

func (VoiceJoinList) Cmd() string {
	return CmdVoiceJoinList
}

type VoiceJoinRoomCountInfo struct {
	ApplyCount  int `json:"apply_count"`  // 1
	NotifyCount int `json:"notify_count"` // 0
	RedPoint    int `json:"red_point"`    // 0
	RoomID      int `json:"room_id"`      // 34348
	RootStatus  int `json:"root_status"`  // 1
	RoomStatus  int `json:"room_status"`  // 1
}

func (VoiceJoinRoomCountInfo) Cmd() string {
	return CmdVoiceJoinRoomCountInfo
}

type WatChedChange struct {
	Num       int    `json:"num"`        // 144450
	TextLarge string `json:"text_large"` // 14.4万人看过
	TextSmall string `json:"text_small"` // 14.4万
}

func (WatChedChange) Cmd() string {
	return CmdWatChedChange
}

type WidgetBanner struct {
	Timestamp  int `json:"timestamp"` // 1735284354
	WidgetList map[int]struct {
		ID               int      `json:"id"`               // 1652
		Title            string   `json:"title"`            // "巅峰全站赛"
		Cover            string   `json:"cover"`            // ""
		WebCover         string   `json:"web_cover"`        // ""
		TipText          string   `json:"tip_text"`         // ""
		TipTextColor     string   `json:"tip_text_color"`   // ""
		TipBottomColor   string   `json:"tip_bottom_color"` // ""
		JumpURL          string   `json:"jump_url"`         // "https://live.bilibili.com/activity/live-activity-battle/index.html?is_live_half_webview=1&hybrid_rotate_d=1&hybrid_half_ui=1,3,100p,60p,0,0,0,0,12,0;2,2,375,100p,0,0,0,0,12,0;3,3,100p,60p,0,0,0,0,12,0;4,2,375,100p,0,0,0,0,12,0;5,3,100p,60p,0,0,0,0,12,0;6,3,100p,60p,0,0,0,0,12,0;7,3,100p,60p,0,0,0,0,12,0;8,3,100p,60p,0,0,0,0,12,0&pc_ui=375,626,17181a,2&room_id=21919321&uid=508963009&tab_id=pendant&selected_id=1652#/"
		URL              string   `json:"url"`              // ""
		StayTime         int      `json:"stay_time"`        // 5
		Site             int      `json:"site"`             // 1
		PlatformIn       []string `json:"platform_in"`
		Type             int      `json:"type"`              // 1
		BandID           int      `json:"band_id"`           // 107561
		SubKey           string   `json:"sub_key"`           // ""
		SubData          string   `json:"sub_data"`          // "%7B%22stage%22%3A1%2C%22is_hit_area%22%3Atrue%2C%22settlement%22%3A0%2C%22team_name%22%3A%22%E8%99%9A%E6%8B%9F%22%2C%22team_id%22%3A2%2C%22timestamp%22%3A1735284354%2C%22anchor_name%22%3A%22HiiroVTuber%22%2C%22head_img%22%3A%22https%3A%2F%2Fi1.hdslb.com%2Fbfs%2Fface%2F101e731be59fb1aaeddc28ba922c4aa7c83fa979.jpg%22%2C%22rank%22%3A%7B%22rank%22%3A462%2C%22score%22%3A24%2C%22rank_type%22%3A3%2C%22rank_distance_score%22%3A39854%2C%22timeout%22%3A0%7D%2C%22pk%22%3Anull%2C%22notice%22%3Anull%7D"
		IsAdd            bool     `json:"is_add"`            // true
		Position         int      `json:"position"`          // 0
		TargetURL        string   `json:"target_url"`        // "https://live.bilibili.com/activity/live-activity-battle/index.html?app_name=bls_winter_2024&config_id=1&show_refresh=0&show_close=0&is_live_half_webview=1&hybrid_rotate_d=1&hybrid_half_ui=1,3,100p,70p,0,0,0,0,12,0;2,2,375,100p,0,0,0,0,12,0;3,3,100p,70p,0,0,0,0,12,0;4,2,375,100p,0,0,0,0,12,0;5,3,100p,70p,0,0,0,0,12,0;6,3,100p,70p,0,0,0,0,12,0;7,3,100p,70p,0,0,0,0,12,0;8,3,100p,70p,0,0,0,0,12,0&merge_tab_id=1&room_id=21919321&uid=508963009#/all0287"
		URLType          int      `json:"url_type"`          // 0
		AndroidJumpurl   string   `json:"android_jumpurl"`   // ""
		IosJumpurl       string   `json:"ios_jumpurl"`       // ""
		AndroidSchameurl string   `json:"android_schameurl"` // ""
		IosSchemaurl     string   `json:"ios_schemaurl"`     // ""
		BizTime          int      `json:"biz_time"`          // 1735284354
		NotMergeTab      int      `json:"not_merge_tab"`     // 0
	} `json:"widget_list"`
}

func (WidgetBanner) Cmd() string {
	return CmdWidgetBanner
}
