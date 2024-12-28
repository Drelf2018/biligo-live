package message

const CmdRecommendCard = "RECOMMEND_CARD"

type RecommendCard struct {
	TitleIcon     string `json:"title_icon"` // "https://i0.hdslb.com/bfs/live/3053f47729c4974b1cfe4cd98482c28d4e23a1c2.png"
	RecommendList []any  `json:"recommend_list"`
	Timestamp     int    `json:"timestamp"` // 1735301684
	UpdateList    []any  `json:"update_list"`
}

func (RecommendCard) Cmd() string {
	return CmdRecommendCard
}
