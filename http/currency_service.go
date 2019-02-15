package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shaunmza/tradesatoshi"
	"net/http"
	"net/url"
)

type CurrencyClient struct {
	CurrencyService CurrencyService
}

func NewCurrencyClient() *CurrencyClient {
	c := &CurrencyClient{}

	return c
}

var _ tradesatoshi.CurrencyService = &CurrencyService{}

type CurrencyService struct {
	CurrencyService *tradesatoshi.CurrencyService
	URL             url.URL
}

func (s *CurrencyService) GetCurrency(symbol tradesatoshi.CurrencySymbol) (*tradesatoshi.Currency, error) {
	res, err := http.Get("https://tradesatoshi.com/api/public/GetCurrency?Symbol=" + string(symbol))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Not successful, return now.
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Invalid server response.")
	}

	var respBody *tradesatoshi.CurrencyResult
	if err := json.NewDecoder(res.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	if respBody.Success != true {
		return nil, errors.New(fmt.Sprintf("Request failed. Message: %s", respBody.Message))
	}

	emptyCurrency := tradesatoshi.Currency{}
	if respBody.Result == emptyCurrency {
		return nil, errors.New("Request failed. No result received")
	}

	return &respBody.Result, err

}

func (s *CurrencyService) GetCurrencies() (*[]tradesatoshi.Currency, error) {
	res, err := http.Get("https://tradesatoshi.com/api/public/GetCurrencies")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Not successful, return now.
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Invalid server response.")
	}

	var respBody *tradesatoshi.CurrenciesResult
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
