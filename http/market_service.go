package http

import (
	"errors"
	"github.com/shaunmza/tradesatoshi"
	"net/url"
)

type MarketClient struct {
	MarketService
}

func NewMarketClient() *MarketClient {
	c := &MarketClient{}

	return c
}

var _ tradesatoshi.MarketService = &MarketService{}

type MarketService struct {
	MarketService *tradesatoshi.MarketService
	URL           url.URL
}

func (s *MarketService) GetTicker(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.TickerResult, error) {
	return nil, errors.New("Not implemented")
}
func (s *MarketService) GetMarketStatus(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.MarketStatusResult, error) {
	return nil, errors.New("Not implemented")
}
func (s *MarketService) GetMarketHistory(symbol tradesatoshi.MarketSymbol, count int) (*tradesatoshi.MarketHistoryResult, error) {
	return nil, errors.New("Not implemented")
}
func (s *MarketService) GetMarketSummary(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.MarketSummaryResult, error) {
	return nil, errors.New("Not implemented")
}
func (s *MarketService) GetMarketSummaries() (*tradesatoshi.MarketSummariesResult, error) {
	return nil, errors.New("Not implemented")
}
func (s *MarketService) GetOrderBook(symbol tradesatoshi.MarketSymbol, orderType string, count int) (*tradesatoshi.OrderBookResult, error) {
	return nil, errors.New("Not implemented")
}
