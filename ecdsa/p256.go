package ecdsa

import (
	"crypto/ecdsa"
	"crypto/x509"
	"math/big"
)

type curveP256 struct {
	ECDSA
}

// NewCurveECDSA creates a new Elliptic Curve Digital Signature Algorithm  instance
func NewCurveECDSA() ECDSA {
	return &curveP256{}
}

// Verify reports whether sig is a valid signature of message by publicKey.
func (e *curveP256) Verify(publicKey, message, signature []byte) bool {
	// obtain r, s *big.Int from []byte signature
	r := new(big.Int)
	s := new(big.Int)
	r.SetBytes(signature[:len(signature)/2])
	s.SetBytes(signature[len(signature)/2:])

	// verify signature
	return pkVerify(message, publicKey, r, s)
}

func pkVerify(message []byte, publicKeyBytes []byte, r *big.Int, s *big.Int) (result bool) {
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		return false
	}

	switch publicKey := publicKey.(type) {
	case *ecdsa.PublicKey:
		return ecdsa.Verify(publicKey, message, r, s)
	default:
		return false
	}
}
