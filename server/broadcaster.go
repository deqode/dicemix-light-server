package server

import (
	"fmt"
	"log"
	"os"
	"time"

	"../commons"
	"github.com/golang/protobuf/proto"
)

// Removes offline peers
// Broadcasts message to active peers
// Broadcasts responses for -
// S_START_DICEMIX,
// S_KEY_EXCHANGE,
// S_SIMPLE_DC_VECTOR,
// S_TX_CONFIRMATION
func broadcastDiceMixResponse(h *Hub, state uint32, message string, errMessage string) {
	// removes offline peers
	// returns true if removed any offline peers
	res := filterPeers(h)

	if res {
		// if any P_Excluded trace back to KE Stage
		switch state {
		case commons.S_SIMPLE_DC_VECTOR, commons.S_TX_CONFIRMATION:
			broadcastKEResponse(h)
			return
		}
	}

	peers, err := proto.Marshal(&commons.DiceMixResponse{
		Code:      state,
		Peers:     h.peers,
		Timestamp: timestamp(),
		Message:   message,
		Err:       errMessage,
	})

	broadcast(h, peers, err, state)
}

// Removes offline peers
// Broadcasts message to active peers
// Broadcasts responses for -
// S_EXP_DC_VECTOR
func broadcastDCExponentialResponse(h *Hub, state uint32, message string, errMessage string) {
	// removes offline peers
	// returns true if removed any offline peers
	res := filterPeers(h)

	if res {
		switch state {
		case commons.S_EXP_DC_VECTOR:
			broadcastKEResponse(h)
			return
		}
	}

	peers, err := proto.Marshal(&commons.DCExpResponse{
		Code:      state,
		Roots:     iDcNet.SolveDCExponential(h.peers),
		Timestamp: timestamp(),
		Message:   message,
		Err:       errMessage,
	})

	broadcast(h, peers, err, state)
}

// creates a new run by broadcast KE Exchange Respose to active peers
// when previous run has been discarded due to some offline peers
func broadcastKEResponse(h *Hub) {
	peers, err := proto.Marshal(&commons.DiceMixResponse{
		Code:      commons.S_KEY_EXCHANGE,
		Peers:     h.peers,
		Timestamp: timestamp(),
		Message:   "Key Exchange Response",
		Err:       "",
	})

	broadcast(h, peers, err, commons.S_KEY_EXCHANGE)
}

// Broadcasts messages to active peers
// sets lastRoundUUID to roundUUID of current Response
// Registers a go routine to handled non responsive peers
func broadcast(h *Hub, message []byte, err error, statusCode uint32) {
	if err != nil {
		fmt.Println(err)
	}

	// minimum peer check
	if len(h.peers) < 2 {
		log.Fatal("MinPeers: Less than two peers")
		os.Exit(1)
	}

	// wait for 1 sec before broadcasting
	time.Sleep(time.Second)

	// Broadcasts messages to active peers
	for client := range h.clients {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}

	// predict next expected RequestCode from client againts current ResponseCode
	h.nextState = nextState(int(statusCode))

	// registers a go-routine to handle offline peers
	for _, state := range h.nextState {
		go registerWorker(h, uint32(state))
	}
}

// registers a go-routine to handle offline peers
func registerWorker(h *Hub, statusCode uint32) {
	select {
	// wait for responseWait seconds then run registerDelayHandler()
	case <-time.After(responseWait * time.Second):
		registerDelayHandler(h, int(statusCode))
	}
}
