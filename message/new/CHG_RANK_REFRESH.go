package message

const CmdChgRankRefresh = "CHG_RANK_REFRESH"

type ChgRankRefresh struct {
	RawCmd      string `json:"cmd"`          // "CHG_RANK_REFRESH"
	RankType    int    `json:"rank_type"`    // 4
	RankModule  string `json:"rank_module"`  // "area"
	RoomID      int    `json:"room_id"`      // 47867
	Ruid        int    `json:"ruid"`         // 67141
	NeedRefresh bool   `json:"need_refresh"` // true
	Version     int64  `json:"version"`      // 1735297046013
}

func (ChgRankRefresh) Cmd() string {
	return CmdChgRankRefresh
}
