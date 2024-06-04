package message

type HotRankChanged struct {
	Rank        int    `json:"rank"`
	Timestamp   int64  `json:"timestamp"`
	WebUrl      string `json:"web_url"`
	LiveUrl     string `json:"live_url"`
	LiveLinkUrl string `json:"live_link_url"`
	AreaName    string `json:"area_name"`
	Trend       int    `json:"trend"`
	Countdown   int    `json:"countdown"`
	BlinkUrl    string `json:"blink_url"`
	PCLinkUrl   string `json:"pc_link_url"`
	Icon        string `json:"icon"`
}

func (HotRankChanged) Cmd() string {
	return CmdHotRankChanged
}

var _ = Register[HotRankChanged]()

/*
	{
		"dm_msg": "恭喜主播 <% 老实憨厚的笑笑 %> 荣登限时热门榜网游榜top3! 即将获得热门流量推荐哦！",
		"dmscore": 144,
		"timestamp": 1635748200,
		"uname": "老实憨厚的笑笑",
		"url": "https://live.bilibili.com/p/html/live-app-hotrank/result.html?is_live_half_webview=1&hybrid_half_ui=1,5,250,200,f4eefa,0,30,0,0,0;2,5,250,200,f4eefa,0,30,0,0,0;3,5,250,200,f4eefa,0,30,0,0,0;4,5,250,200,f4eefa,0,30,0,0,0;5,5,250,200,f4eefa,0,30,0,0,0;6,5,250,200,f4eefa,0,30,0,0,0;7,5,250,200,f4eefa,0,30,0,0,0;8,5,250,200,f4eefa,0,30,0,0,0&areaId=2&cache_key=d04fe4e6d0f2bc24c1a5acc63f574b68",
		"area_name": "网游",
		"cache_key": "d04fe4e6d0f2bc24c1a5acc63f574b68",
		"rank": 3,
		"face": "http://i0.hdslb.com/bfs/face/d92282ac664afffd0ef38c275f3fca17a9567d5a.jpg",
		"icon": "https://i0.hdslb.com/bfs/live/65dbe013f7379c78fc50dfb2fd38d67f5e4895f9.png"
	}
*/
type HotRankSettlement struct {
	DmMsg     string `json:"dm_msg"`
	DmScore   int    `json:"dmscore"`
	Timestamp int64  `json:"timestamp"`
	Uname     string `json:"uname"`
	Url       string `json:"url"`
	AreaName  string `json:"area_name"`
	CacheKey  string `json:"cache_key"`
	Rank      int    `json:"rank"`
	Face      string `json:"face"`
	Icon      string `json:"icon"`
}

func (HotRankSettlement) Cmd() string {
	return CmdHotRankSettlement
}

var _ = Register[HotRankSettlement]()
