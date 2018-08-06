package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var dialer = websocket.Dialer{} // use default options

func main() {
	flag.Parse()
	log.SetFlags(0)

	var connection = connect()
	listener(connection)

	defer connection.Close()

	for {
	}
}

func connect() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())
	c, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	return c
}

func listener(c *websocket.Conn) {
	for {
		_, message, err := c.ReadMessage()

		// serialized := &commons.JoinTx{}
		// err = proto.Unmarshal(message, serialized)
		if err != nil {
			log.Fatal("unmarshaling error: ", err)
		}
		log.Printf("recv: %v", string(message))

	}
}

func send(c *websocket.Conn) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter mesage - ")
	var input, _ = reader.ReadString('\n')
	message := strings.Replace(input, "\n", "", -1)

	err := c.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write:", err)
	}
}
