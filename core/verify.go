package core

import (
	"strconv"

	api "github.com/Drelf2018/go-bilibili-api"
)

// 用户认证数据
type VerifyData struct {
	// 自己的 UID
	UID int `json:"uid"`

	// 目标用户直播间号
	RoomID int `json:"roomid"`

	// 协议版本 通常用 3
	Protover int `json:"protover"`

	// 用户设备信息
	Buvid string `json:"buvid"`

	// 平台 通常用 web
	Platform string `json:"platform"`

	// 类型 通常用 2
	Type int `json:"type"`

	// 用户私钥 要求与 Buvid 匹配
	Key string `json:"key"`
}

// 新建认证数据
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

// 获取认证数据
func GetVerifyData(roomid int, credential *api.Credential) (v VerifyData, err error) {
	// 获取字符格式 UID
	var uid int
	if credential != nil {
		uid, err = strconv.Atoi(credential.DedeUserID)
		if err != nil {
			return
		}
	}
	// 将短号转为真实直播间号
	room, err := api.GetRoomInfo(roomid)
	if err != nil {
		return
	}
	roomid = room.Data.RoomID
	// 获取设备信息
	var buvid3 string
	if credential != nil && credential.Buvid3 != "" {
		buvid3 = credential.Buvid3
	} else {
		var spi api.SPIResponse
		spi, err = api.GetSPI()
		if err != nil {
			return
		}
		buvid3 = spi.Data.B3
	}
	// 获取匹配的密钥
	info, err := api.GetDanmuInfo(roomid, credential)
	if err != nil {
		return
	}
	// 拿下
	return NewVerifyData(uid, roomid, buvid3, info.Data.Token), nil
}
