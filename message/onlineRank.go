package message

/*
	{
		"count": 5679,
		"count_text": "5679",
		"online_count": 6455,
		"online_count_text": "6455"
	}
*/
type OnlineRankCount struct {
	Count           int    `json:"count"`
	CountText       string `json:"count_text"`
	OnlineCount     int    `json:"online_count,omitempty"`
	OnlineCountText string `json:"online_count_text,omitempty"`
}

func (OnlineRankCount) Cmd() string {
	return CmdOnlineRankCount
}

var _ = Register[OnlineRankCount]()

/*
	{
		"dmscore": 112,
		"list": [{
			"msg": "恭喜 <%你们有多腐%> 成为高能榜",
			"rank": 2
		}]
	}
*/
type OnlineRankTop3 struct {
	DmScore int `json:"dmscore"`
	List    []struct {
		Msg  string `json:"msg"`
		Rank int    `json:"rank"`
	} `json:"list"`
}

func (OnlineRankTop3) Cmd() string {
	return CmdOnlineRankTop3
}

var _ = Register[OnlineRankTop3]()

/*
	{
		"list": [
			{
				"guard_level": 2,
				"uid": 277278853,
				"face": "http://i1.hdslb.com/bfs/face/59839130848b8f8d99f8c649f7897ac7f406a052.jpg",
				"score": "15980",
				"uname": "勤俭持家的席撒",
				"rank": 1
			},
			{
				"uid": 12777723,
				"face": "http://i0.hdslb.com/bfs/face/6badd87b9bf8c13c90fcb2c2b1b93b01e4b02664.jpg",
				"score": "2500",
				"uname": "卡纸哥我宣你",
				"rank": 2,
				"guard_level": 2
			},
			{
				"uid": 19229891,
				"face": "http://i2.hdslb.com/bfs/face/3925926f11983d7c2e1736e429aa171761493040.jpg",
				"score": "1580",
				"uname": "大象12183",
				"rank": 3,
				"guard_level": 3
			},
			{
				"rank": 4,
				"guard_level": 3,
				"uid": 271376887,
				"face": "http://i1.hdslb.com/bfs/face/84ef7024aef33ad5d790494130c4081e3a872169.jpg",
				"score": "1380",
				"uname": "w蓄意轰拳w"
			},
			{
				"face": "http://i0.hdslb.com/bfs/face/b0d4640c49ef04f630b103edbec1a277b912fbe1.jpg",
				"score": "1000",
				"uname": "Dys莫的命",
				"rank": 5,
				"guard_level": 3,
				"uid": 16495374
			},
			{
				"uid": 601557387,
				"face": "http://i0.hdslb.com/bfs/face/e5cb2f45e257f337c756521bd73c56814443c8c0.jpg",
				"score": "1000",
				"uname": "秃了送",
				"rank": 6,
				"guard_level": 3
			},
			{
				"guard_level": 3,
				"uid": 143379249,
				"face": "http://i0.hdslb.com/bfs/face/1d581ce73feb42a6d73839047d781c434652195b.jpg",
				"score": "1000",
				"uname": "小泉水噗噗",
				"rank": 7
			}
		],
		"rank_type": "gold-rank"
	}
*/
type OnlineRankV2 struct {
	List []struct {
		GuardLevel int    `json:"guard_level"` // 3:舰长 2:提督 1:总督?
		UID        int64  `json:"uid"`
		Face       string `json:"face"`
		Score      string `json:"score"`
		Uname      string `json:"uname"`
		Rank       int    `json:"rank"`
	} `json:"list"`
	RankType string `json:"rank_type"`
}

func (OnlineRankV2) Cmd() string {
	return CmdOnlineRankV2
}

var _ = Register[OnlineRankV2]()
