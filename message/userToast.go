package message

/*
	{
		"guard_level": 2,
		"op_type": 1,
		"payflow_id": "2111011646287352179657661",
		"unit": "月",
		"is_show": 0,
		"num": 1,
		"price": 1998000,
		"start_time": 1635756388,
		"svga_block": 0,
		"user_show": true,
		"color": "#E17AFF",
		"end_time": 1635756388,
		"role_name": "提督",
		"toast_msg": "<%何图灵%> 开通了提督",
		"uid": 4497965,
		"anchor_show": true,
		"dmscore": 96,
		"target_guard_count": 12089,
		"username": "何图灵"
	}
*/
type UserToastMsg struct {
	GuardLevel       int    `json:"guard_level"`
	OpType           int    `json:"op_type"`
	PayflowID        string `json:"payflow_id"`
	Unit             string `json:"unit"`
	IsShow           int    `json:"is_show"`
	Num              int    `json:"num"`
	Price            int64  `json:"price"`
	StartTime        int64  `json:"start_time"`
	SvgaBlock        int    `json:"svga_block"`
	UserShow         bool   `json:"user_show"`
	Color            string `json:"color"`
	EndTime          int64  `json:"end_time"`
	RoleName         string `json:"role_name"`
	ToastMsg         string `json:"toast_msg"`
	UID              int64  `json:"uid"`
	AnchorShow       bool   `json:"anchor_show"`
	DmScore          int    `json:"dmscore"`
	TargetGuardCount int    `json:"target_guard_count"`
	Username         string `json:"username"`
}

func (UserToastMsg) Cmd() string {
	return CmdUserToastMsg
}

var _ = Register[UserToastMsg]()
