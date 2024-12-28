package message

const CmdPopularityRankTabChg = "POPULARITY_RANK_TAB_CHG"

type PopularityRankTabChg struct {
	RoomID         int    `json:"room_id"`          // 47867
	Ruid           int    `json:"ruid"`             // 67141
	Type           string `json:"type"`             // "area"
	NeedRefreshTab bool   `json:"need_refresh_tab"` // true

}

func (PopularityRankTabChg) Cmd() string {
	return CmdPopularityRankTabChg
}
