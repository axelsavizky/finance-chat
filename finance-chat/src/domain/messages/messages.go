package messages

import (
	"finance-chat/src/domain/messages/entities"
	"finance-chat/src/domain/messages/repository/http"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

const (
	BotCommandPrefix = "/stock="
)

func HandleConnections(c *gin.Context) {
	// Upgrade initial GET request to a websocket
	ws, err := entities.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	entities.Clients[ws] = true

	for {
		var msg entities.Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			//If there is some kind of error with reading from the socket, we assume the client has disconnected for
			// some reason or another. We log the error and remove that client from our global "clients" map so we
			// don't try to read from or send new messages to that client.
			log.Printf("error: %v", err)
			delete(entities.Clients, ws)
			break
		}
		if strings.HasPrefix(msg.Message, BotCommandPrefix) {
			stockCode := strings.TrimPrefix(msg.Message, BotCommandPrefix)
			err = repository.SendStockCodeToChatBot(stockCode)
			if err != nil {
				c.Status(http.StatusBadGateway)
				return
			}

		}
		// Send the newly received message to the broadcast channel
		entities.Broadcast <- msg
	}
}

func HandleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-entities.Broadcast
		// Send it out to every client that is currently connected
		for client := range entities.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(entities.Clients, client)
			}
		}
	}
}
