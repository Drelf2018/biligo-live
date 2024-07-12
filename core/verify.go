package core

type VerifyData struct {
	// 自己的 UID
	UID int `json:"uid"`
	// 台主房间号
	RoomID int `json:"roomid"`
	// 协议版本 通常用 3
	Protover int `json:"protover"`
	// 用于区分用户的字符串
	Buvid string `json:"buvid"`
	// 平台 通常用 web
	Platform string `json:"platform"`
	// 类型 通常用 2
	Type int `json:"type"`
	// 用户私钥 与 Buvid 对应
	Key string `json:"key"`
}

func NewVerifyData(uid, room int, buvid, key string) VerifyData {
	return VerifyData{
		UID:      uid,
		RoomID:   room,
		Protover: 3,
		Buvid:    buvid,
		Platform: "web",
		Type:     2,
		Key:      key,
	}
}
