package ecdh

import (
	"crypto"
)

// ECDH - The main interface ECDH.
type ECDH interface {
	PublicKey([]byte) ([]byte, bool)
	Unmarshal([]byte) (crypto.PublicKey, bool)
	UnmarshalSK([]byte) (crypto.PrivateKey, bool)
	GenerateSharedSecret(crypto.PrivateKey, crypto.PublicKey) ([]byte, error)
}
