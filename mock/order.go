package mock

import (
	"github.com/shaunmza/tradesatoshi"
	"github.com/shopspring/decimal"
)

type OrderService struct {
	GetOrderFn      func(OrderID *tradesatoshi.OrderID) (*tradesatoshi.OrderResponse, error)
	GetOrderInvoked bool

	GetOrdersFn      func(marketSymbol *tradesatoshi.MarketSymbol, count int) (*tradesatoshi.OrdersResponse, error)
	GetOrdersInvoked bool

	SubmitOrderFn      func(marketSymbol *tradesatoshi.MarketSymbol, orderType *tradesatoshi.OrderType, amount decimal.Decimal, price decimal.Decimal) (*tradesatoshi.SubmitOrderResponse, error)
	SubmitOrderInvoked bool

	CancelOrderFn      func(cancelOrderType *tradesatoshi.CancelOrderType, orderID *tradesatoshi.OrderID, marketSymbol *tradesatoshi.MarketSymbol) (*tradesatoshi.SubmitOrderResponse, error)
	CancelOrderInvoked bool

	GetTradeHistoryFn      func(marketSymbol *tradesatoshi.MarketSymbol, count int, page int) (*tradesatoshi.TradeHistoryResponse, error)
	GetTradeHistoryInvoked bool
}

func (s *OrderService) GetOrder(OrderID *tradesatoshi.OrderID) (*tradesatoshi.OrderResponse, error) {
	s.GetOrderInvoked = true
	return s.GetOrderFn(OrderID)
}

func (s *OrderService) GetOrders(marketSymbol *tradesatoshi.MarketSymbol, count int) (*tradesatoshi.OrdersResponse, error) {
	s.GetOrdersInvoked = true
	return s.GetOrdersFn(marketSymbol, count)
}

func (s *OrderService) SubmitOrder(marketSymbol *tradesatoshi.MarketSymbol, orderType *tradesatoshi.OrderType, amount decimal.Decimal, price decimal.Decimal) (*tradesatoshi.SubmitOrderResponse, error) {
	s.SubmitOrderInvoked = true
	return s.SubmitOrderFn(marketSymbol, orderType, amount, price)
}

func (s *OrderService) CancelOrder(cancelOrderType *tradesatoshi.CancelOrderType, orderID *tradesatoshi.OrderID, marketSymbol *tradesatoshi.MarketSymbol) (*tradesatoshi.SubmitOrderResponse, error) {
	s.CancelOrderInvoked = true
	return s.CancelOrderFn(cancelOrderType, orderID, marketSymbol)
}

func (s *OrderService) GetTradeHistory(marketSymbol *tradesatoshi.MarketSymbol, count int, page int) (*tradesatoshi.TradeHistoryResponse, error) {
	s.GetTradeHistoryInvoked = true
	return s.GetTradeHistoryFn(marketSymbol, count, page)
}
