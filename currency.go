package tradesatoshi

import "github.com/shopspring/decimal"

type CurrencySymbol string

type Currency struct {
	Currency        string          `json:"currency"`
	CurrencyLong    string          `json:"currencyLong"`
	MinConfirmation int             `json:"minConfirmation"`
	TxFee           decimal.Decimal `json:"txFee"`
	Status          string          `json:"status"`
	StatusMessage   string          `json:"statusMessage"`
	MinBaseTrade    decimal.Decimal `json:"minBaseTrade"`
	IsTipEnabled    bool            `json:"isTipEnabled"`
	MinTip          decimal.Decimal `json:"minTip"`
	MaxTip          decimal.Decimal `json:"maxTip"`
}

type CurrencyResult struct {
	Success string     `json:"success"`
	Message string     `json:"success"`
	Result  []Currency `json:"result"`
}

type CurrenciesResult struct {
	Success string   `json:"success"`
	Message string   `json:"success"`
	Result  Currency `json:"result"`
}
