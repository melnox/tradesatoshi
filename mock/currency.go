package mock

import "github.com/shaunmza/tradesatoshi"

type CurrencyService struct {
	GetCurrencyFn      func(symbol tradesatoshi.CurrencySymbol) (*tradesatoshi.Currency, error)
	GetCurrencyInvoked bool

	GetCurrenciesFn      func() (*[]tradesatoshi.Currency, error)
	GetCurrenciesInvoked bool
}

func (s *CurrencyService) GetCurrency(symbol tradesatoshi.CurrencySymbol) (*tradesatoshi.Currency, error) {
	s.GetCurrencyInvoked = true
	return s.GetCurrencyFn(symbol)
}

func (s *CurrencyService) GetCurrencies() (*[]tradesatoshi.Currency, error) {
	s.GetCurrenciesInvoked = true
	return s.GetCurrenciesFn()
}
