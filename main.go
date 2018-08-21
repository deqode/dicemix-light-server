package main

import (
	"flag"
	"net/http"

	"./server"
	log "github.com/sirupsen/logrus"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	// setup logger
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)

	flag.Parse()
	connection := server.NewConnection()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		connection.Register(w, r)
	})

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
