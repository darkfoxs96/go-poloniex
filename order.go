package poloniex

type OrderBook struct {
	Asks     [][]interface{} `json:"asks"`
	Bids     [][]interface{} `json:"bids"`
	IsFrozen int             `json:"isFrozen,string"`
	Error    string          `json:"error"`
}

// This can probably be implemented using UnmarshalJSON
/*
type OrderBook struct {
	Bids     []Orderb `json:"bids"`
	Asks     []Orderb `json:"asks"`
	IsFrozen int      `json:"isFrozen,string"`
}
type Orderb struct {
	Rate     string
	Quantity float64
}
*/

type OpenOrder struct {
	OrderNumber int     `json:"orderNumber,string"`
	Type        string  `json:"type"`
	Rate        float64 `json:"rate,string"`
	Amount      float64 `json:"amount,string"`
	Total       float64 `json:"total,string"`
	Date        string  `json:"date"`
}

type OrderStatus struct {
	Status         string  `json:"status"` // 'Open'
	Rate           float64 `json:"rate,string"`
	Amount         float64 `json:"amount,string"`
	CurrencyPair   string  `json:"currencyPair"`
	Date           string  `json:"date"`
	Total          float64 `json:"total,string"`
	Type           string  `json:"type"` // 'buy' or 'sell'
	StartingAmount float64 `json:"startingAmount,string"`
}
