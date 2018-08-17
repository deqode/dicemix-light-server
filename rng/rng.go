package rng

// RNG - The main interface chacha20 DiceMixRng.
type RNG interface {
	GetBytes(dicemix DiceMixRng, bytes uint8) []byte
}
