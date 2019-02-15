package http_test

import (
	"github.com/jarcoal/httpmock"
	"github.com/shaunmza/tradesatoshi"
	"github.com/shaunmza/tradesatoshi/http"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

func TestCurrencyService(t *testing.T) {
	t.Run("CurrencyOK", testCurrencyService_GetCurrency_Success)
	t.Run("CurrencyServerError", testCurrencyService_GetCurrency_ServerError)
	t.Run("CurrencyError", testCurrencyService_GetCurrency_Error)
	t.Run("CurrencyEmptyResult", testCurrencyService_GetCurrency_EmptyResult)
	t.Run("CurrenciesOK", testCurrencyService_GetCurrencies_Success)
	t.Run("CurrenciesServerError", testCurrencyService_GetCurrencies_ServerError)
	t.Run("CurrenciesError", testCurrencyService_GetCurrencies_Error)
	t.Run("CurrenciesEmptyResult", testCurrencyService_GetCurrencies_EmptyResult)
}

func testCurrencyService_GetCurrency_Success(t *testing.T) {
	c := http.NewCurrencyClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/GetCurrency?Symbol=BTC",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": {
                        "currency": "BTC",
                        "currencyLong": "Bitcoin",
                        "minConfirmation": 6,
                        "txFee": 0.00200000,
                        "status": "OK",
                        "statusMessage": "",
                        "minBaseTrade": 0.00010000,
                        "isTipEnabled": true,
                        "minTip": 0.00000100,
                        "maxTip": 0.10000000
                      }
                    }`),
	)

	d, err := c.CurrencyService.GetCurrency("BTC")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &tradesatoshi.Currency{Currency: "BTC", CurrencyLong: "Bitcoin", MinConfirmation: 6,
		TxFee: decimal.NewFromFloat(0.00200000), Status: "OK", StatusMessage: "",
		MinBaseTrade: decimal.NewFromFloat(0.00010000), IsTipEnabled: true,
		MinTip: decimal.NewFromFloat(0.00000100), MaxTip: decimal.NewFromFloat(0.10000000)}) {
		t.Fatalf("unexpected result: %#v", d)
	}
}

func testCurrencyService_GetCurrency_ServerError(t *testing.T) {
	c := http.NewCurrencyClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/GetCurrency?Symbol=BTC",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.CurrencyService.GetCurrency("BTC")
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testCurrencyService_GetCurrency_Error(t *testing.T) {
	c := http.NewCurrencyClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/GetCurrency?Symbol=BTC",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": {}
                    }`),
	)

	_, err := c.CurrencyService.GetCurrency("BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}

func testCurrencyService_GetCurrency_EmptyResult(t *testing.T) {
	c := http.NewCurrencyClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/GetCurrency?Symbol=BTC",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": {}
                    }`),
	)

	d, err := c.CurrencyService.GetCurrency("BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testCurrencyService_GetCurrencies_Success(t *testing.T) {
	c := http.NewCurrencyClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/GetCurrencies",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": [
                        {
                          "currency": "BOLI",
                          "currencyLong": "Bolivarcoin",
                          "minConfirmation": 3,
                          "txFee": 0.00000000,
                          "status": "OK"
                        },
                        {
                          "currency": "BTC",
                          "currencyLong": "Bitcoin",
                          "minConfirmation": 6,
                          "txFee": 0.00000000,
                          "status": "OK"
                        }
                      ]
                    }`),
	)

	d, err := c.CurrencyService.GetCurrencies()
	expectedResult := make([]tradesatoshi.Currency, 0)
	expectedResult = append(expectedResult, tradesatoshi.Currency{Currency: "BOLI",
		CurrencyLong:    "Bolivarcoin",
		MinConfirmation: 3,
		TxFee:           decimal.NewFromFloat(0.00000000),
		Status:          "OK"})
	expectedResult = append(expectedResult, tradesatoshi.Currency{Currency: "BTC",
		CurrencyLong:    "Bitcoin",
		MinConfirmation: 6,
		TxFee:           decimal.NewFromFloat(0.00000000),
		Status:          "OK"})
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &expectedResult) {
		t.Fatalf("unexpected result: %+v", d)
	}
}

func testCurrencyService_GetCurrencies_ServerError(t *testing.T) {
	c := http.NewCurrencyClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/GetCurrencies",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.CurrencyService.GetCurrencies()
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testCurrencyService_GetCurrencies_Error(t *testing.T) {
	c := http.NewCurrencyClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/GetCurrencies",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": []
                    }`),
	)

	_, err := c.CurrencyService.GetCurrencies()
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}

func testCurrencyService_GetCurrencies_EmptyResult(t *testing.T) {
	c := http.NewCurrencyClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/GetCurrencies",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": []
                    }`),
	)

	d, err := c.CurrencyService.GetCurrencies()
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}
