package server

import (
	"fmt"
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
func broadcastDiceMixResponse(h *Hub, code uint32, message string, errMessage string) {
	// removes offline peers
	filterPeers(h)

	peers, err := proto.Marshal(&commons.DiceMixResponse{
		Code:        code,
		Peers:       h.peers,
		MessageUUID: h.roundUUID[code],
		Timestamp:   timestamp(),
		Message:     message,
		Err:         errMessage,
	})
	broadcast(h, peers, err, code, message)
}

// Removes offline peers
// Broadcasts message to active peers
// Broadcasts responses for -
// S_EXP_DC_VECTOR
func broadcastDCExponentialResponse(h *Hub, code uint32, message string, errMessage string) {
	// remove offline peers
	filterPeers(h)

	peers, err := proto.Marshal(&commons.DCExpResponse{
		Code:        code,
		Roots:       iDcNet.SolveDCExponential(h.peers),
		MessageUUID: h.roundUUID[code],
		Timestamp:   timestamp(),
		Message:     message,
		Err:         errMessage,
	})

	broadcast(h, peers, err, code, message)
}

// Broadcasts messages to active peers
// sets lastRoundUUID to roundUUID of current Response
// Registers a go routine to handled non responsive peers
func broadcast(h *Hub, message []byte, err error, statusCode uint32, statusMessage string) {
	if err != nil {
		fmt.Println(err)
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

	// sets lastRoundUUID to roundUUID of current Response
	h.lastRoundUUID = h.roundUUID[statusCode]

	// Registers a go routine to handled non responsive peers
	go func() {
		select {
		// wait for responseWait seconds then run registerDelayHandler()
		case <-time.After(responseWait * time.Second):
			registerDelayHandler(h, statusCode, statusMessage)
		}
	}()
}
