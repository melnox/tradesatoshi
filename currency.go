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
	Message string     `json:"message"`
	Result  []Currency `json:"result"`
}

type CurrenciesResult struct {
	Success string   `json:"success"`
	Message string   `json:"message"`
	Result  Currency `json:"result"`
}

type CurrencyBalance struct {
	Currency        string          `json:"currency"`
	CurrencyLong    string          `json:"currencyLong"`
	Avaliable       decimal.Decimal `json:"Avaliable"`
	Total           decimal.Decimal `json:"Total"`
	HeldForTrades   decimal.Decimal `json:"HeldForTrades"`
	Unconfirmed     decimal.Decimal `json:"Unconfirmed"`
	PendingWithdraw decimal.Decimal `json:"PendingWithdraw"`
	Address         string          `json:"Address"`
}

type GenerateAddressResponse struct {
	Success string        `json:"success"`
	Message string        `json:"message"`
	Result  AddressResult `json:"result"`
}

type AddressResult struct {
	Currency  string `json:"currency"`
	Address   string `json:"Address"`
	PaymentId string `json:"PaymentId"`
}
