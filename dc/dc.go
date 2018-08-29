package dc

import (
	"dicemix_server/messages"
)

// DC - The main interface DC_NET.
type DC interface {
	SolveDCExponential([]*messages.PeersInfo) []uint64
}
