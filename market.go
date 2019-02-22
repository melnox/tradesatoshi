package tradesatoshi

import "github.com/shopspring/decimal"

type MarketSymbol string

type TickerResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  Market `json:"result"`
}

type MarketStatusResult struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Result  MarketStatus `json:"result"`
}

type MarketStatus struct {
	MarketStatus  string `json:"marketStatus"`
	StatusMessage string `json:"statusMessage"`
}

type MarketHistoryResult struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Result  []Order `json:"result"`
}

type Market struct {
	Market         MarketSymbol    `json:"market"`
	High           decimal.Decimal `json:"high"`
	Low            decimal.Decimal `json:"low"`
	Volume         decimal.Decimal `json:"volume"`
	BaseVolume     decimal.Decimal `json:"baseVolume"`
	Last           decimal.Decimal `json:"last"`
	Bid            decimal.Decimal `json:"bid"`
	Ask            decimal.Decimal `json:"ask"`
	OpenBuyOrders  int             `json:"openBuyOrders"`
	OpenSellOrders int             `json:"openSellOrders"`
	MarketStatus   string          `json:"marketStatus"`
	Change         decimal.Decimal `json:"change"`
}

type MarketSummaryResult struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Result  Market `json:"result"`
}

type MarketSummariesResult struct {
	Success string   `json:"success"`
	Message string   `json:"message"`
	Result  []Market `json:"result"`
}
