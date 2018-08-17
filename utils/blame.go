package utils

import (
	"../rng"
)

// Peer - contains information of peers of a Member
type Peer struct {
	ID        int32
	PublicKey []byte
	SharedKey []byte
	Dicemix   rng.DiceMixRng
}

// Member - contains details of participant in Blame stage
type Member struct {
	ID           int32
	PrivateKey   []byte
	PublicKey    []byte
	Peers        []*Peer
	Messages     [][]byte
	MessagesHash []uint64
}
