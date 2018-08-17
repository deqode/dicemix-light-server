package server

import (
	"log"

	"../messages"
	"../utils"
	"github.com/jinzhu/copier"
)

// handles non responsive peers
// after responseWait seconds if all peers have not submitted their response
// then remove them and consider those peers as offline
// and broadcast mesage to active peers
func registerDelayHandler(h *Hub, state int) {
	if h.nextState != state {
		log.Printf("Round has been done already %v\n", state)
		return
	}

	log.Printf("\nRound has not done %v\n", state)
	switch state {
	case messages.C_KEY_EXCHANGE:
		// if some peers have not submitted their PublicKey
		broadcastDiceMixResponse(h, messages.S_KEY_EXCHANGE, "Key Exchange Response", "")
	case messages.C_EXP_DC_VECTOR:
		// if some peers have not submitted their DC-EXP vector
		broadcastDCExponentialResponse(h, messages.S_EXP_DC_VECTOR, "Solved DC Exponential Roots", "")
	case messages.C_SIMPLE_DC_VECTOR:
		// if some peers have not submitted their DC-SIMPLE vector
		broadcastDiceMixResponse(h, messages.S_SIMPLE_DC_VECTOR, "DC Simple Response", "")
	case messages.C_TX_CONFIRMATION:
		// if some peers have not submitted their CONFIRMATION
		checkConfirmations(h)
	case messages.C_KESK_RESPONSE:
		// if some peers have not submitted their KESK
		// TODO: START-BLAME()
		startBlame(h)
	}
}

// removes offline peers from h.peers
// returns true if removed any offline peer
func filterPeers(h *Hub) bool {
	var allPeers []*messages.PeersInfo
	copier.Copy(&allPeers, &h.peers)
	h.peers = make([]*messages.PeersInfo, 0)

	for _, peer := range allPeers {
		// check if client is active and has submitted response
		if peer.MessageReceived {
			peer.MessageReceived = false
			h.peers = append(h.peers, peer)
			continue
		}

		// if client is offline and not submitted response
		removePeer(h, peer.Id)
	}
	// removed any offline peer?
	return len(allPeers) != len(h.peers)
}

// checks if all peers have submitted a valid confirmation for msgs
// if yes then DiceMix protocol is considered as successful
// else moves to BLAME stage
func checkConfirmations(h *Hub) {
	// removes offline peers
	// returns true if removed any offline peers
	res := filterPeers(h)

	// if any P_Excluded trace back to KE Stage
	if res {
		broadcastKEResponse(h)
		return
	}

	msgs := h.peers[0].Messages

	// check if any of peers does'nt agree to continue
	for _, peer := range h.peers {
		if !utils.EqualBytes(peer.Messages, msgs) ||
			len(peer.Confirmation) == 0 {
			// Blame stage - INIT KESK
			log.Printf("BLAME Stage - Peer %v does'nt provide corfirmation", peer.Id)
			broadcastKESKRequest(h)
			return
		}
	}

	// DiceMix is successful
	broadcastTXDone(h)
}

// predicts next expected RequestCodes from client againts current ResponseCode
func nextState(responseCode int) int {
	switch responseCode {
	case messages.S_START_DICEMIX:
		return messages.C_KEY_EXCHANGE
	case messages.S_KEY_EXCHANGE:
		return messages.C_EXP_DC_VECTOR
	case messages.S_EXP_DC_VECTOR:
		return messages.C_SIMPLE_DC_VECTOR
	case messages.S_SIMPLE_DC_VECTOR:
		return messages.C_TX_CONFIRMATION
	case messages.S_KESK_REQUEST:
		return messages.C_KESK_RESPONSE
	}

	return 0
}

// to keep track of number of clients which have already
// submitted the request for corresponding RequestCode (for current run)
func counter(peers []*messages.PeersInfo) (counter int) {
	for _, peer := range peers {
		if peer.MessageReceived {
			counter++
		}
	}
	return
}

// returns key by value from map
func mapkey(m map[*Client]int32, value int32) (key *Client, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}
