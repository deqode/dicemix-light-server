package server

import (
	"fmt"
	"log"

	"../messages"
	"github.com/golang/protobuf/proto"
)

// handles any request message from peers
func handleRequest(message []byte, h *Hub) {
	r := &messages.GenericRequest{}
	h.Lock()
	defer h.Unlock()

	err := proto.Unmarshal(message, r)
	checkError(err)

	// check if request from client was one of
	// the expected Requests or not
	if h.nextState != int(r.Code) {
		return
	}

	// to keep track of number of clients which have already
	// submitted this request (for current run)
	var counter = counter(h.peers)

	if counter >= len(h.peers) {
		return
	}

	switch r.Code {
	case messages.C_KEY_EXCHANGE:
		request := &messages.KeyExchangeRequest{}
		err = proto.Unmarshal(message, request)
		checkError(err)

		handleKeyExchangeRequest(request, h, counter)
	case messages.C_EXP_DC_VECTOR:
		request := &messages.DCExpRequest{}
		err = proto.Unmarshal(message, request)
		checkError(err)

		handleDCExponentialRequest(request, h, counter)
	case messages.C_SIMPLE_DC_VECTOR:
		request := &messages.DCSimpleRequest{}
		err = proto.Unmarshal(message, request)
		checkError(err)

		handleDCSimpleRequest(request, h, counter)
	case messages.C_TX_CONFIRMATION:
		request := &messages.ConfirmationRequest{}
		err = proto.Unmarshal(message, request)
		checkError(err)

		handleConfirmationRequest(request, h, counter)
	case messages.C_KESK_RESPONSE:
		request := &messages.InitiaiteKESKResponse{}
		err = proto.Unmarshal(message, request)
		checkError(err)

		handleInitiateKESKResponse(request, h, counter)
	}
}

// obtains PublicKeys and NumberOfMsgs sent by peers
func handleKeyExchangeRequest(request *messages.KeyExchangeRequest, h *Hub, counter int) {
	for i := 0; i < len(h.peers); i++ {
		if h.peers[i].Id == request.Id {
			h.peers[i].PublicKey = request.PublicKey
			h.peers[i].NumMsgs = request.NumMsgs
			h.peers[i].MessageReceived = true

			fmt.Printf("Recv: handleKeyExchangeRequest PeerId - %v\n", request.Id)
			counter++
			break
		}
	}

	// if all active peers have submitted their response
	if counter == len(h.peers) {
		broadcastDiceMixResponse(h, messages.S_KEY_EXCHANGE, "Key Exchange Response", "")
	}
}

// obtains DC-EXP vector sent by peers
func handleDCExponentialRequest(request *messages.DCExpRequest, h *Hub, counter int) {
	for i := 0; i < len(h.peers); i++ {
		if h.peers[i].Id == request.Id {
			h.peers[i].DCVector = request.DCExpVector
			h.peers[i].MessageReceived = true

			fmt.Printf("Recv: handleDCExponentialRequest PeerId - %v\n", request.Id)
			counter++
			break
		}
	}

	// if all active peers have submitted their response
	if counter == len(h.peers) {
		broadcastDCExponentialResponse(h, messages.S_EXP_DC_VECTOR, "Solved DC Exponential Roots", "")
	}
}

// obtains DC-SIMPLE vector sent by peers
func handleDCSimpleRequest(request *messages.DCSimpleRequest, h *Hub, counter int) {
	for i := 0; i < len(h.peers); i++ {
		if h.peers[i].Id == request.Id {
			h.peers[i].DCSimpleVector = request.DCSimpleVector
			h.peers[i].OK = request.MyOk
			h.peers[i].MessageReceived = true
			h.peers[i].NextPublicKey = request.NextPublicKey

			fmt.Printf("Recv: handleDCSimpleRequest PeerId - %v\n", request.Id)
			counter++
			break
		}
	}

	// if all active peers have submitted their response
	if counter == len(h.peers) {
		broadcastDiceMixResponse(h, messages.S_SIMPLE_DC_VECTOR, "DC Simple Response", "")
	}
}

// obtains confirmations from peers
// if all peers provided valid confirmations then Dicemix is successful
// else moved to BLAME stage
func handleConfirmationRequest(request *messages.ConfirmationRequest, h *Hub, counter int) {
	for i := 0; i < len(h.peers); i++ {
		if h.peers[i].Id == request.Id {
			h.peers[i].Messages = request.Messages
			h.peers[i].Confirmation = request.Confirmation
			h.peers[i].MessageReceived = true

			fmt.Printf("Recv: handleConfirmationRequest PeerId - %v\n", request.Id)
			counter++
			break
		}
	}

	// if all active peers have submitted their response
	if counter == len(h.peers) {
		checkConfirmations(h)
	}
}

// obtains KESK of peers
// used in BLAME stage to identify malicious peer
func handleInitiateKESKResponse(request *messages.InitiaiteKESKResponse, h *Hub, counter int) {
	for i := 0; i < len(h.peers); i++ {
		if h.peers[i].Id == request.Id {
			h.peers[i].PrivateKey = request.PrivateKey
			h.peers[i].MessageReceived = true

			fmt.Printf("Recv: handleInitiateKESKResponse PeerId - %v\n", request.Id)
			counter++
			break
		}
	}

	// if all active peers have submitted their kesk
	if counter == len(h.peers) {
		// TODO: START-BLAME()
		startBlame(h)
	}
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
		if !equals(peer.Messages, msgs) ||
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
