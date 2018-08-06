package server

import (
	"fmt"
	"log"
	"os"
	"time"

	"../commons"
	"github.com/golang/protobuf/proto"
)

func newHub() *Hub {
	return &Hub{
		request:    make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		rounds:     make(map[int]bool),
		peers:      make([]*commons.PeersInfo, maxPeers),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			if h.registration(client) {
				log.Printf("INCOMING C_JOIN_REQUEST - SUCCESSFUL")
			} else {
				log.Printf("INCOMING C_JOIN_REQUEST - FAILED")
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				log.Printf("INCOMING - USER UN-REGISTRATION - SUCCESSFUL")
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.request:
			request := &commons.GenericRequest{}
			if err := proto.Unmarshal(message, request); err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}

			handleRequest(message, h)
		}
	}
}

func (h *Hub) registration(client *Client) bool {
	h.Lock()
	defer h.Unlock()
	counter := int32(len(h.clients))

	if counter < maxPeers {
		h.clients[client] = true
		counter++

		registration, err := proto.Marshal(&commons.RegisterResponse{
			Code:      commons.S_JOIN_RESPONSE,
			Id:        counter,
			Timestamp: timestamp(),
			Message:   "Welcome to CoinShuffle++\nWaiting for other peers to join ...",
			Err:       "",
		})

		if err != nil {
			fmt.Println(err)
		}

		client.send <- registration
		h.peers[counter-1] = &commons.PeersInfo{Id: counter}

		if counter == maxPeers {
			// start DiceMix Light process
			h.rounds[joinTXRound] = true

			time.Sleep(100 * time.Millisecond)
			h.startDicemix()
		}

		return true
	}

	// send message
	registration, err := proto.Marshal(&commons.RegisterResponse{
		Code:      commons.S_JOIN_RESPONSE,
		Id:        -1,
		Timestamp: timestamp(),
		Message:   "",
		Err:       "Limit Exceeded. Kindly try after some time",
	})

	if err != nil {
		fmt.Println(err)
	}

	client.send <- registration

	return false
}

func (h *Hub) startDicemix() {
	peers, err := proto.Marshal(&commons.DiceMixResponse{
		Code:      commons.S_START_DICEMIX,
		Peers:     h.peers,
		Timestamp: timestamp(),
		Message:   "Initiate DiceMix Protocol",
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

func timestamp() string {
	return time.Now().String()
}
