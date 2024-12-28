package message

const CmdGotoBuyFlow = "GOTO_BUY_FLOW"

type GotoBuyFlow struct {
	Text string `json:"text"` // "L**正在去买"
}

func (GotoBuyFlow) Cmd() string {
	return CmdGotoBuyFlow
}
