package tradesatoshi

import "github.com/shopspring/decimal"

type OrderID int64

type OrderType string

type CancelOrderType string

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

type OrderResponse struct {
	Success string       `json:"success"`
	Message string       `json:"message"`
	Result  AccountOrder `json:"result"`
}

type OrdersResponse struct {
	Success string         `json:"success"`
	Message string         `json:"message"`
	Result  []AccountOrder `json:"result"`
}

type AccountOrder struct {
	ID        OrderID         `json:"id"`
	Market    MarketSymbol    `json:"market"`
	Type      OrderType       `json:"Type"`
	Amount    decimal.Decimal `json:"Amount"`
	Rate      decimal.Decimal `json:"Rate"`
	Remaining decimal.Decimal `json:"Remaining"`
	Total     decimal.Decimal `json:"Total"`
	Status    string          `json:"Status"`
	Timestamp Time            `json:"Timestamp"`
	IsApi     bool            `json:"IsApi"`
}

type SubmitOrderResponse struct {
	Success string            `json:"success"`
	Message string            `json:"message"`
	Result  SubmitOrderResult `json:"result"`
}

type SubmitOrderResult struct {
	ID     OrderID   `json:"OrderId"`
	Filled []OrderID `json:"Filled"`
}

type CancelOrderResponse struct {
	Success string            `json:"success"`
	Message string            `json:"message"`
	Result  CancelOrderResult `json:"result"`
}

type CancelOrderResult struct {
	CancelledOrders []OrderID `json:"CanceledOrders"`
}

type TradeHistoryResponse struct {
	Success      string         `json:"success"`
	Message      string         `json:"message"`
	TotalRecords int            `json:"totalRecords"`
	Result       []AccountOrder `json:"result"`
}
