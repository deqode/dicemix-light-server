package server

import (
	"log"
	"os"
	"time"

	"../commons"
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
		if state == commons.S_SIMPLE_DC_VECTOR {
			broadcastKEResponse(h)
			return
		}
	}

	// broadcast response to all active peers
	peers, err := proto.Marshal(&commons.DiceMixResponse{
		Code:      state,
		Peers:     h.peers,
		Timestamp: timestamp(),
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
		if state == commons.S_EXP_DC_VECTOR {
			broadcastKEResponse(h)
			return
		}
	}

	// broadcast response to all active peers
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

// sent if all peers agrees to continue
// and have submitted confirmations
func broadcastTXDone(h *Hub) {
	peers, err := proto.Marshal(&commons.TXDoneResponse{
		Code:      commons.S_TX_SUCCESSFUL,
		Messages:  h.peers[0].Messages,
		Timestamp: timestamp(),
		Message:   "DiceMix Successful Response",
		Err:       "",
	})

	broadcast(h, peers, err, commons.S_TX_SUCCESSFUL)
}

// sent if all peers agrees to continue
// and have submitted confirmations
func broadcastKESKRequest(h *Hub) {
	peers, err := proto.Marshal(&commons.InitiaiteKESK{
		Code:      commons.S_KESK_REQUEST,
		Timestamp: timestamp(),
		Message:   "Blame - send your kesk to identify culprit",
		Err:       "",
	})

	broadcast(h, peers, err, commons.S_KESK_REQUEST)
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
	case <-time.After(responseWait * time.Second):
		registerDelayHandler(h, int(statusCode))
	}
}
