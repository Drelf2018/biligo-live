package message

type WatChedChange struct {
	Num       int    `json:"num"`        //144450
	TextLarge string `json:"text_large"` //14.4万人看过
	TextSmall string `json:"text_small"` //14.4万
}

func (WatChedChange) Cmd() string {
	return CmdWatChedChange
}

var _ = Register[WatChedChange]()
