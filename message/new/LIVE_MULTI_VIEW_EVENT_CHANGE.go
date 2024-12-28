package message

const CmdLiveMultiViewEventChange = "LIVE_MULTI_VIEW_EVENT_CHANGE"

type LiveMultiViewEventChange []struct {
	RoomID        int    `json:"room_id"`         // 21144080
	MatchStatus   int    `json:"match_status"`    // 2
	HomeTeamName  string `json:"home_team_name"`  // "济南RW侠"
	AwayTeamName  string `json:"away_team_name"`  // "北京WB"
	HomeTeamIcon  string `json:"home_team_icon"`  // "https://i0.hdslb.com//bfs/legacy/ebeab9f37ede1658083e612bd4d534def15acc70.png"
	AwayTeamIcon  string `json:"away_team_icon"`  // "https://i0.hdslb.com//bfs/feed-admin/5e4dcfbe906c7e128190bc4a92bf6b0effc125da.png"
	HomeTeamScore int    `json:"home_team_score"` // 1
	AwayTeamScore int    `json:"away_team_score"` // 0
	TimeStamp     int    `json:"time_stamp"`      // 1735299375
}

func (LiveMultiViewEventChange) Cmd() string {
	return CmdLiveMultiViewEventChange
}
