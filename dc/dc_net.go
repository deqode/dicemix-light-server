package dc

import (
	"../field"
	"../messages"
	"../solver"
)

type dcNet struct {
	DC
}

// NewDCNetwork creates a new DC instance
func NewDCNetwork() DC {
	return &dcNet{}
}

// obtains all peers DC-EXP vectors
// combines them and generates DC-COMBINED vector
// solves DC-COMBINED vector and obtain's its roots using Flint
func (d *dcNet) SolveDCExponential(peers []*messages.PeersInfo) []uint64 {
	var i, totalMsgsCount uint32
	var dcCombined = make([]uint64, len(peers[0].DCVector))
	copy(dcCombined, peers[0].DCVector)

	// obtain Total Messages Count
	for _, peer := range peers {
		totalMsgsCount += peer.NumMsgs
	}

	// generates DC-COMBINED vector
	for j := 1; j < len(peers); j++ {
		for i = 0; i < totalMsgsCount; i++ {
			var op1 = field.NewField(field.UInt64(dcCombined[i]))
			var op2 = field.NewField(field.UInt64(peers[j].DCVector[i]))
			dcCombined[i] = uint64(op1.Add(op2).Fp)
		}
	}

	// NOTE: totalMsgsCount should be less than 1000 or else FLINT would fail to obtain roots
	// and [0,0,......] will be considered as roots
	// Basic sanity check to avoid weird inputs
	// check - solver/solver_flint.cpp (46)
	return solver.Solve(dcCombined, int(totalMsgsCount))
}
