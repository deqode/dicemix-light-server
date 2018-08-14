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
		switch state {
		case commons.S_SIMPLE_DC_VECTOR, commons.S_TX_CONFIRMATION:
			state = commons.S_KEY_EXCHANGE
			message = "Key Exchange Response"
		}
	}

	peers, err := proto.Marshal(&commons.DiceMixResponse{
		Code:      state,
		Peers:     h.peers,
		Timestamp: timestamp(),
		Message:   message,
		Err:       errMessage,
	})

	log.Printf("BR for - %v\n", state)
	broadcast(h, peers, err, state, message)
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
			state = commons.S_KEY_EXCHANGE
			message = "Key Exchange Response"

			peers, err := proto.Marshal(&commons.DiceMixResponse{
				Code:      state,
				Peers:     h.peers,
				Timestamp: timestamp(),
				Message:   message,
				Err:       errMessage,
			})

			log.Printf("BR for - %v, %v\n", state, res)
			broadcast(h, peers, err, state, message)
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

	log.Printf("BR for - %v, %v\n", state, res)
	broadcast(h, peers, err, state, message)
}

// Broadcasts messages to active peers
// sets lastRoundUUID to roundUUID of current Response
// Registers a go routine to handled non responsive peers
func broadcast(h *Hub, message []byte, err error, statusCode uint32, statusMessage string) {
	if err != nil {
		fmt.Println(err)
	}

	if len(h.peers) < 2 {
		log.Fatal("Less than two peers")
		os.Exit(1)
	}

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

	h.nextState = nextState(int(statusCode))

	for _, state := range h.nextState {
		go registerWorker(h, uint32(state))
	}
	// sets lastRoundUUID to roundUUID of current Response
	// h.lastRoundUUID = h.roundUUID[statusCode]
}

// Registers a go routine to handle non responsive peers
func registerWorker(h *Hub, statusCode uint32) {

	log.Printf("Registered for - %v\n", statusCode)

	select {
	// wait for responseWait seconds then run registerDelayHandler()
	case <-time.After(responseWait * time.Second):
		registerDelayHandler(h, int(statusCode))
	}
}
