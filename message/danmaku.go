package message

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tidwall/gjson"
)

/*
	{
		"cmd": "DANMU_MSG",
		"info": [
			[
				0,
				1,
				25,
				16777215,
				1717143587990,
				927795592,
				0,
				"4c75c82a",
				0,
				0,
				0,
				"",
				0,
				"{}",
				"{}",
				{
					"mode": 0,
					"show_player_type": 0,
					"extra": "{\"send_from_me\":false,\"mode\":0,\"color\":16777215,\"dm_type\":0,\"font_size\":25,\"player_mode\":1,\"show_player_type\":0,\"content\":\"不上niko小孩真赢不了吧[dog]\",\"user_hash\":\"1282787370\",\"emoticon_unique\":\"\",\"bulge_display\":0,\"recommend_score\":7,\"main_state_dm_color\":\"\",\"objective_state_dm_color\":\"\",\"direction\":0,\"pk_direction\":0,\"quartet_direction\":0,\"anniversary_crowd\":0,\"yeah_space_type\":\"\",\"yeah_space_url\":\"\",\"jump_to_url\":\"\",\"space_type\":\"\",\"space_url\":\"\",\"animation\":{},\"emots\":{\"[dog]\":{\"count\":1,\"descript\":\"[dog]\",\"emoji\":\"[dog]\",\"emoticon_id\":208,\"emoticon_unique\":\"emoji_208\",\"height\":20,\"url\":\"http://i0.hdslb.com/bfs/live/4428c84e694fbf4e0ef6c06e958d9352c3582740.png\",\"width\":20}},\"is_audited\":false,\"id_str\":\"3868686b51d7fa6e6bbbff8bff66598841\",\"icon\":null,\"show_reply\":true,\"reply_mid\":0,\"reply_uname\":\"\",\"reply_uname_color\":\"\",\"reply_is_mystery\":false,\"hit_combo\":0}",
					"user": {
						"uid": 30629377,
						"base": {
							"name": "无名之王QAQ",
							"face": "https://i1.hdslb.com/bfs/face/023e15e625122373a155db37231ad8674faaf488.jpg",
							"name_color": 0,
							"is_mystery": false,
							"risk_ctrl_info": null,
							"origin_info": {
								"name": "无名之王QAQ",
								"face": "https://i1.hdslb.com/bfs/face/023e15e625122373a155db37231ad8674faaf488.jpg"
							},
							"official_info": {
								"role": 0,
								"title": "",
								"desc": "",
								"type": -1
							},
							"name_color_str": ""
						},
						"medal": null,
						"wealth": null,
						"title": {
							"old_title_css_id": "",
							"title_css_id": ""
						},
						"guard": null,
						"uhead_frame": null,
						"guard_leader": {
							"is_guard_leader": false
						}
					}
				},
				{
					"activity_identity": "",
					"activity_source": 0,
					"not_show": 0
				},
				0
			],
			"不上niko小孩真赢不了吧[dog]",
			[
				30629377,
				"无名之王QAQ",
				0,
				0,
				0,
				10000,
				1,
				""
			],
			[
				24,
				"小孩梓",
				"阿梓从小就很可爱",
				80397,
				1725515,
				"",
				0,
				1725515,
				1725515,
				5414290,
				0,
				1,
				7706705
			],
			[
				13,
				0,
				6406234,
				"\u003e50000",
				0
			],
			[
				"",
				""
			],
			0,
			0,
			null,
			{
				"ts": 1717143587,
				"ct": "A21C9A93"
			},
			0,
			0,
			null,
			null,
			0,
			130,
			[
				0
			],
			null
		],
		"dm_v2": ""
	}
*/

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
	CmdRaw string      `json:"cmd"`
	Info   DanmakuInfo `json:"info"`
	DmV2   string      `json:"dm_v2"`
}

func (Danmaku) Cmd() string {
	return CmdDanmaku
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

var _ = RegisterMsg[Danmaku](true)
var _ = RegisterCmd(CmdDanmakuWithCode, UnmarshalDirect[Danmaku])
