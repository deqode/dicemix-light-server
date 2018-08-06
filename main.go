package main

import (
	"flag"
	"log"
	"net/http"

	"./server"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	connection := server.NewConnection()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		connection.Register(w, r)
	})

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
