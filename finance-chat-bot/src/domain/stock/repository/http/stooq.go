package repository

import (
	"encoding/csv"
	"finance-chat-bot/src/domain/stock/entities"
	"fmt"
	"net/http"
)

const (
	stooqGetStockValueURL = "https://stooq.com/q/l/?s=%s&f=sd2t2ohlcv&h&e=csv"

	SymbolHeader      = "Symbol"
	CloseHeader       = "Close"
	CSVNotInformation = "N/D"
)

func GetStockValue(stockCode string) (int, *entities.StockData, error) {
	url := fmt.Sprintf(stooqGetStockValueURL, stockCode)
	response, err := http.Get(url)
	if err != nil {
		return 0, nil, fmt.Errorf("error %s has been raised when requesting to: %s", err.Error(), url)
	}
	if response.StatusCode < http.StatusOK || response.StatusCode > http.StatusIMUsed {
		// Status not 2xx
		return response.StatusCode, nil, nil
	}

	csvReader := csv.NewReader(response.Body)
	result, err := csvReader.ReadAll()
	if err != nil {
		return 0, nil, fmt.Errorf("error parsing csv response: %s", err.Error())
	}

	// We iterate the headers in order to know the position of the symbol and the close
	indexOfSymbol := -1
	indexOfClose := -1
	for index, header := range result[0] {
		if header == SymbolHeader {
			indexOfSymbol = index
		}
		if header == CloseHeader {
			indexOfClose = index
		}
	}

	if indexOfSymbol < 0 || indexOfClose < 0 {
		return 0, nil, fmt.Errorf("error on CSV format")
	}
	rawCloseValue := result[1][indexOfClose]
	if rawCloseValue == CSVNotInformation {
		return http.StatusNotFound, nil, nil
	}

	stockData := entities.StockData{
		StockCode:       result[1][indexOfSymbol],
		StockCloseValue: rawCloseValue,
	}

	return http.StatusOK, &stockData, nil
}
