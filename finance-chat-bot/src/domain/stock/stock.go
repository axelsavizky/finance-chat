package stock

import (
	"crypto/tls"
	"finance-chat-bot/src/configuration"
	"finance-chat-bot/src/domain/stock/entities"
	"finance-chat-bot/src/domain/stock/repository/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HandleGetStockValue(c *gin.Context) {
	stockCode := c.Params.ByName("stock_code")
	if stockCode == "" {
		c.JSON(http.StatusBadRequest, entities.StockValueErrorResponse{Error: "empty stock code"})
	}

	statusCode, stockData, err := repository.GetStockValue(stockCode)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if stockData != nil {
		msg := fmt.Sprintf("%s quote is $%s per share", strings.ToUpper(stockData.StockCode), stockData.StockCloseValue)

		err = sendMessageToWebSocket(msg)
		if err != nil {
			log.Fatal(err)
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	if statusCode == http.StatusNotFound {
		msg := fmt.Sprintf("stock %s not found", stockCode)

		err = sendMessageToWebSocket(msg)
		if err != nil {
			log.Fatal(err)
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	c.Status(statusCode)
	return
}

func sendMessageToWebSocket(msg string) error {
	originURL, err := url.Parse(fmt.Sprintf("https://%s/", configuration.ChatWebsocketAddress))
	if err != nil {
		return err
	}
	locationURL, err := url.Parse(fmt.Sprintf("wss://%s/wss", configuration.ChatWebsocketAddress))
	if err != nil {
		return err
	}
	config := websocket.Config{
		Protocol: []string{"13"},
		Origin: originURL,
		Location: locationURL,
		TlsConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	ws, err := websocket.DialConfig(&config)
	if err != nil {
		return err
	}

	message := entities.Message{
		Username: "chat-bot",
		Message:  msg,
		Time:     time.Now().Format("15:04:05"),
	}

	err = websocket.JSON.Send(ws, &message)
	if err != nil {
		return err
	}

	return nil
}
