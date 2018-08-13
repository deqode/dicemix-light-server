package server

import (
	"fmt"

	"../commons"
	"github.com/jinzhu/copier"
)

func registerDelayHandler(h *Hub, code uint32, message string) {
	if h.roundUUID[code] != h.lastRoundUUID {
		fmt.Printf("\nRound has been done already %v, %v\n", message, code)
		return
	}
	fmt.Printf("\nRound has not done %v, %v\n", message, code)

	switch code {
	case commons.S_START_DICEMIX, commons.C_KEY_EXCHANGE, commons.C_SIMPLE_DC_VECTOR, commons.C_TX_CONFIRMATION:
		broadcastDiceMixResponse(h, code, message, "")
	case commons.C_EXP_DC_VECTOR:
		broadcastDCExponentialResponse(h, code, message, "")
	}
}

func filterPeers(h *Hub) {
	var allPeers []*commons.PeersInfo
	copier.Copy(&allPeers, &h.peers)
	h.peers = make([]*commons.PeersInfo, 0)

	for _, peer := range allPeers {
		if peer.MessageReceived {
			peer.MessageReceived = false
			h.peers = append(h.peers, peer)
		}
	}
}
