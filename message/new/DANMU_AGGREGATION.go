package message

const CmdDanmuAggregation = "DANMU_AGGREGATION"

type DanmuAggregation struct {
	ActivityIdentity string `json:"activity_identity"`  // "25109135"
	ActivitySource   int    `json:"activity_source"`    // 2
	AggregationCycle int    `json:"aggregation_cycle"`  // 1
	AggregationIcon  string `json:"aggregation_icon"`   // "https://i0.hdslb.com/bfs/live/024f7473753c7cc993413e05c69e8b960086e68f.png"
	AggregationNum   int    `json:"aggregation_num"`    // 11
	BroadcastMsgType int    `json:"broadcast_msg_type"` // 0
	Msg              string `json:"msg"`                // "老板大气！点点红包抽礼物"
	ShowRows         int    `json:"show_rows"`          // 1
	ShowTime         int    `json:"show_time"`          // 2
	Timestamp        int    `json:"timestamp"`          // 1735301889
}

func (DanmuAggregation) Cmd() string {
	return CmdDanmuAggregation
}
