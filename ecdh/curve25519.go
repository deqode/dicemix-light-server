package ecdh

import (
	"crypto"
	"sync"

	ecdh "github.com/wsddn/go-ecdh"
)

type curve25519ECDH struct {
	ECDH
	sync.Mutex
}

// NewCurve25519ECDH creates a new ECDH instance that uses djb's curve25519
// elliptical curve.
func NewCurve25519ECDH() ECDH {
	return &curve25519ECDH{}
}

// Unmarshal converts byte[] to crypto.PublicKey
func (e *curve25519ECDH) Unmarshal(data []byte) (crypto.PublicKey, bool) {
	var ecdhCurve = ecdh.NewCurve25519ECDH()
	return ecdhCurve.Unmarshal(data)
}

// Unmarshal converts byte[] to crypto.PrivateKey
func (e *curve25519ECDH) UnmarshalSK(data []byte) (crypto.PrivateKey, bool) {
	var pri [32]byte
	if len(data) != 32 {
		return nil, false
	}
	copy(pri[:], data)
	return &pri, true
}

// GenerateSharedSecret creates shared key using our private key and others public key
func (e *curve25519ECDH) GenerateSharedSecret(privKey crypto.PrivateKey, pubKey crypto.PublicKey) ([]byte, error) {
	var ecdhCurve = ecdh.NewCurve25519ECDH()
	return ecdhCurve.GenerateSharedSecret(privKey, pubKey)
}
