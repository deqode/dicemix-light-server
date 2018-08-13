package dc

import (
	"../commons"
	"../field"
	"../solver"
)

type dcNet struct {
	DC
}

// NewDCNetwork creates a new DC instance
func NewDCNetwork() DC {
	return &dcNet{}
}

// obtains other peers dc[]
// and generates dc_combined[]
func (d *dcNet) SolveDCExponential(peers []*commons.PeersInfo) []uint64 {
	var i, totalMsgsCount uint32
	dcCombined := peers[0].DCVector

	// NOTE: totalMsgsCount should be less than 1000 or else FLINT would fail to obtain roots
	// and [0,0,......] will be considered as roots
	for _, peer := range peers {
		totalMsgsCount += peer.NumMsgs
	}

	for j := 1; j < len(peers); j++ {
		for i = 0; i < totalMsgsCount; i++ {
			var op1 = field.NewField(field.UInt64(dcCombined[i]))
			var op2 = field.NewField(field.UInt64(peers[j].DCVector[i]))
			dcCombined[i] = uint64(op1.Add(op2).Fp)
		}
	}

	return solver.Solve(dcCombined, int(totalMsgsCount))
}
