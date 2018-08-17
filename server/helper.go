package server

import (
	"fmt"
	"log"

	"../messages"
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
		if client, ok := mapkey(h.clients, peer.Id); ok {
			// remove offline peers from clients
			fmt.Printf("USER UN-REGISTRATION - %v\n", peer.Id)
			delete(h.clients, client)
			close(client.send)
		}
	}
	// removed any offline peer?
	return len(allPeers) != len(h.peers)
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
