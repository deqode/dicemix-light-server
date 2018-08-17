package server

import (
	"log"
	"os"
	"time"

	"../messages"
	"../utils"
	"github.com/golang/protobuf/proto"
)

// Removes offline peers
// Broadcasts message to active peers
// Broadcasts responses for -
// S_START_DICEMIX, S_KEY_EXCHANGE, S_SIMPLE_DC_VECTOR, S_TX_CONFIRMATION
func broadcastDiceMixResponse(h *Hub, state uint32, message string, errMessage string) {
	// removes offline peers
	// returns true if removed any offline peers
	res := filterPeers(h)

	if res {
		// if any P_Excluded go back to KE Stage
		if state == messages.S_SIMPLE_DC_VECTOR {
			broadcastKEResponse(h)
			return
		}
	}

	// broadcast response to all active peers
	peers, err := proto.Marshal(&messages.DiceMixResponse{
		Code:      state,
		Peers:     h.peers,
		Timestamp: utils.Timestamp(),
		Message:   message,
		Err:       errMessage,
	})

	broadcast(h, peers, err, state)
}

// Removes offline peers and broadcasts message to active peers
// Broadcasts responses for -
// S_EXP_DC_VECTOR
func broadcastDCExponentialResponse(h *Hub, state uint32, message string, errMessage string) {
	// removes offline peers
	// returns true if removed any offline peers
	res := filterPeers(h)

	if res {
		if state == messages.S_EXP_DC_VECTOR {
			broadcastKEResponse(h)
			return
		}
	}

	// broadcast response to all active peers
	peers, err := proto.Marshal(&messages.DCExpResponse{
		Code:      state,
		Roots:     iDcNet.SolveDCExponential(h.peers),
		Timestamp: utils.Timestamp(),
		Message:   message,
		Err:       errMessage,
	})

	broadcast(h, peers, err, state)
}

// creates a new run by broadcast KE Exchange Respose to active peers
// when previous run has been discarded due to some offline peers
func broadcastKEResponse(h *Hub) {
	peers, err := proto.Marshal(&messages.DiceMixResponse{
		Code:      messages.S_KEY_EXCHANGE,
		Peers:     h.peers,
		Timestamp: utils.Timestamp(),
		Message:   "Key Exchange Response",
		Err:       "",
	})

	broadcast(h, peers, err, messages.S_KEY_EXCHANGE)
}

// sent if all peers agrees to continue
// and have submitted confirmations
func broadcastTXDone(h *Hub) {
	peers, err := proto.Marshal(&messages.TXDoneResponse{
		Code:      messages.S_TX_SUCCESSFUL,
		Messages:  h.peers[0].Messages,
		Timestamp: utils.Timestamp(),
		Message:   "DiceMix Successful Response",
		Err:       "",
	})

	broadcast(h, peers, err, messages.S_TX_SUCCESSFUL)
}

// sent if all peers agrees to continue
// and have submitted confirmations
func broadcastKESKRequest(h *Hub) {
	peers, err := proto.Marshal(&messages.InitiaiteKESK{
		Code:      messages.S_KESK_REQUEST,
		Timestamp: utils.Timestamp(),
		Message:   "Blame - send your kesk to identify culprit",
		Err:       "",
	})

	broadcast(h, peers, err, messages.S_KESK_REQUEST)
}

// Broadcasts messages to active peers
// sets lastRoundUUID to roundUUID of current Response
// Registers a go routine to handled non responsive peers
func broadcast(h *Hub, message []byte, err error, statusCode uint32) {
	checkError(err)

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
	go registerWorker(h, uint32(h.nextState))
}

// registers a go-routine to handle offline peers
func registerWorker(h *Hub, statusCode uint32) {
	select {
	// wait for responseWait seconds then run registerDelayHandler()
	case <-time.After(utils.ResponseWait * time.Second):
		registerDelayHandler(h, int(statusCode))
	}
}
