package http_test

import (
	"github.com/jarcoal/httpmock"
	"github.com/shaunmza/tradesatoshi"
	"github.com/shaunmza/tradesatoshi/http"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

func TestMarketService(t *testing.T) {
	t.Run("TickerOK", testMarketService_GetTicker_Success)
	t.Run("TickerServerError", testMarketService_GetTicker_ServerError)
	t.Run("TickerError", testMarketService_GetTicker_Error)
	/*t.Run("CurrencyEmptyResult", testCurrencyService_GetCurrency_EmptyResult)
	t.Run("CurrenciesOK", testCurrencyService_GetCurrencies_Success)
	t.Run("CurrenciesServerError", testCurrencyService_GetCurrencies_ServerError)
	t.Run("CurrenciesError", testCurrencyService_GetCurrencies_Error)
	t.Run("CurrenciesEmptyResult", testCurrencyService_GetCurrencies_EmptyResult)*/
}

func testMarketService_GetTicker_Success(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getticker?market=LTC_BTC",
		httpmock.NewStringResponder(200, `
                {
                    "success":true,
                    "message":null,
                    "result":{
                    "bid":0.01500000,
                    "ask":100.00000000,
                    "last":0.01000000,
                    "market":"LTC_BTC"
                    }
                }`),
	)

	d, err := c.MarketService.GetTicker("LTC_BTC")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &tradesatoshi.Market{Bid: decimal.NewFromFloat(0.015), Ask: decimal.NewFromFloat(100), Last: decimal.NewFromFloat(0.01), Market: "LTC_BTC"}) {
		t.Fatalf("unexpected result: %#v", d)
	}
}

func testMarketService_GetTicker_ServerError(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getticker?market=LTC_BTC",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.MarketService.GetTicker("LTC_BTC")
	if err == nil {
		t.Fatal("did not fail")
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetTicker_Error(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getticker?market=LTC_BTC",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": {}
                    }`),
	)

	d, err := c.MarketService.GetTicker("LTC_BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

/*func testCurrencyService_GetCurrency_EmptyResult(t *testing.T) {
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
}*/
