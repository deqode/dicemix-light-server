package server

import (
	"log"

	"../nike"
	"../rng"
	"../utils"

	op "github.com/adam-hanna/arrayOperations"
)

// Peer - contains information of peers of a Participant
type peerInfo struct {
	ID        int32
	SharedKey []byte
	Dicemix   rng.DiceMixRng
}

// Participant - contains details of participant in Blame stage
type participant struct {
	ID           int32
	Peers        []*peerInfo
	Messages     [][]byte
	MessagesHash []uint64
}

func startBlame(h *hub, sessionID uint64) {
	h.Lock()
	defer h.Unlock()

	var participants = make([]*participant, 0)
	var roots = iDcNet.SolveDCExponential(h.runs[sessionID].peers)

	// identifies honest peers (who have expected protocol messages)
	participants = initBlame(h, sessionID, participants, roots)

	// identify and exclude peers involved in slot collision
	slotCollision(h, sessionID, participants)

	// removes malicious and offline peers
	// i.e. those peers who have sent unexpected protocol messages
	filterPeers(h, sessionID)

	rotateKeys(h, sessionID)
	broadcastKEResponse(h, sessionID)
}

// Exclude peers who have sent unexpected protocol messages
func initBlame(h *hub, sessionID uint64, participants []*participant, roots []uint64) []*participant {
	nike := nike.NewNike()

	for i := 0; i < len(h.runs[sessionID].peers); i++ {
		peer := h.runs[sessionID].peers[i]

		// TODO: check if peer has sent correct private key
		// validate(privateKey, publicKey) key pairs

		// do not perform following actions for
		// those peers whcih have not sent their KESK
		if !peer.MessageReceived {
			continue
		}

		// create participant object to store info of a valid
		// participant of blame (i.e. whcih has sent his KESK)
		var participant = &participant{}
		participant.ID = peer.Id
		participant.Messages = peer.DCSimpleVector
		participant.Peers = make([]*peerInfo, 0)

		privateKey := peer.PrivateKey

		// for every peer active till confirmation
		// irrespective of he has sent his kesk or not
		for _, otherPeer := range h.runs[sessionID].peers {
			if peer.Id == otherPeer.Id {
				continue
			}

			// derive sharedSecret with otherPeers
			var peer = &peerInfo{}
			peer.ID = otherPeer.Id
			peer.SharedKey, peer.Dicemix = nike.DeriveSharedKeys(privateKey, otherPeer.PublicKey)

			// append peer to peers of participant
			participant.Peers = append(participant.Peers, peer)
		}

		// recover messages - obtains messages of participant from his DC-Simple broadcast
		participant.Messages = recoverMessages(participant.Peers, participant.Messages)

		// verify messages checks if peer has sent
		// unexpected message and corresponding hash
		hashes, ok := verifyMessagesHash(participant.Messages, roots)
		participant.MessagesHash = hashes

		if !ok {
			// set peer.MessageReceived to false
			// so it would be removed by filterPeers()
			h.runs[sessionID].peers[i].MessageReceived = false
			continue
		}

		// if participant is valid add to valid participants
		participants = append(participants, participant)
	}

	return participants
}

// to identify peers who are involved in slot collision
// Exclude peers who are involved in a slot collision,
// i.e., a message hash collision
func slotCollision(h *hub, sessionID uint64, participants []*participant) {
	// store id's of peers involved in slot collision
	var collisions = make([]int32, 0)

	// for all (p1, p2) in P^2
	for i := 0; i < len(participants); i++ {
		p1 := participants[i]
		for j := i + 1; j < len(participants); j++ {
			p2 := participants[j]

			intersection, ok := op.Intersect(p1.MessagesHash, p2.MessagesHash)
			slice, ok := intersection.Interface().([]uint64)

			if !ok {
				log.Fatalf("Error: Cannot convert reflect.Value to []uint64")
			}

			// if |intersection| == 0 {no collision occured between peer1 and peer2}
			if len(slice) == 0 {
				continue
			}

			// collsion occured between peer1 and peer2
			// add them to collisions slice
			// P_exclude := P_exclude U {p1, p2}
			collisions = append(collisions, p1.ID)
			collisions = append(collisions, p2.ID)
		}
	}

	// remove every peer involved in slot collision
	for i := 0; i < len(collisions); i++ {
		for j := 0; j < len(h.runs[sessionID].peers); j++ {
			// if peer is not involved in collision
			if collisions[i] != h.runs[sessionID].peers[j].Id {
				continue
			}

			// set peer.MessageReceived to false
			// so it would be removed by filterPeers()
			h.runs[sessionID].peers[j].MessageReceived = false
		}
	}
}

// recovers honest peers messages from his DC-SIMPLE vector
// by cancelling out randomness
func recoverMessages(peers []*peerInfo, messages [][]byte) [][]byte {
	messages = decodeMessages(peers, messages)
	messages = utils.RemoveEmpty(messages)
	return messages
}

// decodes messages from slots
func decodeMessages(peers []*peerInfo, messages [][]byte) [][]byte {
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
// returns hashes of messages from roots (if valid)
// bool - roots contains valid hashes of messages or not
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
// (kepk) := (my_next_kepk)
// (my_next_kepk) := (undef)
func rotateKeys(h *hub, sessionID uint64) {
	for i := 0; i < len(h.runs[sessionID].peers); i++ {
		h.runs[sessionID].peers[i].PublicKey = h.runs[sessionID].peers[i].NextPublicKey
		h.runs[sessionID].peers[i].NextPublicKey = nil
	}
}
