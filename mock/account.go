package mock

import (
	"github.com/shaunmza/tradesatoshi"
	"github.com/shopspring/decimal"
)

type AccountService struct {
	GetBalanceFn      func(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.GetBalanceResponse, error)
	GetBalanceInvoked bool

	GetBalancesFn      func() (*tradesatoshi.GetBalancesResponse, error)
	GetBalancesInvoked bool

	GenerateAddressFn      func(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.GenerateAddressResponse, error)
	GenerateAddressInvoked bool

	SubmitWithdrawFn      func(symbol tradesatoshi.MarketSymbol, address string, amount decimal.Decimal, paymentId string) (*tradesatoshi.SubmitWithdrawResponse, error)
	SubmitWithdrawInvoked bool

	GetDepositsFn      func(symbol tradesatoshi.MarketSymbol, count int) (*tradesatoshi.GetDepositsResponse, error)
	GetDepositsInvoked bool

	GetWithdrawalsFn      func(symbol tradesatoshi.MarketSymbol, count int) (*tradesatoshi.GetWithdrawalsResponse, error)
	GetWithdrawalsInvoked bool

	SubmitTransferFn      func(symbol tradesatoshi.MarketSymbol, username string, amount decimal.Decimal) (*tradesatoshi.SubmitTransferResponse, error)
	SubmitTransferInvoked bool

	SubmitTipFn      func(symbol tradesatoshi.MarketSymbol, numberActiveUsers int, amount decimal.Decimal, reason string) (*tradesatoshi.SubmitTipResponse, error)
	SubmitTipInvoked bool
}

func (s *AccountService) GetBalance(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.GetBalanceResponse, error) {
	s.GetBalanceInvoked = true
	return s.GetBalanceFn(symbol)
}

func (s *AccountService) GetBalances() (*tradesatoshi.GetBalancesResponse, error) {
	s.GetBalancesInvoked = true
	return s.GetBalancesFn()
}

func (s *AccountService) GenerateAddress(symbol tradesatoshi.MarketSymbol) (*tradesatoshi.GenerateAddressResponse, error) {
	s.GenerateAddressInvoked = true
	return s.GenerateAddressFn(symbol)
}

func (s *AccountService) SubmitWithdraw(symbol tradesatoshi.MarketSymbol, address string, amount decimal.Decimal, paymentId string) (*tradesatoshi.SubmitWithdrawResponse, error) {
	s.SubmitWithdrawInvoked = true
	return s.SubmitWithdrawFn(symbol, address, amount, paymentId)
}

func (s *AccountService) GetDeposits(symbol tradesatoshi.MarketSymbol, count int) (*tradesatoshi.GetDepositsResponse, error) {
	s.GetDepositsInvoked = true
	return s.GetDepositsFn(symbol, count)
}

func (s *AccountService) GetWithdrawals(symbol tradesatoshi.MarketSymbol, count int) (*tradesatoshi.GetWithdrawalsResponse, error) {
	s.GetWithdrawalsInvoked = true
	return s.GetWithdrawalsFn(symbol, count)
}

func (s *AccountService) SubmitTransfer(symbol tradesatoshi.MarketSymbol, username string, amount decimal.Decimal) (*tradesatoshi.SubmitTransferResponse, error) {
	s.SubmitTransferInvoked = true
	return s.SubmitTransferFn(symbol, username, amount)
}

func (s *AccountService) SubmitTip(symbol tradesatoshi.MarketSymbol, numberActiveUsers int, amount decimal.Decimal, reason string) (*tradesatoshi.SubmitTipResponse, error) {
	s.SubmitTipInvoked = true
	return s.SubmitTipFn(symbol, numberActiveUsers, amount, reason)
}
