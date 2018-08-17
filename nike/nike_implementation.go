package nike

import (
	"sync"

	"../ecdh"
	"../rng"
	"../server"
	log "github.com/sirupsen/logrus"
)

type nike struct {
	NIKE
	sync.Mutex
}

// NewNike creates a new NIKE instance
func NewNike() NIKE {
	return &nike{}
}

// DeriveSharedKeys - derives shared keys for all peers
// generates RNG based on shared key using ChaCha20
func (n *nike) DeriveSharedKeys(member *server.Member) {
	ecdh := ecdh.NewCurve25519ECDH()
	peersCount := len((*member).Peers)
	privateKey, _ := ecdh.UnmarshalSK(member.PrivateKey)

	for i := 0; i < peersCount; i++ {
		var pubkey, res = ecdh.Unmarshal((*member).Peers[i].PublicKey)
		if !res {
			log.Fatalf("Error: generating NIKE Shared Keys %v", res)
		}
		var err error
		(*member).Peers[i].SharedKey, err = ecdh.GenerateSharedSecret(privateKey, pubkey)

		if err != nil {
			log.Fatalf("Error: generating NIKE Shared Keys %v", err)
		}

		(*member).Peers[i].Dicemix = rng.NewRng((*member).Peers[i].SharedKey)
	}
}
