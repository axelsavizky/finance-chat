package repository

import (
	"finance-chat/src/configuration"
	"fmt"
	"net/http"
)

const chatBotStockValuePath = "stock/%s"

func SendStockCodeToChatBot(stockCode string) error {
	stockValuePath := fmt.Sprintf(chatBotStockValuePath, stockCode)
	url := fmt.Sprintf(configuration.ChatBotBasePath, stockValuePath)
	_, err := http.Post(url, "application/json", nil)
	return err
}
