package message

var _ = Register(FansUpdate{})

type FansUpdate struct {
	// {
	// 	"fans_club": 49182,
	// 	"roomid": 545068,
	// 	"fans": 1384297,
	// 	"red_notice": -1
	// }
	FansClub  int   `json:"fans_club"`
	RoomID    int64 `json:"roomid"`
	Fans      int   `json:"fans"`
	RedNotice int   `json:"red_notice"`
}

func (FansUpdate) Cmd() string {
	return CmdRoomRealTimeMessageUpdate
}

func (FansUpdate) Parse(raw []byte) (Msg, error) {
	return GetData[FansUpdate](raw)
}
