package eddsa

import (
	"golang.org/x/crypto/ed25519"
)

type curveED25519 struct {
	EdDSA
}

// NewCurveED25519 creates a new Edwards-Curve Digital Signature Algorithm (EdDSA) instance
func NewCurveED25519() EdDSA {
	return &curveED25519{}
}

// Verify reports whether sig is a valid signature of message by publicKey. It
// will panic if len(publicKey) is not PublicKeySize = 32	.
func (e *curveED25519) Verify(publicKey, message, signature []byte) bool {
	return ed25519.Verify(publicKey, message, signature)
}
