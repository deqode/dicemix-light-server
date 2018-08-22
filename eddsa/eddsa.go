package eddsa

// EdDSA - The main interface ed25519.
type EdDSA interface {
	Verify([]byte, []byte, []byte) bool
}
