package message

const CmdGiftStarProcess = "GIFT_STAR_PROCESS"

type GiftStarProcess struct {
	Status int    `json:"status"` // 1
	Tip    string `json:"tip"`    // "为你打call已点亮"
}

func (GiftStarProcess) Cmd() string {
	return CmdGiftStarProcess
}
