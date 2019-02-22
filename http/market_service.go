package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shaunmza/tradesatoshi"
	"net/http"
	"net/url"
	"strconv"
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

func (s *MarketService) GetTicker(symbol, baseSymbol tradesatoshi.CurrencySymbol) (*tradesatoshi.Market, error) {
	ms := buildMarketSymbol(symbol, baseSymbol)

	res, err := http.Get("https://tradesatoshi.com/api/public/getticker?market=" + string(ms))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Not successful, return now.
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Invalid server response.")
	}

	var respBody *tradesatoshi.TickerResult
	if err := json.NewDecoder(res.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	if respBody.Success != true {
		return nil, errors.New(fmt.Sprintf("Request failed. Message: %s", respBody.Message))
	}

	emptyMarket := tradesatoshi.Market{}
	if respBody.Result == emptyMarket {
		return nil, errors.New("Request failed. No result received")
	}

	return &respBody.Result, err
}
func (s *MarketService) GetMarketStatus(symbol, baseSymbol tradesatoshi.CurrencySymbol) (*tradesatoshi.MarketStatus, error) {
	ms := buildMarketSymbol(symbol, baseSymbol)

	res, err := http.Get("https://tradesatoshi.com/api/public/GetMarketStatus?market=" + string(ms))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Not successful, return now.
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Invalid server response.")
	}

	var respBody *tradesatoshi.MarketStatusResult
	if err := json.NewDecoder(res.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	if respBody.Success != true {
		return nil, errors.New(fmt.Sprintf("Request failed. Message: %s", respBody.Message))
	}

	emptyMarketStatus := tradesatoshi.MarketStatus{}
	if respBody.Result == emptyMarketStatus {
		return nil, errors.New("Request failed. No result received")
	}

	return &respBody.Result, err
}
func (s *MarketService) GetMarketHistory(symbol, baseSymbol tradesatoshi.CurrencySymbol, count int) (*[]tradesatoshi.Order, error) {
	ms := buildMarketSymbol(symbol, baseSymbol)

	res, err := http.Get("https://tradesatoshi.com/api/public/getmarkethistory?market=" + string(ms) + "&count=" + strconv.Itoa(count))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Not successful, return now.
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Invalid server response.")
	}

	var respBody *tradesatoshi.MarketHistoryResult
	if err := json.NewDecoder(res.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	if respBody.Success != true {
		return nil, errors.New(fmt.Sprintf("Request failed. Message: %s", respBody.Message))
	}

	if len(respBody.Result) == 0 {
		return nil, errors.New("Request failed. No result received")
	}

	return &respBody.Result, err
}
func (s *MarketService) GetMarketSummary(symbol, baseSymbol tradesatoshi.CurrencySymbol) (*tradesatoshi.MarketSummaryResult, error) {
	return nil, errors.New("Not implemented")
}
func (s *MarketService) GetMarketSummaries() (*tradesatoshi.MarketSummariesResult, error) {
	return nil, errors.New("Not implemented")
}
func (s *MarketService) GetOrderBook(symbol, baseSymbol tradesatoshi.CurrencySymbol, orderType string, count int) (*tradesatoshi.OrderBookResult, error) {
	return nil, errors.New("Not implemented")
}

func buildMarketSymbol(symbol, baseSymbol tradesatoshi.CurrencySymbol) tradesatoshi.MarketSymbol {
	return tradesatoshi.MarketSymbol(symbol + "_" + baseSymbol)
}
