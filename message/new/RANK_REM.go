package message

const CmdRankRem = "RANK_REM"

type RankRem struct {
	Name   string `json:"name"`    // "online_rank"
	RoomID int    `json:"room_id"` // 47867
	Ruid   int    `json:"ruid"`    // 67141
	Time   int    `json:"time"`    // 1735291145
	UID    int    `json:"uid"`     // 38870608

}

func (RankRem) Cmd() string {
	return CmdRankRem
}
