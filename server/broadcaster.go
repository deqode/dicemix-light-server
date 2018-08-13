package server

import (
	"fmt"
	"time"

	"../commons"
	"github.com/golang/protobuf/proto"
)

func broadcastDiceMixResponse(h *Hub, code uint32, message string, errMessage string) {
	// TODO : if not BR yet then only
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

func broadcastDCExponentialResponse(h *Hub, code uint32, message string, errMessage string) {
	// TODO : if not BR yet then only
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

func broadcast(h *Hub, message []byte, err error, statusCode uint32, statusMessage string) {
	if err != nil {
		fmt.Println(err)
	}

	for client := range h.clients {
		select {
		case client.send <- message:
		default:
			close(client.send)
			delete(h.clients, client)
		}
	}

	h.lastRoundUUID = h.roundUUID[statusCode]

	go func() {
		select {
		case <-time.After(responseWait * time.Second):
			registerDelayHandler(h, statusCode, statusMessage)
		}
	}()
}
