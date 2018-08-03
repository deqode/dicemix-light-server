package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"

	"../commons"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	conn *websocket.Conn
}

type Hub struct {
	clients map[*Client]bool
	counter int32
	sync.Mutex
}

func main() {
	hub := initialize()

	http.HandleFunc("/ws/join", func(w http.ResponseWriter, r *http.Request) {
		handleNewUser(&hub, w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func initialize() Hub {
	return Hub{clients: make(map[*Client]bool), counter: 0}
}

func handleNewUser(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error joining transaction: ", err)
	}

	hub.Lock()

	defer hub.Unlock()
	if hub.counter != 3 {
		client := &Client{conn: conn}
		hub.clients[client] = true
		hub.counter++

		// send JoinTx response
		joinTx := commons.JoinTx{
			Status:  commons.S_JOIN_RESPONSE,
			ID:      hub.counter,
			Message: "Welcome to CoinShuffle++",
			Err:     "",
		}

		if err = conn.WriteJSON(joinTx); err != nil {
			fmt.Println(err)
		}
	} else {
		for client := range hub.clients {
			if err := client.conn.WriteJSON("{message : 'Start TX'}"); err != nil {
				fmt.Println(err)
			}
		}

		// send JoinTx response
		joinTx := commons.JoinTx{
			Status:  commons.S_JOIN_RESPONSE,
			ID:      -1,
			Message: "",
			Err:     "Limit Exceeded. Kindly try after some time",
		}

		if err = conn.WriteJSON(joinTx); err != nil {
			fmt.Println(err)
		}
	}
}
