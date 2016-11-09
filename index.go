package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/dcb9/steamer/httpHandler"
	"github.com/dcb9/steamer/app"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8080", "http service address")
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	flag.Parse()

	http.Handle("/", http.HandlerFunc(httpHandler.Index))
	http.Handle("/r/", http.StripPrefix("/r/", http.FileServer(http.Dir(app.OutputDir))))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		httpHandler.WebSocket(conn, w, r)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
