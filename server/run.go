package server

import (
	"log"
	"sync"

	"../messages"
	"../utils"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

// Client is a middleman between the websocket connection and the hub.
type client struct {
	hub *hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

type run struct {
	sessionID uint64
	peers     []*messages.PeersInfo
	nextState int
	sync.Mutex
}

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type hub struct {
	clients      map[*client]int32
	runs         map[uint64]*run
	waitingQueue []int32
	request      chan []byte
	register     chan *client
	unregister   chan *client
	sync.Mutex
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func newHub() *hub {
	return &hub{
		clients:      make(map[*client]int32),
		runs:         make(map[uint64]*run),
		waitingQueue: make([]int32, 0),
		request:      make(chan []byte),
		register:     make(chan *client),
		unregister:   make(chan *client),
	}
}

func newRun() *run {
	return &run{
		sessionID: 0,
		peers:     make([]*messages.PeersInfo, utils.MinPeers),
		nextState: 0,
	}
}

// starts a run
// registers a peer when he want to participate in TX
// unregisters a peer
// listens for requests from peers and calls its corresponding handler
func (h *hub) listener() {
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
			handleRequest(message, h)
		}
	}
}

// adds a peer in h.peers if |h.peers| < MaxPeers
// send a failure message response to other peers
func (h *hub) registration(client *client) bool {
	h.Lock()
	defer h.Unlock()
	counter := int32(len(h.waitingQueue))

	// if counter >= utils.MinPeers {
	// 	registration, err := proto.Marshal(&messages.RegisterResponse{
	// 		Code:      messages.S_JOIN_RESPONSE,
	// 		Id:        -1,
	// 		Timestamp: utils.Timestamp(),
	// 		Message:   "",
	// 		Err:       "Limit Exceeded. Kindly try after some time",
	// 	})

	// 	checkError(err)
	// 	client.send <- registration
	// 	return false
	// }

	userID := utils.RandInt31()

	counter++
	registration, err := proto.Marshal(&messages.RegisterResponse{
		Code:      messages.S_JOIN_RESPONSE,
		SessionId: 0,
		Id:        userID,
		Timestamp: utils.Timestamp(),
		Message:   "Welcome to CoinShuffle++. Waiting for other peers to join ...",
		Err:       "",
	})

	checkError(err)
	client.send <- registration

	h.clients[client] = userID
	h.waitingQueue = append(h.waitingQueue, userID)

	if counter == utils.MinPeers {
		// start DiceMix Light process
		// initRoundUUID(h)
		h.startDicemix()
	}
	return true
}

// initiates DiceMix-Light protocol
// send all peers ID's
func (h *hub) startDicemix() {
	// generate session id
	// create run with waitingPeers
	// clear waitingPeers
	// br start dicemix response

	sessionID := utils.RandUint64()
	run := newRun()
	run.peers = make([]*messages.PeersInfo, utils.MinPeers)
	run.sessionID = sessionID

	for i, userID := range h.waitingQueue {
		run.peers[i] = &messages.PeersInfo{Id: userID}
		run.peers[i].MessageReceived = true
	}

	h.runs[sessionID] = run
	h.waitingQueue = make([]int32, 0)

	// fmt.Printf("Run - %v\n\n", h.runs[sessionID])

	go broadcastDiceMixResponse(h, sessionID, messages.S_START_DICEMIX, "Initiate DiceMix Protocol", "")
}
