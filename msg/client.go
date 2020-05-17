package msg

import (
	"log"
	"strconv"

	//tc "github.com/athul/termchat"
	wss "github.com/gorilla/websocket"
)

//Hub data type
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

//Client dtaa
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *wss.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

//NewHub instantiates a new Hub struct
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

//cName=fmt.Sprintf(`%s joined`, tc.ClientName)
var clients = 0

//Run startes the hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			clients++
			log.Printf("Server join New, Total Clients are " + strconv.Itoa(clients))
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
