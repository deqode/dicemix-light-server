package utils

import (
	"bytes"
	"time"

	"../field"
	base58 "github.com/jbenet/go-base58"
	"github.com/shomali11/util/xhashes"
)

// ContainsHash - returns true if (hash ∈ roots)
func ContainsHash(hash uint64, roots []uint64) bool {
	for _, root := range roots {
		if hash == root {
			return true
		}
	}
	return false
}

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

// RemoveEmpty - removes empty byte slices from messages
func RemoveEmpty(messages [][]byte) [][]byte {
	emptyByte := make([]byte, 20)
	output := make([][]byte, 0)

	for _, message := range messages {
		if !bytes.Equal(message, emptyByte) {
			output = append(output, message)
		}
	}

	return output
}

// BytesToBase58String - converts []byte to Base58 Encoded string
func BytesToBase58String(bytes []byte) string {
	return base58.Encode(bytes)
}

// ShortHash - returns 64bit hash of input string message
func ShortHash(message string) uint64 {
	// NOTE: after DC-EXP roots would contain hash reduced into field
	// (as final result would be in field)
	return xhashes.FNV64(message)
}

// Reduce - reduces value into field range
func Reduce(value uint64) uint64 {
	return uint64(field.NewField(field.UInt64(value)).Fp)
}

// Timestamp - to identify time of occurence of an event
// returns current timestamp
// example - 2018-08-07 12:04:46.456601867 +0000 UTC m=+0.000753626
func Timestamp() string {
	return time.Now().String()
}
