package utils

import (
	"time"
)

// EqualBytes - compares two slice's of [][]byte for equality
func EqualBytes(slice1 [][]byte, slice2 [][]byte) bool {
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

// Timestamp - to identify time of occurence of an event
// returns current timestamp
// example - 2018-08-07 12:04:46.456601867 +0000 UTC m=+0.000753626
func Timestamp() string {
	return time.Now().String()
}
