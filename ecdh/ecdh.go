package ecdh

import (
	"crypto"
)

// ECDH - The main interface ECDH.
type ECDH interface {
	Unmarshal([]byte) (crypto.PublicKey, bool)
	UnmarshalSK([]byte) (crypto.PrivateKey, bool)
	GenerateSharedSecret(crypto.PrivateKey, crypto.PublicKey) ([]byte, error)
}
