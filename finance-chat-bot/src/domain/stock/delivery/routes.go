package stock

import (
	"finance-chat-bot/src/domain/stock"
	"github.com/gin-gonic/gin"
)

func AddRoutes(engine *gin.Engine) {
	engine.POST("/stock/:stock_code", stock.HandleGetStockValue)
}
