package server

// import (
// 	"../nike"
// 	"../rng"
// )

// // Peer - contains information of peers of a Member
// type Peer struct {
// 	ID        int32
// 	PublicKey []byte
// 	SharedKey []byte
// 	Dicemix   rng.DiceMixRng
// }

// // Member - contains details of participant in Blame stage
// type Member struct {
// 	ID           int32
// 	PrivateKey   []byte
// 	PublicKey    []byte
// 	Peers        []*Peer
// 	Messages     [][]byte
// 	MessagesHash []uint64
// }

func startBlame(h *Hub) {
	// 	var members = make([]*Member, 0)

	// 	initBlame(h, members)

	// 	// var roots = generateroots()
}

// func initBlame(h *Hub, members []*Member) {
// 	for _, peer := range h.peers {
// 		if !peer.MessageReceived {
// 			continue
// 		}

// 		var member = &Member{}
// 		member.ID = peer.Id
// 		member.PrivateKey = peer.PrivateKey
// 		member.PublicKey = peer.PublicKey
// 		member.Messages = peer.DCSimpleVector
// 		member.Peers = make([]*Peer, 0)

// 		for _, otherPeer := range h.peers {
// 			if peer.Id == otherPeer.Id {
// 				continue
// 			}

// 			var peer = &Peer{}
// 			peer.ID = otherPeer.Id
// 			peer.PublicKey = otherPeer.PublicKey
// 			nike.DeriveSharedKeys(member)
// 			member.Peers = append(member.Peers, peer)
// 		}

// 		// recover messages
// 		members = append(members, member)
// 	}
// }

// func recoverMessages() {}
