package messages

import (
	"finance-chat/src/domain/messages"
	"github.com/gin-gonic/gin"
)

func AddRoutes(engine *gin.Engine) {
	// Configure websocket route
	engine.GET("/wss", messages.HandleConnections)
}

func HandleMessages() {
	// Start listening for incoming chat messages
	go messages.HandleMessages()
}
