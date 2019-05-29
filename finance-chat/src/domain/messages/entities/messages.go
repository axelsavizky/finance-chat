package entities

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	Clients   = make(map[*websocket.Conn]bool) // connected clients
	Broadcast = make(chan Message)             // broadcast channel

	// Configure the upgrader
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// Define our message object
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Time     string `json:"time"`
}
