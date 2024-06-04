package message

/*
	{
		"from": "fans_medal",
		"platform": {
		"android": true,
		"ios": true,
		"web": true
		},
		"msg": "今日首条弹幕发送成功~亲密度+100",
		"room_id": 8432038,
		"type": 1
	}
*/
type LittleMessageBox struct {
	From     string `json:"from"`
	Platform struct {
		Android bool `json:"android"`
		Ios     bool `json:"ios"`
		Web     bool `json:"web"`
	} `json:"platform"`
	Msg    string `json:"msg"`
	RoomID int    `json:"room_id"`
	Type   int    `json:"type"`
}

func (LittleMessageBox) Cmd() string {
	return CmdLittleMessageBox
}

var _ = Register[LittleMessageBox]()

/*
	{
		"type": 2,
		"uid": 188888131,
		"up_uid": 267766441,
		"medal_level": 15,
		"medal_name": "全撤了",
		"medal_color_start": 12478086,
		"medal_color_end": 12478086,
		"medal_color_border": 12478086,
		"is_lighted": 1,
		"is_lighted_v2": true,
		"guard_level": 0,
		"unlock": 0,
		"unlock_level": 0,
		"multi_unlock_level": "",
		"upper_bound_content": ""
	}
*/
type MessageboxUserMedalChange struct {
	Type              int    `json:"type"`
	UID               int    `json:"uid"`
	UpUID             int    `json:"up_uid"`
	MedalLevel        int    `json:"medal_level"`
	MedalName         string `json:"medal_name"`
	MedalColorStart   int    `json:"medal_color_start"`
	MedalColorEnd     int    `json:"medal_color_end"`
	MedalColorBorder  int    `json:"medal_color_border"`
	IsLighted         int    `json:"is_lighted"`
	IsLightedV2       bool   `json:"is_lighted_v2"`
	GuardLevel        int    `json:"guard_level"`
	Unlock            int    `json:"unlock"`
	UnlockLevel       int    `json:"unlock_level"`
	MultiUnlockLevel  string `json:"multi_unlock_level"`
	UpperBoundContent string `json:"upper_bound_content"`
}

func (MessageboxUserMedalChange) Cmd() string {
	return CmdMessageboxUserMedalChange
}

var _ = Register[MessageboxUserMedalChange]()
