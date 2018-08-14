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

	// check if request from client was one of
	// the expected Requests or not
	if !contains(h.nextState, int(r.Code)) {
		return
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
		fmt.Printf("RECV: CODE - %v, NEXT-STATE - %v\n", r.Code, h.nextState)
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
	// to keep track of number of clients which have already
	// submitted this request (for current run)
	var counter = counter(h.peers)

	if counter < len(h.peers) {
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

		if counter == len(h.peers) {
			broadcastDiceMixResponse(h, commons.S_KEY_EXCHANGE, "Key Exchange Response", "")
		}
	}
}

func handleDCExponentialRequest(request *commons.DCExpRequest, h *Hub) {
	// to keep track of number of clients which have already
	// submitted this request (for current run)
	var counter = counter(h.peers)

	if counter < len(h.peers) {
		for i := 0; i < len(h.peers); i++ {
			if h.peers[i].Id == request.Id {
				h.peers[i].DCVector = request.DCExpVector
				h.peers[i].MessageReceived = true

				fmt.Printf("Recv: handleDCExponentialRequest PeerId - %v\n", request.Id)
				counter++
				break
			}
		}

		if counter == len(h.peers) {
			broadcastDCExponentialResponse(h, commons.S_EXP_DC_VECTOR, "Solved DC Exponential Roots", "")
		}
	}

}

func handleDCSimpleRequest(request *commons.DCSimpleRequest, h *Hub) {
	// to keep track of number of clients which have already
	// submitted this request (for current run)
	var counter = counter(h.peers)

	if counter < len(h.peers) {
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

		if counter == len(h.peers) {
			broadcastDiceMixResponse(h, commons.S_SIMPLE_DC_VECTOR, "DC Simple Response", "")
		}
	}
}

func handleConfirmationRequest(request *commons.ConfirmationRequest, h *Hub) {
	// to keep track of number of clients which have already
	// submitted this request (for current run)
	var counter = counter(h.peers)

	if counter < len(h.peers) {
		for i := 0; i < len(h.peers); i++ {
			if h.peers[i].Id == request.Id {
				h.peers[i].Messages = request.Messages
				h.peers[i].Confirm = request.Confirm
				h.peers[i].MessageReceived = true

				fmt.Printf("Recv: handleConfirmationRequest PeerId - %v\n", request.Id)
				counter++
				break
			}
		}

		if counter == len(h.peers) {
			broadcastDiceMixResponse(h, commons.S_TX_CONFIRMATION, "Confirmation Response", "")
		}
	}
}
