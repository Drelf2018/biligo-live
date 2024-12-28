package message

const CmdCommonNoticeDanmaku = "COMMON_NOTICE_DANMAKU"

type CommonNoticeDanmaku struct {
	ContentSegments []struct {
		FontColor              string `json:"font_color"`                // "#CCCCCC"
		FontColorDark          string `json:"font_color_dark"`           // "#CCCCCC"
		HighlightFontColor     string `json:"highlight_font_color"`      // "#FFC73E"
		HighlightFontColorDark string `json:"highlight_font_color_dark"` // "#FFC73E"
		Text                   string `json:"text"`                      // "恭喜用户 板斧青凤 <%荣耀等级升级至11级%>"
		Type                   int    `json:"type"`                      // 1
	} `json:"content_segments"`
	Dmscore   int   `json:"dmscore"` // 1008
	Terminals []int `json:"terminals"`
}

func (CommonNoticeDanmaku) Cmd() string {
	return CmdCommonNoticeDanmaku
}
