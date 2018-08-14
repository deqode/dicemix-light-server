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
func registerDelayHandler(h *Hub, state int) {
	if !contains(h.nextState, state) {
		log.Printf("Round has been done already %v\n", state)
		return
	}

	log.Printf("\nRound has not done %v\n", state)
	switch state {
	case commons.C_KEY_EXCHANGE:
		broadcastDiceMixResponse(h, commons.S_KEY_EXCHANGE, "Key Exchange Response", "")
	case commons.C_EXP_DC_VECTOR:
		broadcastDCExponentialResponse(h, commons.S_EXP_DC_VECTOR, "Solved DC Exponential Roots", "")
	case commons.C_SIMPLE_DC_VECTOR:
		broadcastDiceMixResponse(h, commons.S_SIMPLE_DC_VECTOR, "DC Simple Response", "")
	case commons.C_TX_CONFIRMATION:
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
		fmt.Printf("Check for PeerId - %v\n", peer.Id)
		if peer.MessageReceived {
			fmt.Printf("Recv PeerId - %v\n", peer.Id)
			peer.MessageReceived = false
			h.peers = append(h.peers, peer)
		} else {
			fmt.Printf("Not Recv PeerId - %v\n", peer.Id)
			if client, ok := mapkey(h.clients, peer.Id); ok {
				fmt.Printf("USER UN-REGISTRATION - %v\n", peer.Id)
				delete(h.clients, client)
				close(client.send)
			}
		}
	}

	return len(allPeers) != len(h.peers)
}

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
