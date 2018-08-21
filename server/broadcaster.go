package server

import (
	"os"
	"time"

	"../messages"
	"../utils"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

// Removes offline peers
// Broadcasts message to active peers
// Broadcasts responses for -
// S_START_DICEMIX, S_KEY_EXCHANGE, S_SIMPLE_DC_VECTOR, S_TX_CONFIRMATION
func broadcastDiceMixResponse(h *hub, sessionID uint64, state uint32, message string, errMessage string) {
	// removes offline peers
	// returns true if removed any offline peers
	res := filterPeers(h, sessionID)

	if res {
		// if any P_Excluded go back to KE Stage
		if state == messages.S_SIMPLE_DC_VECTOR {
			broadcastKEResponse(h, sessionID)
			return
		}
	}

	// broadcast response to all active peers
	peers, err := proto.Marshal(&messages.DiceMixResponse{
		Code:      state,
		SessionId: sessionID,
		Peers:     h.runs[sessionID].peers,
		Timestamp: utils.Timestamp(),
		Message:   message,
		Err:       errMessage,
	})

	broadcast(h, sessionID, peers, err, state)
}

// Removes offline peers and broadcasts message to active peers
// Broadcasts responses for -
// S_EXP_DC_VECTOR
func broadcastDCExponentialResponse(h *hub, sessionID uint64, state uint32, message string, errMessage string) {
	// removes offline peers
	// returns true if removed any offline peers
	res := filterPeers(h, sessionID)

	if res {
		if state == messages.S_EXP_DC_VECTOR {
			broadcastKEResponse(h, sessionID)
			return
		}
	}

	// broadcast response to all active peers
	peers, err := proto.Marshal(&messages.DCExpResponse{
		Code:      state,
		SessionId: sessionID,
		Roots:     iDcNet.SolveDCExponential(h.runs[sessionID].peers),
		Timestamp: utils.Timestamp(),
		Message:   message,
		Err:       errMessage,
	})

	broadcast(h, sessionID, peers, err, state)
}

// creates a new run by broadcast KE Exchange Respose to active peers
// when previous run has been discarded due to some offline peers
func broadcastKEResponse(h *hub, sessionID uint64) {
	peers, err := proto.Marshal(&messages.DiceMixResponse{
		Code:      messages.S_KEY_EXCHANGE,
		SessionId: sessionID,
		Peers:     h.runs[sessionID].peers,
		Timestamp: utils.Timestamp(),
		Message:   "Key Exchange Response",
		Err:       "",
	})

	broadcast(h, sessionID, peers, err, messages.S_KEY_EXCHANGE)
}

// sent if all peers agrees to continue
// and have submitted confirmations
func broadcastTXDone(h *hub, sessionID uint64) {
	peers, err := proto.Marshal(&messages.TXDoneResponse{
		Code:      messages.S_TX_SUCCESSFUL,
		SessionId: sessionID,
		Messages:  h.runs[sessionID].peers[0].Messages,
		Timestamp: utils.Timestamp(),
		Message:   "DiceMix Successful Response",
		Err:       "",
	})

	broadcast(h, sessionID, peers, err, messages.S_TX_SUCCESSFUL)
}

// sent if all peers agrees to continue
// and have submitted confirmations
func broadcastKESKRequest(h *hub, sessionID uint64) {
	peers, err := proto.Marshal(&messages.InitiaiteKESK{
		Code:      messages.S_KESK_REQUEST,
		SessionId: sessionID,
		Timestamp: utils.Timestamp(),
		Message:   "Blame - send your kesk to identify culprit",
		Err:       "",
	})

	broadcast(h, sessionID, peers, err, messages.S_KESK_REQUEST)
}

// Broadcasts messages to active peers
// sets lastRoundUUID to roundUUID of current Response
// Registers a go routine to handled non responsive peers
func broadcast(h *hub, sessionID uint64, message []byte, err error, statusCode uint32) {
	checkError(err)

	// minimum peer check
	if len(h.runs[sessionID].peers) < 2 {
		log.Fatal("MinPeers: Less than two peers")
		os.Exit(1)
	}

	// wait for 1 sec before broadcasting
	time.Sleep(time.Second)

	for _, peerInfo := range h.runs[sessionID].peers {
		client, _ := getClient(h.clients, peerInfo.Id)
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(h.clients, client)
		}

		log.Info("SENT: SessionId - ", sessionID, ", ResponseCode - ", statusCode, ", PeerId - ", peerInfo.Id)
	}

	// predict next expected RequestCode from client againts current ResponseCode
	h.runs[sessionID].nextState = nextState(int(statusCode))

	log.Info("SessionId - ", sessionID, ", Expected Next State - ", h.runs[sessionID].nextState)

	// registers a go-routine to handle offline peers
	go registerWorker(h, sessionID, uint32(h.runs[sessionID].nextState))
}

// registers a go-routine to handle offline peers
func registerWorker(h *hub, sessionID uint64, statusCode uint32) {
	select {
	// wait for responseWait seconds then run registerDelayHandler()
	case <-time.After(utils.ResponseWait * time.Second):
		registerDelayHandler(h, sessionID, int(statusCode))
	}
}
