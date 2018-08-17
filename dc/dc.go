package dc

import (
	"../messages"
)

// DC - The main interface DC_NET.
type DC interface {
	SolveDCExponential([]*messages.PeersInfo) []uint64
}
