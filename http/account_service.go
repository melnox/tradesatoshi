package http

import (
	"errors"
	"github.com/shaunmza/tradesatoshi"
	"github.com/shopspring/decimal"
	"net/url"
)

type AccountClient struct {
	AccountService
}

func NewAccountClient() *AccountClient {
	c := &AccountClient{}

	return c
}

var _ tradesatoshi.AccountService = &AccountService{}

type AccountService struct {
	AccountService *tradesatoshi.AccountService
	URL            url.URL
}

func (s *AccountService) GetBalance(symbol tradesatoshi.CurrencySymbol) (*tradesatoshi.GetBalanceResponse, error) {
	return nil, errors.New("Not implemented")
}

func (s *AccountService) GetBalances() (*tradesatoshi.GetBalancesResponse, error) {
	return nil, errors.New("Not implemented")
}

func (s *AccountService) GenerateAddress(symbol tradesatoshi.CurrencySymbol) (*tradesatoshi.GenerateAddressResponse, error) {
	return nil, errors.New("Not implemented")
}

func (s *AccountService) SubmitWithdraw(symbol tradesatoshi.CurrencySymbol, address string, amount decimal.Decimal, paymentId string) (*tradesatoshi.SubmitWithdrawResponse, error) {
	return nil, errors.New("Not implemented")
}

func (s *AccountService) GetDeposits(symbol tradesatoshi.CurrencySymbol, count int) (*tradesatoshi.GetDepositsResponse, error) {
	return nil, errors.New("Not implemented")
}

func (s *AccountService) GetWithdrawals(symbol tradesatoshi.CurrencySymbol, count int) (*tradesatoshi.GetWithdrawalsResponse, error) {
	return nil, errors.New("Not implemented")
}

func (s *AccountService) SubmitTransfer(symbol tradesatoshi.CurrencySymbol, username string, amount decimal.Decimal) (*tradesatoshi.SubmitTransferResponse, error) {
	return nil, errors.New("Not implemented")
}

func (s *AccountService) SubmitTip(symbol tradesatoshi.CurrencySymbol, numberActiveUsers int, amount decimal.Decimal, reason string) (*tradesatoshi.SubmitTipResponse, error) {
	return nil, errors.New("Not implemented")
}
