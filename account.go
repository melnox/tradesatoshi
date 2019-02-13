package tradesatoshi

import "github.com/shopspring/decimal"

type WithdrawalID int64
type DepositID int64

type GetBalanceResponse struct {
	Success string          `json:"success"`
	Message string          `json:"message"`
	Result  CurrencyBalance `json:"result"`
}

type GetBalancesResponse struct {
	Success string            `json:"success"`
	Message string            `json:"message"`
	Result  []CurrencyBalance `json:"result"`
}

type SubmitWithdrawResponse struct {
	Success string                 `json:"success"`
	Message string                 `json:"message"`
	Result  SubmitWithdrawalResult `json:"result"`
}

type SubmitWithdrawalResult struct {
	WithdrawalID WithdrawalID `json:"WithdrawalId"`
}

type GetDepositsResponse struct {
	Success string          `json:"success"`
	Message string          `json:"message"`
	Result  []DepositResult `json:"result"`
}

type DepositResult struct {
	ID            DepositID       `json:"Id"`
	Currency      string          `json:"Currency"`
	CurrencyLong  string          `json:"CurrencyLong"`
	Amount        decimal.Decimal `json:"Amount"`
	Status        string          `json:"Status"`
	Txid          string          `json:"Txid"`
	Confirmations int             `json:"Confirmations"`
	Timestamp     Time            `json:"Timestamp"`
}

type GetWithdrawalsResponse struct {
	Success string             `json:"success"`
	Message string             `json:"message"`
	Result  []WithdrawalResult `json:"result"`
}

type WithdrawalResult struct {
	ID            WithdrawalID    `json:"Id"`
	Currency      string          `json:"Currency"`
	CurrencyLong  string          `json:"CurrencyLong"`
	Amount        decimal.Decimal `json:"Amount"`
	Fee           decimal.Decimal `json:"Fee"`
	Status        string          `json:"Status"`
	Txid          string          `json:"Txid"`
	Confirmations int             `json:"Confirmations"`
	Timestamp     Time            `json:"Timestamp"`
	IsApi         bool            `json:"IsApi"`
}

type SubmitTransferResponse struct {
	Success string               `json:"success"`
	Message string               `json:"message"`
	Result  SubmitTransferResult `json:"result"`
}

type SubmitTransferResult struct {
	Data string `json:"data"`
}

type SubmitTipResponse struct {
	Success string          `json:"success"`
	Message string          `json:"message"`
	Result  SubmitTipResult `json:"result"`
}

type SubmitTipResult struct {
	Data string `json:"data"`
}
