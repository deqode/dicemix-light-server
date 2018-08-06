package dc

import (
	"../commons"
)

// DC - The main interface DC_NET.
type DC interface {
	SolveDCExponential([]*commons.PeersInfo) []uint64
}
