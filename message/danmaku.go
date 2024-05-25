package message

import "encoding/json"

var _ = Register(Danmaku{})

type Danmaku struct {
	SendMode     int    `json:"send_mode"`
	SendFontSize int    `json:"send_font_size"`
	DanmakuColor int64  `json:"danmaku_color"`
	Time         int64  `json:"time"`
	DMID         int64  `json:"dmid"`
	MsgType      int    `json:"msg_type"`
	Bubble       string `json:"bubble"`
	Content      string `json:"content"`
	MID          int64  `json:"mid"`
	Uname        string `json:"uname"`
	RoomAdmin    int    `json:"room_admin"`
	Vip          int    `json:"vip"`
	SVip         int    `json:"svip"`
	Rank         int    `json:"rank"`
	MobileVerify int    `json:"mobile_verify"`
	UnameColor   string `json:"uname_color"`
	MedalName    string `json:"medal_name"`
	UpName       string `json:"up_name"`
	MedalLevel   int    `json:"medal_level"`
	UserLevel    int    `json:"user_level"`
}

func (Danmaku) Cmd() string {
	return CmdDanmaku
}

func (dm Danmaku) Parse(raw []byte) (Msg, error) {
	var data struct {
		Info []any `json:"info"`
	}

	err := json.Unmarshal(raw, &data)
	if err != nil {
		return nil, err
	}

	info := data.Info
	l := len(info)
	if l > 5 {
		l = 5
	}

	switch l {
	case 5:
		dm.UserLevel = int(info[4].([]any)[0].(float64))
		fallthrough
	case 4:
		h := info[3].([]any)
		l2 := len(h)
		if l2 >= 1 {
			dm.MedalLevel = int(h[0].(float64))
		}
		if l2 >= 2 {
			dm.MedalName = h[1].(string)
		}
		if l2 >= 3 {
			dm.UpName = h[2].(string)
		}
		fallthrough
	case 3:
		h := info[2].([]any)
		dm.MID = int64(h[0].(float64))
		dm.Uname = h[1].(string)
		dm.RoomAdmin = int(h[2].(float64))
		dm.Vip = int(h[3].(float64))
		dm.SVip = int(h[4].(float64))
		dm.Rank = int(h[5].(float64))
		dm.MobileVerify = int(h[6].(float64))
		dm.UnameColor = h[7].(string)
		fallthrough
	case 2:
		dm.Content = info[1].(string)
		fallthrough
	case 1:
		h := info[0].([]any)
		dm.SendMode = int(h[1].(float64))
		dm.SendFontSize = int(h[2].(float64))
		dm.DanmakuColor = int64(h[3].(float64))
		dm.Time = int64(h[4].(float64))
		dm.DMID = int64(h[5].(float64))
		dm.MsgType = int(h[10].(float64))
		dm.Bubble = h[11].(string)
	}

	return dm, nil
}
