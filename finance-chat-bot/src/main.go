package main

import (
	"finance-chat-bot/src/configuration"
	"finance-chat-bot/src/domain/stock/delivery"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := createEngine()

	addRoutes(engine)

	startEngine(engine)
}

func createEngine() *gin.Engine {
	return gin.Default()
}

func addRoutes(engine *gin.Engine) {
	stock.AddRoutes(engine)
}

func startEngine(engine *gin.Engine) {
	// Start the server on localhost port 8001 and log any errors
	err := engine.Run(configuration.Port)
	if err != nil {
		panic(err)
	}
}
