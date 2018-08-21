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
func registerDelayHandler(h *hub, sessionID uint64, state int) {
	if h.runs[sessionID].nextState != state {
		log.Printf("Round has been done already %v\n", state)
		return
	}

	log.Printf("\nRound has not done %v\n", state)
	switch state {
	case messages.C_KEY_EXCHANGE:
		// if some peers have not submitted their PublicKey
		broadcastDiceMixResponse(h, sessionID, messages.S_KEY_EXCHANGE, "Key Exchange Response", "")
	case messages.C_EXP_DC_VECTOR:
		// if some peers have not submitted their DC-EXP vector
		broadcastDCExponentialResponse(h, sessionID, messages.S_EXP_DC_VECTOR, "Solved DC Exponential Roots", "")
	case messages.C_SIMPLE_DC_VECTOR:
		// if some peers have not submitted their DC-SIMPLE vector
		broadcastDiceMixResponse(h, sessionID, messages.S_SIMPLE_DC_VECTOR, "DC Simple Response", "")
	case messages.C_TX_CONFIRMATION:
		// if some peers have not submitted their CONFIRMATION
		checkConfirmations(h, sessionID)
	case messages.C_KESK_RESPONSE:
		// if some peers have not submitted their KESK
		// TODO: START-BLAME()
		startBlame(h, sessionID)
	}
}

// removes offline peers from h.runs[sessionID].peers
// returns true if removed any offline peer
func filterPeers(h *hub, sessionID uint64) bool {
	var allPeers []*messages.PeersInfo
	copier.Copy(&allPeers, &h.runs[sessionID].peers)
	h.runs[sessionID].peers = make([]*messages.PeersInfo, 0)

	for _, peer := range allPeers {
		// check if client is active and has submitted response
		if peer.MessageReceived {
			peer.MessageReceived = false
			h.runs[sessionID].peers = append(h.runs[sessionID].peers, peer)
			continue
		}

		// if client is offline and not submitted response
		removePeer(h, peer.Id)
	}
	// removed any offline peer?
	return len(allPeers) != len(h.runs[sessionID].peers)
}

// checks if all peers have submitted a valid confirmation for msgs
// if yes then DiceMix protocol is considered as successful
// else moves to BLAME stage
func checkConfirmations(h *hub, sessionID uint64) {
	// removes offline peers
	// returns true if removed any offline peers
	res := filterPeers(h, sessionID)

	// if any P_Excluded trace back to KE Stage
	if res {
		broadcastKEResponse(h, sessionID)
		return
	}

	msgs := h.runs[sessionID].peers[0].Messages

	// check if any of peers does'nt agree to continue
	for _, peer := range h.runs[sessionID].peers {
		if !utils.EqualBytes(peer.Messages, msgs) ||
			len(peer.Confirmation) == 0 {
			// Blame stage - INIT KESK
			log.Printf("BLAME Stage - Peer %v does'nt provide corfirmation", peer.Id)
			broadcastKESKRequest(h, sessionID)
			return
		}
	}

	// DiceMix is successful
	broadcastTXDone(h, sessionID)
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

// returns client connection object from client id
func getClient(m map[*client]int32, value int32) (key *client, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}
