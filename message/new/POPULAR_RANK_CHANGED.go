package message

const CmdPopularRankChanged = "POPULAR_RANK_CHANGED"

type PopularRankChanged struct {
	UID              int    `json:"uid"`                  // 67141
	Rank             int    `json:"rank"`                 // 0
	Countdown        int    `json:"countdown"`            // 2701
	Timestamp        int    `json:"timestamp"`            // 1735298100
	CacheKey         string `json:"cache_key"`            // "rank_change:da11cc989d764bc1b28042f3b272f058"
	OnRankNameByType string `json:"on_rank_name_by_type"` // ""
	RankNameByType   string `json:"rank_name_by_type"`    // "人气榜"
	URLByType        string `json:"url_by_type"`          // "https://live.bilibili.com/p/html/live-app-hotrank/index.html?is_live_half_webview=1&hybrid_rotate_d=1&hybrid_half_ui=1,3,100p,70p,0,0,30,100,12;2,2,375,100p,0,0,30,100,0;3,3,100p,70p,0,0,30,100,12;4,2,375,100p,0,0,30,100,0;5,3,100p,70p,0,0,30,100,0;6,3,100p,70p,0,0,30,100,0;7,3,100p,70p,0,0,30,100,0;8,3,100p,70p,0,0,30,100,0&pc_ui=338,465,f4eefa,0&redirect=v2&anchorId=67141"
	RankByType       int    `json:"rank_by_type"`         // 0
	DefaultURL       string `json:"default_url"`          // "https://live.bilibili.com/p/html/live-app-hotrank/index.html?is_live_half_webview=1&hybrid_rotate_d=1&hybrid_half_ui=1,3,100p,70p,0,0,30,100,12;2,2,375,100p,0,0,30,100,0;3,3,100p,70p,0,0,30,100,12;4,2,375,100p,0,0,30,100,0;5,3,100p,70p,0,0,30,100,0;6,3,100p,70p,0,0,30,100,0;7,3,100p,70p,0,0,30,100,0;8,3,100p,70p,0,0,30,100,0&pc_ui=338,465,f4eefa,0&redirect=v2&anchorId=67141"

}

func (PopularRankChanged) Cmd() string {
	return CmdPopularRankChanged
}
