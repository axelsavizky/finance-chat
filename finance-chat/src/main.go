package main

import (
	"finance-chat/src/configuration"
	messagesDelivery "finance-chat/src/domain/messages/delivery"
	"finance-chat/src/domain/users/delivery"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := createEngineWithStaticFileServer()

	addRoutes(engine)

	messagesDelivery.HandleMessages()

	startEngine(engine)
}

func createEngineWithStaticFileServer() *gin.Engine {
	engine := gin.Default()

	// Create a simple file server
	engine.Static("/index", "./public")

	return engine
}

func addRoutes(engine *gin.Engine) {
	users.AddRoutes(engine)
	messagesDelivery.AddRoutes(engine)
}

func startEngine(engine *gin.Engine) {
	// Start the server on localhost port 8000 and log any errors
	err := engine.RunTLS(configuration.Port, configuration.CertFilePath, configuration.KeyFilePath)
	if err != nil {
		panic(err)
	}
}
