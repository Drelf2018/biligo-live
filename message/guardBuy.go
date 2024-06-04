package message

/*
	{
		"guard_level": 2,
		"price": 1998000,
		"uid": 1751924,
		"num": 1,
		"gift_id": 10002,
		"gift_name": "提督",
		"start_time": 1635766940,
		"end_time": 1635766940,
		"username": "ppmann"
	}
*/
type GuardBuy struct {
	GuardLevel int    `json:"guard_level"`
	Price      int    `json:"price"`
	UID        int64  `json:"uid"`
	Num        int    `json:"num"`
	GiftID     int64  `json:"gift_id"`
	GiftName   string `json:"gift_name"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Username   string `json:"username"`
}

func (GuardBuy) Cmd() string {
	return CmdGuardBuy
}

var _ = Register[GuardBuy]()
