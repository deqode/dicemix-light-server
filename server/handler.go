package server

import (
	"fmt"

	"../messages"
	"github.com/golang/protobuf/proto"
)

// handles any request message from peers
func handleRequest(message []byte, h *hub) {
	h.Lock()
	defer h.Unlock()

	r := &messages.GenericRequest{}
	err := proto.Unmarshal(message, r)
	checkError(err)

	runInfo := h.runs[r.SessionId]

	// check if request from client was one of
	// the expected Requests or not
	if runInfo.nextState != int(r.Code) {
		return
	}

	// to keep track of number of clients which have already
	// submitted this request (for current run)
	var counter = counter(runInfo.peers)

	if counter >= len(runInfo.peers) {
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
func handleKeyExchangeRequest(request *messages.KeyExchangeRequest, h *hub, counter int) {
	for i := 0; i < len(h.runs[request.SessionId].peers); i++ {
		if h.runs[request.SessionId].peers[i].Id == request.Id {
			h.runs[request.SessionId].peers[i].PublicKey = request.PublicKey
			h.runs[request.SessionId].peers[i].NumMsgs = request.NumMsgs
			h.runs[request.SessionId].peers[i].MessageReceived = true

			fmt.Printf("Recv: handleKeyExchangeRequest PeerId - %v\n", request.Id)
			counter++
			break
		}
	}

	// if all active peers have submitted their response
	if counter == len(h.runs[request.SessionId].peers) {
		broadcastDiceMixResponse(h, request.SessionId, messages.S_KEY_EXCHANGE, "Key Exchange Response", "")
	}
}

// obtains DC-EXP vector sent by peers
func handleDCExponentialRequest(request *messages.DCExpRequest, h *hub, counter int) {
	for i := 0; i < len(h.runs[request.SessionId].peers); i++ {
		if h.runs[request.SessionId].peers[i].Id == request.Id {
			h.runs[request.SessionId].peers[i].DCVector = request.DCExpVector
			h.runs[request.SessionId].peers[i].MessageReceived = true

			fmt.Printf("Recv: handleDCExponentialRequest PeerId - %v\n", request.Id)
			counter++
			break
		}
	}

	// if all active peers have submitted their response
	if counter == len(h.runs[request.SessionId].peers) {
		broadcastDCExponentialResponse(h, request.SessionId, messages.S_EXP_DC_VECTOR, "Solved DC Exponential Roots", "")
	}
}

// obtains DC-SIMPLE vector sent by peers
func handleDCSimpleRequest(request *messages.DCSimpleRequest, h *hub, counter int) {
	for i := 0; i < len(h.runs[request.SessionId].peers); i++ {
		if h.runs[request.SessionId].peers[i].Id == request.Id {
			h.runs[request.SessionId].peers[i].DCSimpleVector = request.DCSimpleVector
			h.runs[request.SessionId].peers[i].OK = request.MyOk
			h.runs[request.SessionId].peers[i].MessageReceived = true
			h.runs[request.SessionId].peers[i].NextPublicKey = request.NextPublicKey

			fmt.Printf("Recv: handleDCSimpleRequest PeerId - %v\n", request.Id)
			counter++
			break
		}
	}

	// if all active peers have submitted their response
	if counter == len(h.runs[request.SessionId].peers) {
		broadcastDiceMixResponse(h, request.SessionId, messages.S_SIMPLE_DC_VECTOR, "DC Simple Response", "")
	}
}

// obtains confirmations from peers
// if all peers provided valid confirmations then Dicemix is successful
// else moved to BLAME stage
func handleConfirmationRequest(request *messages.ConfirmationRequest, h *hub, counter int) {
	for i := 0; i < len(h.runs[request.SessionId].peers); i++ {
		if h.runs[request.SessionId].peers[i].Id == request.Id {
			h.runs[request.SessionId].peers[i].Messages = request.Messages
			h.runs[request.SessionId].peers[i].Confirmation = request.Confirmation
			h.runs[request.SessionId].peers[i].MessageReceived = true

			fmt.Printf("Recv: handleConfirmationRequest PeerId - %v\n", request.Id)
			counter++
			break
		}
	}

	// if all active peers have submitted their response
	if counter == len(h.runs[request.SessionId].peers) {
		checkConfirmations(h, request.SessionId)
	}
}

// obtains KESK of peers
// used in BLAME stage to identify malicious peer
func handleInitiateKESKResponse(request *messages.InitiaiteKESKResponse, h *hub, counter int) {
	for i := 0; i < len(h.runs[request.SessionId].peers); i++ {
		if h.runs[request.SessionId].peers[i].Id == request.Id {
			h.runs[request.SessionId].peers[i].PrivateKey = request.PrivateKey
			h.runs[request.SessionId].peers[i].MessageReceived = true

			fmt.Printf("Recv: handleInitiateKESKResponse PeerId - %v\n", request.Id)
			counter++
			break
		}
	}

	// if all active peers have submitted their kesk
	if counter == len(h.runs[request.SessionId].peers) {
		// initiate blame
		startBlame(h, request.SessionId)
	}
}
