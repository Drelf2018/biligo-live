package message

type Msg interface {
	Cmd() string
	Parse([]byte) (Msg, error)
}

var messages = make(map[string]Msg)

func Register(msg ...Msg) (n int) {
	for _, m := range msg {
		if _, exists := messages[m.Cmd()]; !exists {
			messages[m.Cmd()] = m
			n++
		}
	}
	return n
}

func MustRegister(msg ...Msg) {
	for _, m := range msg {
		messages[m.Cmd()] = m
	}
}

var UseRaw = true

func Parse(b []byte) (Msg, error) {
	raw := Raw(b)
	if v, ok := messages[raw.Cmd()]; ok {
		return v.Parse(b)
	}
	if UseRaw {
		return raw, nil
	}
	return nil, nil
}

func Delete(cmd string) {
	delete(messages, cmd)
}

func Clear() {
	for k := range messages {
		delete(messages, k)
	}
}

const (
	CmdAttention                 = "ATTENTION"                     // 用户关注
	CmdShare                     = "SHARE"                         // 用户分享直播间
	CmdSpecialAttention          = "SPECIAL_ATTENTION"             // 特别关注直播间
	CmdSysMsg                    = "SYS_MSG"                       //
	CmdPreparing                 = "PREPARING"                     // 下播
	CmdLive                      = "LIVE"                          // 开播
	CmdRoomChange                = "ROOM_CHANGE"                   // 房间信息改变
	CmdRoomRank                  = "ROOM_RANK"                     // 排名改变
	CmdRoomLimit                 = "ROOM_LIMIT"                    //
	CmdBlock                     = "BLOCK"                         //
	CmdPkPre                     = "PK_PRE"                        //
	CmdPkEnd                     = "PK_END"                        // PK判断胜负
	CmdPkSettle                  = "PK_SETTLE"                     //
	CmdSysGift                   = "SYS_GIFT"                      //
	CmdHotRankSettlement         = "HOT_RANK_SETTLEMENT"           // 荣登热门榜topX
	CmdHotRank                   = "HOT_RANK"                      // 热门榜xx榜topX
	CmdOnlineRankTop3            = "ONLINE_RANK_TOP3"              // 高能榜TOP3改变
	CmdActivityRedPacket         = "ACTIVITY_RED_PACKET"           //
	CmdPkMicEnd                  = "PK_MIC_END"                    //
	CmdStopLiveRoomList          = "STOP_LIVE_ROOM_LIST"           // 刚刚停止了直播的直播间
	CmdPlayTag                   = "PLAY_TAG"                      //
	CmdGuardMsg                  = "GUARD_MSG"                     // 舰长消息
	CmdPlayProgressBar           = "PLAY_PROGRESS_BAR"             //
	CmdHotRoomNotify             = "HOT_ROOM_NOTIFY"               //
	CmdRefresh                   = "REFRESH"                       //
	CmdRound                     = "ROUND"                         //
	CmdDanmaku                   = "DANMU_MSG"                     // 弹幕消息
	CmdWelcomeGuard              = "WELCOME_GUARD"                 //
	CmdEntryEffect               = "ENTRY_EFFECT"                  // 舰长、高能榜、老爷进入
	CmdWelcome                   = "WELCOME"                       // 欢迎进入房间(似乎已废弃)
	CmdSuperChatMessageJPN       = "SUPER_CHAT_MESSAGE_JPN"        // 醒目留言日文翻译?
	CmdSuperChatMessage          = "SUPER_CHAT_MESSAGE"            // 醒目留言
	CmdSuperChatMessageDelete    = "SUPER_CHAT_MESSAGE_DELETE"     // 删除醒目留言 (似乎有时候并不会发,同时结束时间在 CmdSuperChatMessage 可以获取)
	CmdLiveInteractiveGame       = "LIVE_INTERACTIVE_GAME"         //
	CmdSendGift                  = "SEND_GIFT"                     // 投喂礼物
	CmdRoomBlockMsg              = "ROOM_BLOCK_MSG"                // 用户被禁言
	CmdComboSend                 = "COMBO_SEND"                    // 连击礼物
	CmdAnchorLotStart            = "ANCHOR_LOT_START"              // 天选之人开始完整信息
	CmdAnchorLotEnd              = "ANCHOR_LOT_END"                // 天选之人获奖id
	CmdAnchorLotAward            = "ANCHOR_LOT_AWARD"              // 天选结果推送
	CmdVoiceJoinRoomCountInfo    = "VOICE_JOIN_ROOM_COUNT_INFO"    // 申请连麦队列变化
	CmdVoiceJoinList             = "VOICE_JOIN_LIST"               // 连麦申请、取消连麦申请
	CmdVoiceJoinStatus           = "VOICE_JOIN_STATUS"             // 开始连麦、结束连麦
	CmdCutOff                    = "CUT_OFF"                       // 被超管切断
	CmdSpecialGift               = "SPECIAL_GIFT"                  // 节奏风暴
	CmdNewGuardCount             = "NEW_GUARD_COUNT"               // 船员数量改变事件
	CmdRoomAdmins                = "ROOM_ADMINS"                   // 房管数量改变
	CmdGuardBuy                  = "GUARD_BUY"                     // 用户上舰长
	CmdUserToastMsg              = "USER_TOAST_MSG"                // 上船附带的通知
	CmdNoticeMsg                 = "NOTICE_MSG"                    // 广播消息(别的直播间投递高价礼物对所有直播间发起的广播)
	CmdAnchorLotCheckStatus      = "ANCHOR_LOT_CHECKSTATUS"        // 天选时刻前的审核
	CmdActivityBannerUpdateV2    = "ACTIVITY_BANNER_UPDATE_V2"     //
	CmdRoomRealTimeMessageUpdate = "ROOM_REAL_TIME_MESSAGE_UPDATE" // 粉丝数量改变
	CmdInteractWord              = "INTERACT_WORD"                 // 用户进入直播间
	CmdOnlineRankCount           = "ONLINE_RANK_COUNT"             // 高能榜数量更新
	CmdOnlineRankV2              = "ONLINE_RANK_V2"                // 高能榜数据
	CmdPkBattlePre               = "PK_BATTLE_PRE"                 // 大乱斗准备，10秒后开始
	CmdPkBattleSettle            = "PK_BATTLE_SETTLE"              //
	CmdHotRankChanged            = "HOT_RANK_CHANGED"              // 热门榜改变
	CmdPkBattleStart             = "PK_BATTLE_START"               // 大乱斗开始
	CmdPkBattleProcess           = "PK_BATTLE_PROCESS"             // 大乱斗双方送礼
	CmdPkEnding                  = "PK_ENDING"                     // 大乱斗尾声，最后几秒
	CmdPkBattleEnd               = "PK_BATTLE_END"                 // 大乱斗结束
	CmdPkBattleSettleUser        = "PK_BATTLE_SETTLE_USER"         //
	CmdPkBattleSettleV2          = "PK_BATTLE_SETTLE_V2"           //
	CmdPkLotteryStart            = "PK_LOTTERY_START"              // 大乱斗胜利后的抽奖
	CmdPkBestUname               = "PK_BEST_UNAME"                 // PK最佳助攻
	CmdCallOnOpposite            = "CALL_ON_OPPOSITE"              // 本直播间的观众跑去对面串门
	CmdAttentionOpposite         = "ATTENTION_OPPOSITE"            // 本直播间观众关注了对面主播
	CmdShareOpposite             = "SHARE_OPPOSITE"                // 本直播间观众分享了对面直播间
	CmdAttentionOnOpposite       = "ATTENTION_ON_OPPOSITE"         // 对面观众关注了本直播间
	CmdPkMatchInfo               = "PK_MATCH_INFO"                 // 获取对面直播间信息
	CmdPkMatchOnlineGuard        = "PK_MATCH_ONLINE_GUARD"         // 获取对面直播间舰长在线人数
	CmdPkWinningStreak           = "PK_WINNING_STREAK"             // 大乱斗连胜事件
	CmdPkDanmuMsg                = "PK_DANMU_MSG"                  // 对面的弹幕消息
	CmdPkSendGift                = "PK_SEND_GIFT"                  // 对面的礼物消息
	CmdPkInteractWord            = "PK_INTERACT_WORD"              // 对面的用户进入
	CmdPkAttention               = "PK_ATTENTION"                  // 对面新增关注
	CmdPkShare                   = "PK_SHARE"                      // 对面有人分享直播间
	CmdWatChedChange             = "WATCHED_CHANGE"                // 直播间看过人数变化
)
