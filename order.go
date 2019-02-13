package tradesatoshi

import "github.com/shopspring/decimal"

type OrderID int64

type OrderType string

type Order struct {
	ID        OrderID         `json:"id"`
	TimeStamp Time            `json:"timeStamp"`
	Quantity  decimal.Decimal `json:"quantity"`
	Price     decimal.Decimal `json:"price"`
	OrderType OrderType       `json:"orderType"`
	Total     decimal.Decimal `json:"total"`
}

type OrderBookResult struct {
	Success string       `json:"success"`
	Message string       `json:"message"`
	Result  MarketStatus `json:"result"`
}

type OrderBookType struct {
	Buy  []OrderBookAggregate `json:"buy"`
	Sell []OrderBookAggregate `json:"sell"`
}

type OrderBookAggregate struct {
	Quantity decimal.Decimal `json:"quantity"`
	Rate     decimal.Decimal `json:"rate"`
}
