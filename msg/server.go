package msg

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to TermChat, Chat from your terminal with Websockets")
}

func setupRoutes() {
	hub := NewHub()
	go hub.Run()
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})
}

// StartServer starts the ws server
func StartServer() {
	flag.Parse()
	log.Println("Server Started at port 8080")
	setupRoutes()
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
