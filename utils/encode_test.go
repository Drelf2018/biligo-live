package utils_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/iyear/biligo-live/utils"
)

func TestEncode(t *testing.T) {
	authParams := utils.VerifyData{
		Platform: "web",
		Protover: 1,
		RoomID:   21852,
		UID:      int(rand.Float64()*200000000000000.0 + 100000000000000.0),
		Type:     2,
		Key:      "",
	}

	body, err := json.Marshal(authParams)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(utils.Encode(utils.WsVerInt, utils.WsOpEnterRoom, body))
}
