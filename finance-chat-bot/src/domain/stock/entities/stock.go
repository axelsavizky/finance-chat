package entities

type (
	StockValueErrorResponse struct {
		Error string `json:"error"`
	}

	StockValueResponse struct {
		Message string `json:"message"`
	}

	StockData struct {
		StockCode       string `json:"stock_code"`
		StockCloseValue string `json:"stock_close_value"`
	}

	// Define our message object
	Message struct {
		Username string `json:"username"`
		Message  string `json:"message"`
		Time     string `json:"time"`
	}
)
