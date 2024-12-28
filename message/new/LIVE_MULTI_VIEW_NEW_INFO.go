package message

const CmdLiveMultiViewNewInfo = "LIVE_MULTI_VIEW_NEW_INFO"

type LiveMultiViewNewInfo struct {
	Title        string `json:"title"`          // "尖叫浴场"
	RoomID       int    `json:"room_id"`        // 21470454
	CopyWriting  string `json:"copy_writing"`   // "更多视角"
	BgImage      string `json:"bg_image"`       // "https://i0.hdslb.com/bfs/live/edaa9477a1d8325dd0c36c419b6fd5f9646b2419.png"
	SubSltColor  string `json:"sub_slt_color"`  // "#FFFFFF"
	SubBgColor   string `json:"sub_bg_color"`   // "#333333"
	SubTextColor string `json:"sub_text_color"` // "#FFFFFF"
	ViewType     int    `json:"view_type"`      // 0
	RoomList     []struct {
		OrderID    int    `json:"order_id"`    // 1
		RoomID     int    `json:"room_id"`     // 21470454
		RoomName   string `json:"room_name"`   // "主视角"
		LiveStatus int    `json:"live_status"` // 1
		JumpURL    string `json:"jump_url"`    // "https://live.bilibili.com/21470454?accept_quality=%5B10000%2C250%5D&av1_current_qn=0&broadcast_type=0&current_qn=250&current_quality=250&h264_current_qn=250&h265_current_qn=250&is_room_feed=1&live_play_network=other&master_url=&p2p_type=-1&playurl_av1=&playurl_h264=http%3A%2F%2Fd1--cn-gotcha04.bilivideo.com%2Flive-bvc%2F785333%2Flive_21470454_bs_2532759_2500.flv%3Fexpires%3D1735304880%26len%3D0%26oi%3D0%26pt%3Dandroid%26qn%3D250%26trid%3D10000960dfa5d247b4b00d01a8cc13676e98%26sigparams%3Dcdn%2Cexpires%2Clen%2Coi%2Cpt%2Cqn%2Ctrid%26cdn%3Dcn-gotcha04%26sign%3D605d0c11809a58ad1de14f36c49e9aac%26site%3D3c6d1d7e8f607bc2fb1256db5848de7f%26free_type%3D0%26mid%3D0%26sche%3Dban%26trace%3D0%26isp%3Dother%26rg%3Dother%26pv%3Dother%26sl%3D2%26deploy_env%3Dprod%26p2p_type%3D-1%26source%3Dpuv3_batch%26score%3D1%26pp%3Drtmp%26sk%3D4207df3de646838b084f14f252be3aff4b18779093814a14930f0aa3303e094e%26info_source%3Dcache%26hot_cdn%3D909449%26origin_bitrate%3D1159886%26suffix%3D2500%26vd%3Dbc%26src%3Dpuv3%26order%3D1&playurl_h265=http%3A%2F%2Fd1--cn-gotcha04.bilivideo.com%2Flive-bvc%2F468868%2Flive_413748120_72326941_minihevc.flv%3Fexpires%3D1735304880%26len%3D0%26oi%3D0%26pt%3Dandroid%26qn%3D250%26trid%3D10000960dfa5d247b4b00d01a8cc13676e98%26sigparams%3Dcdn%2Cexpires%2Clen%2Coi%2Cpt%2Cqn%2Ctrid%26cdn%3Dcn-gotcha04%26sign%3D1b2c54ade2403b62e3ebefab6551302b%26site%3D3c6d1d7e8f607bc2fb1256db5848de7f%26free_type%3D0%26mid%3D0%26sche%3Dban%26trace%3D0%26isp%3Dother%26rg%3Dother%26pv%3Dother%26sl%3D2%26p2p_type%3D-1%26sk%3D1304f646dfeb4df8b6e7ff33c167d3ad491555cae985818bb2a5ec341f048d3d%26deploy_env%3Dprod%26score%3D73%26hot_cdn%3D909449%26origin_bitrate%3D999991%26pp%3Drtmp%26source%3Dpuv3_batch%26info_source%3Dcache%26suffix%3Dminihevc%26vd%3Dbc%26src%3Dpuv3%26order%3D1&quality_description=%5B%7B%22qn%22%3A10000%2C%22desc%22%3A%22%E5%8E%9F%E7%94%BB%22%2C%22hdr_type%22%3A0%7D%2C%7B%22qn%22%3A250%2C%22desc%22%3A%22%E8%B6%85%E6%B8%85%22%2C%22hdr_type%22%3A0%7D%5D"
	} `json:"room_list"`
	RelationView []struct {
		OrderID       int    `json:"order_id"`        // 1
		ViewType      int    `json:"view_type"`       // 0
		ViewID        int    `json:"view_id"`         // 21470454
		ViewName      string `json:"view_name"`       // "主视角"
		Title         string `json:"title"`           // "虚拟整蛊综艺《尖叫浴场》EP01"
		Cover         string `json:"cover"`           // "https://i0.hdslb.com/bfs/live/new_room_cover/acf9a185dc1066f3d97c45cbdb40f1acb78d2434.jpg"
		JumpURL       string `json:"jump_url"`        // "https://live.bilibili.com/21470454?accept_quality=%5B10000%2C250%5D&av1_current_qn=0&broadcast_type=0&current_qn=250&current_quality=250&h264_current_qn=250&h265_current_qn=250&is_room_feed=1&live_play_network=other&master_url=&p2p_type=-1&playurl_av1=&playurl_h264=http%3A%2F%2Fd1--cn-gotcha04.bilivideo.com%2Flive-bvc%2F785333%2Flive_21470454_bs_2532759_2500.flv%3Fexpires%3D1735304880%26len%3D0%26oi%3D0%26pt%3Dandroid%26qn%3D250%26trid%3D10000960dfa5d247b4b00d01a8cc13676e98%26sigparams%3Dcdn%2Cexpires%2Clen%2Coi%2Cpt%2Cqn%2Ctrid%26cdn%3Dcn-gotcha04%26sign%3D605d0c11809a58ad1de14f36c49e9aac%26site%3D3c6d1d7e8f607bc2fb1256db5848de7f%26free_type%3D0%26mid%3D0%26sche%3Dban%26trace%3D0%26isp%3Dother%26rg%3Dother%26pv%3Dother%26sl%3D2%26deploy_env%3Dprod%26p2p_type%3D-1%26source%3Dpuv3_batch%26score%3D1%26pp%3Drtmp%26sk%3D4207df3de646838b084f14f252be3aff4b18779093814a14930f0aa3303e094e%26info_source%3Dcache%26hot_cdn%3D909449%26origin_bitrate%3D1159886%26suffix%3D2500%26vd%3Dbc%26src%3Dpuv3%26order%3D1&playurl_h265=http%3A%2F%2Fd1--cn-gotcha04.bilivideo.com%2Flive-bvc%2F468868%2Flive_413748120_72326941_minihevc.flv%3Fexpires%3D1735304880%26len%3D0%26oi%3D0%26pt%3Dandroid%26qn%3D250%26trid%3D10000960dfa5d247b4b00d01a8cc13676e98%26sigparams%3Dcdn%2Cexpires%2Clen%2Coi%2Cpt%2Cqn%2Ctrid%26cdn%3Dcn-gotcha04%26sign%3D1b2c54ade2403b62e3ebefab6551302b%26site%3D3c6d1d7e8f607bc2fb1256db5848de7f%26free_type%3D0%26mid%3D0%26sche%3Dban%26trace%3D0%26isp%3Dother%26rg%3Dother%26pv%3Dother%26sl%3D2%26p2p_type%3D-1%26sk%3D1304f646dfeb4df8b6e7ff33c167d3ad491555cae985818bb2a5ec341f048d3d%26deploy_env%3Dprod%26score%3D73%26hot_cdn%3D909449%26origin_bitrate%3D999991%26pp%3Drtmp%26source%3Dpuv3_batch%26info_source%3Dcache%26suffix%3Dminihevc%26vd%3Dbc%26src%3Dpuv3%26order%3D1&quality_description=%5B%7B%22qn%22%3A10000%2C%22desc%22%3A%22%E5%8E%9F%E7%94%BB%22%2C%22hdr_type%22%3A0%7D%2C%7B%22qn%22%3A250%2C%22desc%22%3A%22%E8%B6%85%E6%B8%85%22%2C%22hdr_type%22%3A0%7D%5D"
		Switch        bool   `json:"switch"`          // true
		Num           int    `json:"num"`             // 3444
		WatchIcon     string `json:"watch_icon"`      // "https://i0.hdslb.com/bfs/live/a725a9e61242ef44d764ac911691a7ce07f36c1d.png"
		LiveStatus    int    `json:"live_status"`     // 1
		TextSmall     string `json:"text_small"`      // "3444"
		UseViewVt     bool   `json:"use_view_vt"`     // false
		AnchorFace    string `json:"anchor_face"`     // "https://i0.hdslb.com/bfs/face/f12e48c0a9cfab94dc8efd9c54a24832298f1772.jpg"
		MatchLiveRoom bool   `json:"match_live_room"` // false
		MatchInfo     any    `json:"match_info"`
		Duration      int    `json:"duration"`  // 0
		UpName        string `json:"up_name"`   // ""
		PubDate       string `json:"pub_date"`  // ""
		GatherID      int    `json:"gather_id"` // 0
	} `json:"relation_view"`
	ViewPattern    int   `json:"view_pattern"` // 0
	GatherRoomList []any `json:"gather_room_list"`
}

func (LiveMultiViewNewInfo) Cmd() string {
	return CmdLiveMultiViewNewInfo
}
