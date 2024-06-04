package message

import (
	"encoding/json"
)

/*
	{
		"data": "{\"combo\":[{\"id\":37699755023360,\"status\":4,\"content\":\"人妻\",\"cnt\":25,\"guide\":\"他们都在说:\",\"left_duration\":6000,\"fade_duration\":10000,\"prefix_icon\":\"\"},{\"id\":37699834614784,\"status\":4,\"content\":\"？\",\"cnt\":5,\"guide\":\"他们都在说:\",\"left_duration\":19000,\"fade_duration\":10000,\"prefix_icon\":\"\"}],\"merge_interval\":1000,\"card_appear_interval\":1000,\"send_interval\":1000,\"reset_cnt\":1,\"display_flag\":0}",
		"dmscore": 36,
		"id": 37699755023360,
		"status": 4,
		"type": 102
	}
*/
type DmInteraction struct {
	DataRaw string `json:"data"`
	Dmscore int    `json:"dmscore"`
	ID      int64  `json:"id"`
	Status  int    `json:"status"`
	Type    int    `json:"type"`
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

var _ = Register[DmInteraction]()
