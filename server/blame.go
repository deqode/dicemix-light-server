package server

import (
	"../nike"
	"../rng"
	"../utils"
)

// Peer - contains information of peers of a Participant
type Peer struct {
	ID        int32
	SharedKey []byte
	Dicemix   rng.DiceMixRng
}

// Participant - contains details of participant in Blame stage
type Participant struct {
	ID           int32
	Peers        []*Peer
	Messages     [][]byte
	MessagesHash []uint64
}

func startBlame(h *Hub) {
	var participants = make([]*Participant, 0)
	var roots = iDcNet.SolveDCExponential(h.peers)

	// identifies honest peers (who have expected protocol messages)
	participants = initBlame(h, participants, roots)

	// TODO: detection of slot collision

	// removes malicious and offline peers
	// i.e. those peers who have sent unexpected protocol messages
	filterPeers(h)

	rotateKeys(h)
	broadcastKEResponse(h)
}

// Exclude peers who have sent unexpected protocol messages
func initBlame(h *Hub, participants []*Participant, roots []uint64) []*Participant {
	nike := nike.NewNike()

	for i := 0; i < len(h.peers); i++ {
		peer := h.peers[i]

		// TODO: check if peer has sent correct private key
		// validate(privateKey, publicKey) key pairs
		if !peer.MessageReceived {
			continue
		}

		var participant = &Participant{}
		participant.ID = peer.Id
		participant.Messages = peer.DCSimpleVector
		participant.Peers = make([]*Peer, 0)

		privateKey := peer.PrivateKey

		for _, otherPeer := range h.peers {
			if peer.Id == otherPeer.Id {
				continue
			}

			// derive sharedSecret with otherPeers
			var peer = &Peer{}
			peer.ID = otherPeer.Id
			peer.SharedKey, peer.Dicemix = nike.DeriveSharedKeys(privateKey, otherPeer.PublicKey)
			participant.Peers = append(participant.Peers, peer)
		}

		// recover messages
		participant.Messages = recoverMessages(participant.Peers, participant.Messages)

		// verify messages
		hashes, ok := verifyMessagesHash(participant.Messages, roots)
		participant.MessagesHash = hashes

		if !ok {
			// set peer.MessageReceived to false
			// so it would be removed by filterPeers()
			h.peers[i].MessageReceived = false
			continue
		}

		participants = append(participants, participant)
	}

	return participants
}

// recovers honest peers messages from his DC-SIMPLE vector
// by cancelling out randomness
func recoverMessages(peers []*Peer, messages [][]byte) [][]byte {
	messages = decodeMessages(peers, messages)
	messages = utils.RemoveEmpty(messages)
	return messages
}

// decodes messages from slots
func decodeMessages(peers []*Peer, messages [][]byte) [][]byte {
	for i := 0; i < len(peers); i++ {
		for j := 0; j < len(messages); j++ {
			// decodes messages
			// xor operation - messages[j] = dc_simple_vector[j] + <randomness for chacha20>
			utils.XorBytes(messages[j], messages[j], peers[i].Dicemix.GetBytes(20))
		}
	}

	return messages
}

// checks if message sent by peer in DC-Simple
// and Hash sent by him in DC-EXP are related or not
func verifyMessagesHash(messages [][]byte, roots []uint64) ([]uint64, bool) {
	var hashes = make([]uint64, 0)
	for _, message := range messages {
		messageHash := utils.Reduce(utils.ShortHash(utils.BytesToBase58String(message)))
		hashes = append(hashes, messageHash)
		if !utils.ContainsHash(messageHash, roots) {
			return nil, false
		}
	}
	return hashes, true
}

// rotate keys to be used in next run
func rotateKeys(h *Hub) {
	for i := 0; i < len(h.peers); i++ {
		h.peers[i].PublicKey = h.peers[i].NextPublicKey
		h.peers[i].NextPublicKey = nil
	}
}
