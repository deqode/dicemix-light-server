package server

import (
	"fmt"
	"os"

	"../commons"
	"github.com/golang/protobuf/proto"
)

func handleRequest(message []byte, h *Hub) {
	r := &commons.GenericRequest{}
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

	}
}

func handleKeyExchangeRequest(request *commons.KeyExchangeRequest, h *Hub) {
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
			peers, err := proto.Marshal(&commons.DiceMixResponse{
				Code:      commons.S_KEY_EXCHANGE,
				Peers:     h.peers,
				Timestamp: timestamp(),
				Message:   "Key Exchange Response",
				Err:       "",
			})

			if err != nil {
				fmt.Println(err)
			}

			for client := range h.clients {
				select {
				case client.send <- peers:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func handleDCExponentialRequest(request *commons.DCExpRequest, h *Hub) {
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

			peers, err := proto.Marshal(&commons.DCExpResponse{
				Code:      commons.S_EXP_DC_VECTOR,
				Roots:     iDcNet.SolveDCExponential(h.peers),
				Timestamp: timestamp(),
				Message:   "Solved DC Exponential Roots",
				Err:       "",
			})

			if err != nil {
				fmt.Println(err)
			}

			for client := range h.clients {
				select {
				case client.send <- peers:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}

		}
	}
}

func handleDCSimpleRequest(request *commons.DCSimpleRequest, h *Hub) {
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
				counter++
			}
		}

		if counter == maxPeers {

			peers, err := proto.Marshal(&commons.DiceMixResponse{
				Code:      commons.S_SIMPLE_DC_VECTOR,
				Peers:     h.peers,
				Timestamp: timestamp(),
				Message:   "DC Simple Response",
				Err:       "",
			})

			if err != nil {
				fmt.Println(err)
			}

			for client := range h.clients {
				select {
				case client.send <- peers:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}

		}
	}
}
