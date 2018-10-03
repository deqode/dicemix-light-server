package dc

import (
	"github.com/techracers-blockchain/dicemix-light-server/messages"
)

// DC - The main interface DC_NET.
type DC interface {
	SolveDCExponential([]*messages.PeersInfo) []uint64
	ResolveDCNet([]*messages.PeersInfo, int) [][]byte
}
