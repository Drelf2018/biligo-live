package message

/*
	{
		"roomid": "14703541",
		"round": 1,
		"cmd": "PREPARING"
	}
*/
type Preparing struct {
	RoomID string `json:"roomid"`
	Round  int    `json:"round"`
}

func (Preparing) Cmd() string {
	return CmdPreparing
}

var _ = RegisterMsg[Preparing](true)
