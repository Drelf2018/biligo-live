package message

/*
	{
		"cmd": "NOTICE_MSG",
		"business_id": "31087",
		"full": {
			"head_icon": "http://i0.hdslb.com/bfs/live/00f26756182b2e9d06c00af23001bc8e10da67d0.webp",
			"tail_icon": "http://i0.hdslb.com/bfs/live/822da481fdaba986d738db5d8fd469ffa95a8fa1.webp",
			"head_icon_fa": "http://i0.hdslb.com/bfs/live/77983005023dc3f31cd599b637c83a764c842f87.png",
			"tail_icon_fa": "http://i0.hdslb.com/bfs/live/38cb2a9f1209b16c0f15162b0b553e3b28d9f16f.png",
			"background": "#6098FFFF",
			"highlight": "#FDFF2FFF",
			"head_icon_fan": 36,
			"tail_icon_fan": 4,
			"color": "#FFFFFFFF",
			"time": 20
		},
		"half": {
			"time": 15,
			"head_icon": "http://i0.hdslb.com/bfs/live/358cc52e974b315e83eee429858de4fee97a1ef5.png",
			"tail_icon": "",
			"background": "#7BB6F2FF",
			"color": "#FFFFFFFF",
			"highlight": "#FDFF2FFF"
		},
		"id": 2,
		"link_url": "https://live.bilibili.com/5655865?accept_quality=%5B10000%2C150%5D&broadcast_type=0&current_qn=150&current_quality=150&is_room_feed=1&live_play_network=other&p2p_type=-2&playurl_h264=http%3A%2F%2Fd1--cn-gotcha03.bilivideo.com%2Flive-bvc%2F429443%2Flive_2257663_5953069_1500.flv%3Fexpires%3D1635753433%26len%3D0%26oi%3D0%26pt%3D%26qn%3D150%26trid%3D10004aaecf5169e74b51b5932933468e0364%26sigparams%3Dcdn%2Cexpires%2Clen%2Coi%2Cpt%2Cqn%2Ctrid%26cdn%3Dcn-gotcha03%26sign%3De0b8728896efe026833d99655b05c084%26p2p_type%3D4294967294%26src%3D5%26sl%3D1%26flowtype%3D1%26source%3Dbatch%26order%3D1%26machinezone%3Dylf%26sk%3D2935686d6cb9146c7a6a6a0b4e120e2594e074fa0760377f1a7a2b2fa0ee6443&playurl_h265=&quality_description=%5B%7B%22qn%22%3A10000%2C%22desc%22%3A%22%E5%8E%9F%E7%94%BB%22%7D%2C%7B%22qn%22%3A150%2C%22desc%22%3A%22%E9%AB%98%E6%B8%85%22%7D%5D&from=28003&extra_jump_from=28003&live_lottery_type=1",
		"msg_common": "<%JamesTuT%>投喂:<%木之本切%>1个次元之城，点击前往TA的房间吧！",
		"msg_self": "<%JamesTuT%>投喂:<%木之本切%>1个次元之城，快来围观吧！",
		"msg_type": 2,
		"name": "分区道具抽奖广播样式",
		"real_roomid": 5655865,
		"roomid": 5655865,
		"scatter": {
			"min": 0,
			"max": 0
		},
		"shield_uid": -1,
		"side": {
			"head_icon": "",
			"background": "",
			"color": "",
			"highlight": "",
			"border": ""
		}
	}
*/
type NoticeMsg struct {
	BusinessID string `json:"business_id"`
	Full       struct {
		HeadIcon    string `json:"head_icon"`
		TailIcon    string `json:"tail_icon"`
		HeadIconFa  string `json:"head_icon_fa"`
		TailIconFa  string `json:"tail_icon_fa"`
		Background  string `json:"background"`
		Highlight   string `json:"highlight"`
		HeadIconFan int    `json:"head_icon_fan"`
		TailIconFan int    `json:"tail_icon_fan"`
		Color       string `json:"color"`
		Time        int64  `json:"time"`
	} `json:"full"`
	Half struct {
		Time       int64  `json:"time"`
		HeadIcon   string `json:"head_icon"`
		TailIcon   string `json:"tail_icon"`
		Background string `json:"background"`
		Color      string `json:"color"`
		Highlight  string `json:"highlight"`
	} `json:"half"`
	ID         int64  `json:"id"`
	LinkUrl    string `json:"link_url"`
	MsgCommon  string `json:"msg_common"`
	MsgSelf    string `json:"msg_self"`
	MsgType    int    `json:"msg_type"`
	Name       string `json:"name"`
	RealRoomID int64  `json:"real_roomid"`
	RoomID     int64  `json:"roomid"`
	Scatter    struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"scatter"`
	ShieldUID int64 `json:"shield_uid"`
	Side      struct {
		HeadIcon   string `json:"head_icon"`
		Background string `json:"background"`
		Color      string `json:"color"`
		Highlight  string `json:"highlight"`
		Border     string `json:"border"`
	} `json:"side"`
}

func (NoticeMsg) Cmd() string {
	return CmdNoticeMsg
}

var _ = RegisterMsg[NoticeMsg](true)

/*
	{
		"notice_msg": "为保护用户隐私，未登录无法查看他人昵称",
		"image_web": "http://i0.hdslb.com/bfs/dm/75e7c16b99208df259fe0a93354fd3440cbab412.png",
		"image_app": "http://i0.hdslb.com/bfs/dm/b632f7dcd3acf47deffb5f9ccc9546ae97a3415b.png"
	}
*/
type LogInNotice struct {
	NoticeMsg string `json:"notice_msg"`
	ImageWeb  string `json:"image_web"`
	ImageApp  string `json:"image_app"`
}

func (LogInNotice) Cmd() string {
	return CmdLogInNotice
}

var _ = Register[LogInNotice]()
