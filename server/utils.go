package server

import (
	"math/rand"
	"time"

	"../commons"
)

func timestamp() string {
	return time.Now().String()
}

func initRoundUUID(h *Hub) {
	h.roundUUID[commons.S_JOIN_RESPONSE] = randUUIDString()
	h.roundUUID[commons.S_START_DICEMIX] = randUUIDString()
	h.roundUUID[commons.S_KEY_EXCHANGE] = randUUIDString()
	h.roundUUID[commons.S_EXP_DC_VECTOR] = randUUIDString()
	h.roundUUID[commons.S_SIMPLE_DC_VECTOR] = randUUIDString()
	h.roundUUID[commons.S_TX_CONFIRMATION] = randUUIDString()
}

func randUUIDString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 20)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
