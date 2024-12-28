package message

const CmdPopularityRedPocketWinnerList = "POPULARITY_RED_POCKET_WINNER_LIST"

type PopularityRedPocketWinnerList struct {
	LotID      int   `json:"lot_id"`    // 25109135
	TotalNum   int   `json:"total_num"` // 21
	AwardNum   int   `json:"award_num"` // 21
	WinnerInfo []any `json:"winner_info"`
	Awards     struct {
		Num31213 struct {
			AwardType   int    `json:"award_type"`    // 1
			AwardName   string `json:"award_name"`    // "这个好诶"
			AwardPic    string `json:"award_pic"`     // "https://s1.hdslb.com/bfs/live/9260c680959428c45b3a2742e42ea7ae75e457ef.png"
			AwardBigPic string `json:"award_big_pic"` // "https://i0.hdslb.com/bfs/live/dafd2e9e5c3086377124a9328e840cd21a3f6847.png"
			AwardPrice  int    `json:"award_price"`   // 1000
		} `json:"31213"`
		Num31212 struct {
			AwardType   int    `json:"award_type"`    // 1
			AwardName   string `json:"award_name"`    // "打call"
			AwardPic    string `json:"award_pic"`     // "https://s1.hdslb.com/bfs/live/461be640f60788c1d159ec8d6c5d5cf1ef3d1830.png"
			AwardBigPic string `json:"award_big_pic"` // "https://i0.hdslb.com/bfs/live/9e6521c57f24c7149c054d265818d4b82059f2ef.png"
			AwardPrice  int    `json:"award_price"`   // 500
		} `json:"31212"`
		Num34003 struct {
			AwardType   int    `json:"award_type"`    // 1
			AwardName   string `json:"award_name"`    // "人气票"
			AwardPic    string `json:"award_pic"`     // "https://s1.hdslb.com/bfs/live/7164c955ec0ed7537491d189b821cc68f1bea20d.png"
			AwardBigPic string `json:"award_big_pic"` // "https://i0.hdslb.com/bfs/live/5bfaddf9a78e677501bb6d440f4d690668136496.png"
			AwardPrice  int    `json:"award_price"`   // 100
		} `json:"34003"`
	} `json:"awards"`
	Version   int `json:"version"`   // 1
	RpType    int `json:"rp_type"`   // 0
	Timestamp int `json:"timestamp"` // 1735302061
}

func (PopularityRedPocketWinnerList) Cmd() string {
	return CmdPopularityRedPocketWinnerList
}
