package server

import (
	"sync"

	"../messages"
	"../utils"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// Client is a middleman between the websocket connection and the hub.
type client struct {
	hub *hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// to isolate clients from parallel dicemix executations
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
				log.Info("INCOMING C_JOIN_REQUEST - SUCCESSFUL")
			} else {
				log.Info("INCOMING C_JOIN_REQUEST - FAILED")
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				log.Info("INCOMING - USER UN-REGISTRATION - ", h.clients[client])
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

	// generates a random user id for new client
	userID := utils.RandInt31()

	// send registration response to client
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

	// map client with its userID
	h.clients[client] = userID

	// store client in waiting queue till |clients| == MinPeers
	h.waitingQueue = append(h.waitingQueue, userID)

	counter++

	// if min number of clients registers then start Dicemix protocol
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
	// generate session id for clients involved in current dicemix execution
	sessionID := utils.RandUint64()

	// create new run
	run := newRun()
	run.peers = make([]*messages.PeersInfo, utils.MinPeers)
	run.sessionID = sessionID

	// copy peersInfo from waiting queue to run
	for i, userID := range h.waitingQueue {
		run.peers[i] = &messages.PeersInfo{Id: userID}
		run.peers[i].MessageReceived = true
	}

	// creates an association between sessionID and run
	h.runs[sessionID] = run

	// clear the waiting queue
	h.waitingQueue = make([]int32, 0)

	// broadcasts - initiates DiceMix-Light protocol
	go broadcastDiceMixResponse(h, sessionID, messages.S_START_DICEMIX, "Initiate DiceMix Protocol", "")
}
