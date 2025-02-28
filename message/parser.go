package message

import (
	"encoding/json"
)

// 过滤解析器
type FilterParser struct {
	// 需要解析的命令
	Command map[string]struct{}

	// 不在 Command 中的消息
	// 是否以 Raw 对象返回
	UseRaw bool
}

// 添加需要解析的命令
func (f *FilterParser) Add(cmd ...string) {
	if f.Command == nil {
		f.Command = make(map[string]struct{})
	}
	for _, c := range cmd {
		f.Command[c] = struct{}{}
	}
}

// 添加需要解析的消息
func (f *FilterParser) AddMsg(msg ...Msg) {
	if f.Command == nil {
		f.Command = make(map[string]struct{})
	}
	for _, m := range msg {
		f.Command[m.Cmd()] = struct{}{}
		if m, ok := m.(CmdAlias); ok {
			for _, cmd := range m.CmdAlias() {
				f.Command[cmd] = struct{}{}
			}
		}
	}
}

// 包装器
type Wrapper struct {
	CMD  string          `json:"cmd"`
	Data json.RawMessage `json:"data"`
}

func Unmarshal[T Msg](b []byte) (msg T, err error) {
	err = json.Unmarshal(b, &msg)
	return
}

func (f *FilterParser) UnmarshalMessage(b []byte) (Msg, error) {
	// 获取消息的命令
	var w Wrapper
	err := json.Unmarshal(b, &w)
	if err != nil {
		return nil, err
	}
	// 在过滤器中查找是否需要
	if f.Command != nil {
		if _, ok := f.Command[w.CMD]; !ok {
			if f.UseRaw {
				return Raw(b), nil
			}
			return nil, nil
		}
	}
	// 需要 继续判断消息类型
	switch w.CMD {

	case CmdDanmaku, CmdDanmakuWithCode:
		return Unmarshal[Danmaku](b)
	case CmdEnterRoomSuccess:
		return Unmarshal[EnterRoomSuccess](b)
	case CmdLive:
		return Unmarshal[Live](b)
	case CmdNoticeMsg:
		return Unmarshal[NoticeMsg](b)
	case CmdPreparing:
		return Unmarshal[Preparing](b)
	case CmdSysMsg:
		return Unmarshal[SysMsg](b)

	case CmdAnchorLotAward:
		return Unmarshal[AnchorLotAward](w.Data)
	case CmdAnchorLotCheckStatus:
		return Unmarshal[AnchorLotCheckStatus](w.Data)
	case CmdAnchorLotStart:
		return Unmarshal[AnchorLotStart](w.Data)
	case CmdAnchorLotEnd:
		return Unmarshal[AnchorLotEnd](w.Data)
	case CmdAreaRankChanged:
		return Unmarshal[AreaRankChanged](w.Data)

	case CmdComboSend:
		return Unmarshal[ComboSend](w.Data)

	case CmdDmInteraction:
		return Unmarshal[DmInteraction](w.Data)

	case CmdEntryEffect:
		return Unmarshal[EntryEffect](w.Data)

	case CmdGuardBuy:
		return Unmarshal[GuardBuy](w.Data)

	case CmdHotRankChanged:
		return Unmarshal[HotRankChanged](w.Data)
	case CmdHotRankSettlement:
		return Unmarshal[HotRankSettlement](w.Data)

	case CmdInteractWord:
		return Unmarshal[InteractWord](w.Data)

	case CmdLikeInfoV3Click:
		return Unmarshal[LikeInfoV3Click](w.Data)
	case CmdLikeInfoV3Update:
		return Unmarshal[LikeInfoV3Update](w.Data)
	case CmdLittleMessageBox:
		return Unmarshal[LittleMessageBox](w.Data)

	case CmdLogInNotice:
		return Unmarshal[LogInNotice](w.Data)

	case CmdMessageboxUserMedalChange:
		return Unmarshal[MessageboxUserMedalChange](w.Data)

	case CmdOnlineRankCount:
		return Unmarshal[OnlineRankCount](w.Data)
	case CmdOnlineRankTop3:
		return Unmarshal[OnlineRankTop3](w.Data)
	case CmdOnlineRankV2:
		return Unmarshal[OnlineRankV2](w.Data)

	case CmdRankChanged:
		return Unmarshal[RankChanged](w.Data)
	case CmdRevenueRankChanged:
		return Unmarshal[RevenueRankChanged](w.Data)
	case CmdRoomBlockMsg:
		return Unmarshal[RoomBlockMsg](w.Data)
	case CmdRoomChange:
		return Unmarshal[RoomChange](w.Data)
	case CmdRoomRealTimeMessageUpdate:
		return Unmarshal[RoomRealTimeMessageUpdate](w.Data)

	case CmdSendGift:
		return Unmarshal[SendGift](w.Data)
	case CmdStopLiveRoomList:
		return Unmarshal[StopLiveRoomList](w.Data)
	case CmdSuperChatMessage:
		return Unmarshal[SuperChatMessage](w.Data)
	case CmdSuperChatMessageJPN:
		return Unmarshal[SuperChatMessageJPN](w.Data)

	case CmdUserToastMsg:
		return Unmarshal[UserToastMsg](w.Data)

	case CmdVoiceJoinList:
		return Unmarshal[VoiceJoinList](w.Data)
	case CmdVoiceJoinRoomCountInfo:
		return Unmarshal[VoiceJoinRoomCountInfo](w.Data)

	case CmdWatChedChange:
		return Unmarshal[WatChedChange](w.Data)
	case CmdWidgetBanner:
		return Unmarshal[WidgetBanner](w.Data)
	default:
		if f.UseRaw {
			return Raw(b), nil
		}
		return nil, nil
	}
}

var DefaultFilterParser Parser = &FilterParser{UseRaw: true}
