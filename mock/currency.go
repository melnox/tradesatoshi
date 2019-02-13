package mock

import "github.com/shaunmza/tradesatoshi"

type CurrencyService struct {
	GetCurrencyFn      func(symbol tradesatoshi.CurrencySymbol) (*tradesatoshi.CurrencyResult, error)
	GetCurrencyInvoked bool

	GetCurrenciesFn      func() (*tradesatoshi.CurrenciesResult, error)
	GetCurrenciesInvoked bool
}

func (s *CurrencyService) GetCurrency(symbol tradesatoshi.CurrencySymbol) (*tradesatoshi.CurrencyResult, error) {
	s.GetCurrencyInvoked = true
	return s.GetCurrencyFn(symbol)
}

func (s *CurrencyService) GetCurrencies() (*tradesatoshi.CurrenciesResult, error) {
	s.GetCurrenciesInvoked = true
	return s.GetCurrenciesFn()
}
