package server

import (
	"math/rand"
	"time"

	"../commons"
)

// associates 20 letter string to each response from server side
// client replying (to response from server side) need to add last obtained string in request
// used for error handling cases like
// handling delayed responses from client side
// to protect from malicious peers trying to send messages in back and forth rounds
// func initRoundUUID(h *Hub) {
// 	h.roundUUID[commons.S_JOIN_RESPONSE] = randUUIDString()
// 	h.roundUUID[commons.S_START_DICEMIX] = randUUIDString()
// 	h.roundUUID[commons.S_KEY_EXCHANGE] = randUUIDString()
// 	h.roundUUID[commons.S_EXP_DC_VECTOR] = randUUIDString()
// 	h.roundUUID[commons.S_SIMPLE_DC_VECTOR] = randUUIDString()
// 	h.roundUUID[commons.S_TX_CONFIRMATION] = randUUIDString()
// }

func contains(expected []int, obtained int) bool {
	for _, value := range expected {
		if value == obtained {
			return true
		}
	}
	return false
}

func counter(peers []*commons.PeersInfo) (counter int) {
	for _, peer := range peers {
		if peer.MessageReceived {
			counter++
		}
	}
	return
}

// generates a 20 letter random string
func randUUIDString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 20)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// to identify time of occurence of an event
// returns current timestamp
// example - 2018-08-07 12:04:46.456601867 +0000 UTC m=+0.000753626
func timestamp() string {
	return time.Now().String()
}

// returns key by value from map
func mapkey(m map[*Client]int32, value int32) (key *Client, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}
