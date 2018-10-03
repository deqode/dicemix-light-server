package nike

import "github.com/techracers-blockchain/dicemix-light-server/rng"

// NIKE - The main interface for Non-interactive Key Exchange (NIKE).
type NIKE interface {
	DeriveSharedKeys([]byte, []byte) ([]byte, rng.DiceMixRng)
}
