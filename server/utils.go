package server

import (
	"time"

	"../commons"
)

// returns true if (obtained ∈ expected)
func contains(expected []int, obtained int) bool {
	for _, value := range expected {
		if value == obtained {
			return true
		}
	}
	return false
}

// compares two slice's of [][]byte for equality
func equals(slice1 [][]byte, slice2 [][]byte) bool {
	if (slice1 == nil) != (slice2 == nil) {
		return false
	}

	if len(slice1) != len(slice2) {
		return false
	}

	for i := 0; i < len(slice1); i++ {
		if len(slice1[i]) != len(slice2[i]) {
			return false
		}

		for j := 0; j < len(slice1[i]); j++ {
			if slice1[i][j] != slice2[i][j] {
				return false
			}
		}
	}
	return true
}

// to keep track of number of clients which have already
// submitted the request for corresponding RequestCode (for current run)
func counter(peers []*commons.PeersInfo) (counter int) {
	for _, peer := range peers {
		if peer.MessageReceived {
			counter++
		}
	}
	return
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
