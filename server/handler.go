package server

import (
	"fmt"
	"os"

	"../commons"
	"github.com/golang/protobuf/proto"
)

func handleRequest(message []byte, h *Hub) {
	r := &commons.GenericRequest{}
	h.Lock()
	defer h.Unlock()

	if err := proto.Unmarshal(message, r); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	switch r.Code {
	case commons.C_KEY_EXCHANGE:
		request := &commons.KeyExchangeRequest{}
		if err := proto.Unmarshal(message, request); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		handleKeyExchangeRequest(request, h)
	case commons.C_EXP_DC_VECTOR:
		request := &commons.DCExpRequest{}
		if err := proto.Unmarshal(message, request); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		handleDCExponentialRequest(request, h)
	case commons.C_SIMPLE_DC_VECTOR:
		request := &commons.DCSimpleRequest{}
		if err := proto.Unmarshal(message, request); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		handleDCSimpleRequest(request, h)
	case commons.C_TX_CONFIRMATION:
		request := &commons.ConfirmationRequest{}
		if err := proto.Unmarshal(message, request); err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		handleConfirmationRequest(request, h)

	}
}

func handleKeyExchangeRequest(request *commons.KeyExchangeRequest, h *Hub) {
	if h.lastRoundUUID == request.LastMessageUUID {
		var counter int
		for _, peer := range h.peers {
			if len(peer.PublicKey) != 0 {
				counter++
			}
		}

		if counter < maxPeers {
			for i := 0; i < len(h.peers); i++ {
				if h.peers[i].Id == request.Id {
					h.peers[i].PublicKey = request.PublicKey
					h.peers[i].NumMsgs = request.NumMsgs
					counter++
				}
			}

			if counter == maxPeers {
				broadcastDiceMixResponse(h, commons.S_KEY_EXCHANGE, "Key Exchange Response", "")
			}
		}
	}
}

func handleDCExponentialRequest(request *commons.DCExpRequest, h *Hub) {
	if h.lastRoundUUID == request.LastMessageUUID {
		var counter int
		for _, peer := range h.peers {
			if len(peer.DCVector) != 0 {
				counter++
			}
		}

		if counter < maxPeers {
			for i := 0; i < len(h.peers); i++ {
				if h.peers[i].Id == request.Id {
					h.peers[i].DCVector = request.DCExpVector
					counter++
				}
			}

			if counter == maxPeers {
				broadcastDCExponentialResponse(h, commons.S_EXP_DC_VECTOR, "Solved DC Exponential Roots", "")
			}
		}
	}
}

func handleDCSimpleRequest(request *commons.DCSimpleRequest, h *Hub) {
	if h.lastRoundUUID == request.LastMessageUUID {
		var counter int
		for _, peer := range h.peers {
			if len(peer.DCSimpleVector) != 0 {
				counter++
			}
		}

		if counter < maxPeers {
			for i := 0; i < len(h.peers); i++ {
				if h.peers[i].Id == request.Id {
					h.peers[i].DCSimpleVector = request.DCSimpleVector
					h.peers[i].OK = request.MyOk
					counter++
				}
			}

			if counter == maxPeers {
				broadcastDiceMixResponse(h, commons.S_SIMPLE_DC_VECTOR, "DC Simple Response", "")
			}
		}
	}
}

func handleConfirmationRequest(request *commons.ConfirmationRequest, h *Hub) {
	if h.lastRoundUUID == request.LastMessageUUID {
		var counter int
		for _, peer := range h.peers {
			if len(peer.Messages) != 0 {
				counter++
			}
		}

		if counter < maxPeers {
			for i := 0; i < len(h.peers); i++ {
				if h.peers[i].Id == request.Id {
					h.peers[i].Messages = request.Messages
					h.peers[i].Confirm = request.Confirm
					counter++
				}
			}

			if counter == maxPeers {
				broadcastDiceMixResponse(h, commons.S_TX_CONFIRMATION, "Confirmation Response", "")
			}
		}
	}
}
