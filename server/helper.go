package server

import (
	"fmt"
	"log"

	"../commons"
	"github.com/jinzhu/copier"
)

// handles non responsive peers
// after responseWait seconds if all peers have not submitted their response
// then remove them and consider those peers as offline
// and broadcast mesage to active peers
func registerDelayHandler(h *Hub, state int) {
	if !contains(h.nextState, state) {
		log.Printf("Round has been done already %v\n", state)
		return
	}

	log.Printf("\nRound has not done %v\n", state)
	switch state {
	case commons.C_KEY_EXCHANGE:
		// if some peers have not submitted their PublicKey
		broadcastDiceMixResponse(h, commons.S_KEY_EXCHANGE, "Key Exchange Response", "")
	case commons.C_EXP_DC_VECTOR:
		// if some peers have not submitted their DC-EXP vector
		broadcastDCExponentialResponse(h, commons.S_EXP_DC_VECTOR, "Solved DC Exponential Roots", "")
	case commons.C_SIMPLE_DC_VECTOR:
		// if some peers have not submitted their DC-SIMPLE vector
		broadcastDiceMixResponse(h, commons.S_SIMPLE_DC_VECTOR, "DC Simple Response", "")
	case commons.C_TX_CONFIRMATION:
		// if some peers have not submitted their CONFIRMATION
		broadcastDiceMixResponse(h, commons.S_TX_CONFIRMATION, "Confirmation Response", "")

	}
}

// removes offline peers from h.peers
// returns true if removed any offline peer
func filterPeers(h *Hub) bool {
	var allPeers []*commons.PeersInfo
	copier.Copy(&allPeers, &h.peers)
	h.peers = make([]*commons.PeersInfo, 0)

	for _, peer := range allPeers {
		if peer.MessageReceived {
			peer.MessageReceived = false
			h.peers = append(h.peers, peer)
		} else {
			if client, ok := mapkey(h.clients, peer.Id); ok {
				// remove offline peers from clients
				fmt.Printf("USER UN-REGISTRATION - %v\n", peer.Id)
				delete(h.clients, client)
				close(client.send)
			}
		}
	}
	// removed any offline peer?
	return len(allPeers) != len(h.peers)
}

// predicts next expected RequestCodes from client againts current ResponseCode
func nextState(responseCode int) (nextState []int) {
	nextState = make([]int, 0)

	switch responseCode {
	case commons.S_START_DICEMIX:
		nextState = append(nextState, commons.C_KEY_EXCHANGE)
	case commons.S_KEY_EXCHANGE:
		nextState = append(nextState, commons.C_EXP_DC_VECTOR)
	case commons.S_EXP_DC_VECTOR:
		nextState = append(nextState, commons.C_SIMPLE_DC_VECTOR)
	case commons.S_SIMPLE_DC_VECTOR:
		nextState = append(nextState, commons.C_TX_CONFIRMATION)
		nextState = append(nextState, commons.C_BLAME)
		// case commons.S_TX_CONFIRMATION:
		// 		nextState.append(commons.C_KEY_EXCHANGE)
	}

	return
}
