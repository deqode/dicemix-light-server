package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"../dc"
	"../utils"
	"github.com/gorilla/websocket"
)

// using expose interfaces
var iDcNet dc.DC

type connection struct {
	hub *hub
	Server
}

// NewConnection creates a new Server instance
func NewConnection() Server {
	iDcNet = dc.NewDCNetwork()

	hub := newHub()
	go hub.listener()

	return &connection{hub: hub}
}

// Register handles websocket requests from the peer.
func (s *connection) Register(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &client{hub: s.hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writeMessage()
	go client.readMessage()
}

// readMessage pumps messages from the websocket connection to the hub.
//
// The application runs readMessage in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *client) readMessage() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadDeadline(time.Now().Add(utils.PongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(utils.PongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.hub.request <- message
	}
}

// writeMessage pumps messages from the hub to the websocket connection.
//
// A goroutine running writeMessage is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *client) writeMessage() {
	ticker := time.NewTicker(utils.PingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(utils.WriteWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(utils.Newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(utils.WriteWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// remove a peer from set of all peers
func removePeer(h *hub, id int32) {
	// if client is offline and not submitted response
	if client, ok := mapkey(h.clients, id); ok {
		// remove offline peers from clients
		fmt.Printf("USER UN-REGISTRATION - %v\n", id)
		delete(h.clients, client)
		close(client.send)
	}
}

// checks for any potential errors
// exists program if one found
func checkError(err error) {
	if err != nil {
		log.Fatal("Error Occured:", err)
		os.Exit(1)
	}
}
