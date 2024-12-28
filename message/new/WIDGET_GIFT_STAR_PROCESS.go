package message

const CmdWidgetGiftStarProcess = "WIDGET_GIFT_STAR_PROCESS"

type WidgetGiftStarProcess struct {
	StartDate   int `json:"start_date"` // 20241223
	ProcessList []struct {
		GiftID       int    `json:"gift_id"`       // 34383
		GiftImg      string `json:"gift_img"`      // "https://s1.hdslb.com/bfs/live/f1ed9d627f75b46fe621558951847e0e4e95d6a2.png"
		GiftName     string `json:"gift_name"`     // "礼物星球"
		CompletedNum int    `json:"completed_num"` // 0
		TargetNum    int    `json:"target_num"`    // 1
	} `json:"process_list"`
	Finished       bool   `json:"finished"`         // false
	DdlTimestamp   int    `json:"ddl_timestamp"`    // 1735488000
	Version        int64  `json:"version"`          // 1735294008274
	RewardGift     int    `json:"reward_gift"`      // 0
	RewardGiftImg  string `json:"reward_gift_img"`  // "https://i0.hdslb.com/bfs/live/52edb4ab7377ece34ac15b21154d13d188874b01.png"
	RewardGiftName string `json:"reward_gift_name"` // "礼物星球"
	LevelInfo      struct {
		StarName string `json:"star_name"` // "礼物星球"
		LevelTip string `json:"level_tip"` // "暂未达成"
		LevelImg string `json:"level_img"` // "https://i0.hdslb.com/bfs/live/91292553d02aa10a133cdf6d8e579d8588e2067a.png"
		LevelID  int    `json:"level_id"`  // 0
	} `json:"level_info"`
}

func (WidgetGiftStarProcess) Cmd() string {
	return CmdWidgetGiftStarProcess
}
