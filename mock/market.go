package mock

import "github.com/shaunmza/tradesatoshi"

type MarketService struct {
	GetTickerFn      func(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.TickerResult, error)
	GetTickerInvoked bool

	GetMarketStatusFn      func(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.MarketStatusResult, error)
	GetMarketStatusInvoked bool

	GetMarketHistoryFn      func(symbol tradesatoshi.MarketSymbol, count int) (*tradesatoshi.MarketHistoryResult, error)
	GetMarketHistoryInvoked bool

	GetMarketSummaryFn      func(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.MarketSummaryResult, error)
	GetMarketSummaryInvoked bool

	GetMarketSummariesFn      func() (*tradesatoshi.MarketSummariesResult, error)
	GetMarketSummariesInvoked bool

	GetOrderBookFn      func(symbol tradesatoshi.MarketSymbol, orderType string, count int) (*tradesatoshi.OrderBookResult, error)
	GetOrderBookInvoked bool
}

func (s *MarketService) GetTicker(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.TickerResult, error) {
	s.GetTickerInvoked = true
	return s.GetTickerFn(symbol)
}
func (s *MarketService) GetMarketStatus(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.MarketStatusResult, error) {
	s.GetMarketStatusInvoked = true
	return s.GetMarketStatusFn(symbol)
}
func (s *MarketService) GetMarketHistory(symbol tradesatoshi.MarketSymbol, count int) (*tradesatoshi.MarketHistoryResult, error) {
	s.GetMarketHistoryInvoked = true
	return s.GetMarketHistoryFn(symbol, count)
}
func (s *MarketService) GetMarketSummary(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.MarketSummaryResult, error) {
	s.GetMarketSummaryInvoked = true
	return s.GetMarketSummaryFn(symbol)
}
func (s *MarketService) GetMarketSummaries() (*tradesatoshi.MarketSummariesResult, error) {
	s.GetMarketSummariesInvoked = true
	return s.GetMarketSummariesFn()
}
func (s *MarketService) GetOrderBook(symbol tradesatoshi.MarketSymbol, orderType string, count int) (*tradesatoshi.OrderBookResult, error) {
	s.GetOrderBookInvoked = true
	return s.GetOrderBookFn(symbol, orderType, count)
}
