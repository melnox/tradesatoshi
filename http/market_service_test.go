package http_test

import (
	"github.com/jarcoal/httpmock"
	"github.com/shaunmza/tradesatoshi"
	"github.com/shaunmza/tradesatoshi/http"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
	"time"
)

func TestMarketService(t *testing.T) {
	t.Run("TickerOK", testMarketService_GetTicker_Success)
	t.Run("TickerServerError", testMarketService_GetTicker_ServerError)
	t.Run("TickerError", testMarketService_GetTicker_Error)
	t.Run("TickerEmptyResult", testMarketService_GetTicker_EmptyResult)
	t.Run("MarketStatusOK", testMarketService_GetMarketStatus_Success)
	t.Run("MarketStatusServerError", testMarketService_GetMarketStatus_ServerError)
	t.Run("MarketStatusError", testMarketService_GetMarketStatus_Error)
	t.Run("MarketStatusEmptyResult", testMarketService_GetMarketStatus_EmptyResult)
	t.Run("MarketHistoryOK", testMarketService_GetMarketHistory_Success)
	t.Run("MarketHistoryServerError", testMarketService_GetMarketHistory_ServerError)
	t.Run("MarketHistoryError", testMarketService_GetMarketHistory_Error)
	t.Run("MarketStatusEmptyResult", testMarketService_GetMarketHistory_EmptyResult)
	t.Run("MarketSummaryOK", testMarketService_GetMarketSummary_Success)
	t.Run("MarketSummaryServerError", testMarketService_GetMarketSummary_ServerError)
	t.Run("MarketSummaryError", testMarketService_GetMarketSummary_Error)
	t.Run("MarketSummaryEmptyResult", testMarketService_GetMarketSummary_EmptyResult)
	t.Run("MarketSummariesOK", testMarketService_GetMarketSummaries_Success)
	t.Run("MarketSummariesServerError", testMarketService_GetMarketSummaries_ServerError)
	t.Run("MarketSummariesError", testMarketService_GetMarketSummaries_Error)
	t.Run("MarketSummariesEmptyResult", testMarketService_GetMarketSummaries_EmptyResult)
	t.Run("OrderBookOK", testMarketService_GetOrderBook_Success)
	t.Run("OrderBookServerError", testMarketService_GetOrderBook_ServerError)
	t.Run("OrderBookError", testMarketService_GetOrderBook_Error)
	t.Run("OrderBookEmptyResult", testMarketService_GetOrderBook_EmptyResult)
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

	d, err := c.MarketService.GetTicker("LTC", "BTC")
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

	d, err := c.MarketService.GetTicker("LTC", "BTC")
	if err == nil {
		t.Fatal("did not fail")
	} else if err.Error() != "Invalid server response." {
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

	d, err := c.MarketService.GetTicker("LTC", "BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetTicker_EmptyResult(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getticker?market=LTC_BTC",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": {}
                    }`),
	)

	d, err := c.MarketService.GetTicker("LTC", "BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetMarketStatus_Success(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/GetMarketStatus?market=LTC_BTC",
		httpmock.NewStringResponder(200, `
                    {
                      "success": true,
                      "message": null,
                      "result": {
                        "marketStatus": "OK",
                        "statusMessage": null
                      }
                    }`),
	)

	d, err := c.MarketService.GetMarketStatus("LTC", "BTC")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &tradesatoshi.MarketStatus{MarketStatus: "OK", StatusMessage: ""}) {
		t.Fatalf("unexpected result: %#v", d)
	}
}

func testMarketService_GetMarketStatus_ServerError(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/GetMarketStatus?market=LTC_BTC",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.MarketService.GetMarketStatus("LTC", "BTC")
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetMarketStatus_Error(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/GetMarketStatus?market=LTC_BTC",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message"
                    }`),
	)

	d, err := c.MarketService.GetMarketStatus("LTC", "BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetMarketStatus_EmptyResult(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/GetMarketStatus?market=LTC_BTC",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": {}
                    }`),
	)

	d, err := c.MarketService.GetMarketStatus("LTC", "BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetMarketHistory_Success(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getmarkethistory?market=LTC_BTC&count=20",
		httpmock.NewStringResponder(200, `
                {
                  "success": true,
                  "message": null,
                  "result": [
                    {
                      "id": 512,
                      "timeStamp": "2016-04-28T01:34:02.397",
                      "quantity": 0.00784797,
                      "price": 0.01000000,
                      "orderType": "Buy",
                      "total": 0.00007848
                    },
                    {
                      "id": 503,
                      "timeStamp": "2016-04-23T08:16:38.087",
                      "quantity": 0.00134797,
                      "price": 0.08555000,
                      "orderType": "Buy",
                      "total": 0.00011532
                    },
                    {
                      "id": 502,
                      "timeStamp": "2016-04-23T08:16:34.91",
                      "quantity": 0.00650000,
                      "price": 0.07900000,
                      "orderType": "Buy",
                      "total": 0.00051350
                    }
                  ]
                } `),
	)

	d, err := c.MarketService.GetMarketHistory("LTC", "BTC", 20)

	if err != nil {
		t.Fatal(err)
	}
	expectedResult := make([]tradesatoshi.Order, 0)

	t1, err := time.Parse(tradesatoshi.TIME_FORMAT, "2016-04-28T01:34:02.397")
	if err != nil {
		t.Fatal(err)
	}

	t2, err := time.Parse(tradesatoshi.TIME_FORMAT, "2016-04-23T08:16:38.087")
	if err != nil {
		t.Fatal(err)
	}

	t3, err := time.Parse(tradesatoshi.TIME_FORMAT, "2016-04-23T08:16:34.91")
	if err != nil {
		t.Fatal(err)
	}

	expectedResult = append(expectedResult, tradesatoshi.Order{ID: 512,
		TimeStamp: tradesatoshi.Time{t1},
		Quantity:  decimal.NewFromFloat(0.00784797),
		Price:     decimal.NewFromFloat(0.01),
		OrderType: "Buy",
		Total:     decimal.NewFromFloat(0.00007848)})
	expectedResult = append(expectedResult, tradesatoshi.Order{ID: 503,
		TimeStamp: tradesatoshi.Time{t2},
		Quantity:  decimal.NewFromFloat(0.00134797),
		Price:     decimal.NewFromFloat(0.08555000),
		OrderType: "Buy",
		Total:     decimal.NewFromFloat(0.00011532)})
	expectedResult = append(expectedResult, tradesatoshi.Order{ID: 502,
		TimeStamp: tradesatoshi.Time{t3},
		Quantity:  decimal.NewFromFloat(0.0065),
		Price:     decimal.NewFromFloat(0.079),
		OrderType: "Buy",
		Total:     decimal.NewFromFloat(0.00051350)})
	if !reflect.DeepEqual(d, &expectedResult) {
		t.Fatalf("unexpected result: %+v", d)
	}
}

func testMarketService_GetMarketHistory_ServerError(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getmarkethistory?market=LTC_BTC&count=20",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.MarketService.GetMarketHistory("LTC", "BTC", 20)
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetMarketHistory_Error(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getmarkethistory?market=LTC_BTC&count=20",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": []
                    }`),
	)

	d, err := c.MarketService.GetMarketHistory("LTC", "BTC", 20)
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetMarketHistory_EmptyResult(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getmarkethistory?market=LTC_BTC&count=20",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": []
                    }`),
	)

	d, err := c.MarketService.GetMarketHistory("LTC", "BTC", 20)
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetMarketSummary_Success(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getmarketsummary?market=LTC_BTC",
		httpmock.NewStringResponder(200, `
                {
                    "success":true,
                    "message":null,
                    "result":{
                    "market":"LTC_BTC",
                    "high":0.01000000,
                    "low":0.01000000,
                    "volume":0.00784797,
                    "baseVolume":0.00007848,
                    "last":0.01000000,
                    "bid":0.01500000,
                    "ask":100.00000000,
                    "openBuyOrders":2,
                    "openSellOrders":7
                    }
                }`),
	)

	d, err := c.MarketService.GetMarketSummary("LTC", "BTC")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &tradesatoshi.Market{Market: "LTC_BTC", High: decimal.NewFromFloat(0.01),
		Low: decimal.NewFromFloat(0.01), Volume: decimal.NewFromFloat(0.00784797),
		BaseVolume: decimal.NewFromFloat(0.00007848), Last: decimal.NewFromFloat(0.01),
		Bid: decimal.NewFromFloat(0.015), Ask: decimal.NewFromFloat(100), OpenBuyOrders: 2, OpenSellOrders: 7}) {
		t.Fatalf("unexpected result: %#v", d)
	}
}

func testMarketService_GetMarketSummary_ServerError(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getmarketsummary?market=LTC_BTC",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.MarketService.GetMarketSummary("LTC", "BTC")
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetMarketSummary_Error(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getmarketsummary?market=LTC_BTC",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": []
                    }`),
	)

	d, err := c.MarketService.GetMarketSummary("LTC", "BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetMarketSummary_EmptyResult(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getmarketsummary?market=LTC_BTC",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": []
                    }`),
	)

	d, err := c.MarketService.GetMarketSummary("LTC", "BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetMarketSummaries_Success(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getmarketsummaries",
		httpmock.NewStringResponder(200, `
                {
                    "success":true,
                    "message":null,
                    "result":[
                    {
                        "market":"TRAD_BTC",
                        "high":0.00000000,
                        "low":0.00000000,
                        "volume":0.00000000,
                        "baseVolume":0.00000000,
                        "last":0.00000000,
                        "bid":0.00000000,
                        "ask":0.00000000,
                        "openBuyOrders":0,
                        "openSellOrders":0,
                        "marketStatus":"OK",
                        "change":21.20
                    },
                    {
                        "market":"BUMBA_BTC",
                        "high":0.00000000,
                        "low":0.00000000,
                        "volume":0.00000000,
                        "baseVolume":0.00000000,
                        "last":0.00000000,
                        "bid":0.00000000,
                        "ask":0.00000000,
                        "openBuyOrders":0,
                        "openSellOrders":0,
                        "marketStatus":"Paused",
                        "change":125.20
                    }]
                } `),
	)

	d, err := c.MarketService.GetMarketSummaries()
	expectedResult := make([]tradesatoshi.Market, 0)

	expectedResult = append(expectedResult, tradesatoshi.Market{Market: "TRAD_BTC", High: decimal.NewFromFloat(0),
		Low: decimal.NewFromFloat(0), Volume: decimal.NewFromFloat(0),
		BaseVolume: decimal.NewFromFloat(0), Last: decimal.NewFromFloat(0),
		Bid: decimal.NewFromFloat(0), Ask: decimal.NewFromFloat(0), OpenBuyOrders: 0, OpenSellOrders: 0,
		MarketStatus: "OK", Change: decimal.NewFromFloat(21.2)})
	expectedResult = append(expectedResult, tradesatoshi.Market{Market: "BUMBA_BTC", High: decimal.NewFromFloat(0),
		Low: decimal.NewFromFloat(0), Volume: decimal.NewFromFloat(0),
		BaseVolume: decimal.NewFromFloat(0), Last: decimal.NewFromFloat(0),
		Bid: decimal.NewFromFloat(0), Ask: decimal.NewFromFloat(0), OpenBuyOrders: 0, OpenSellOrders: 0,
		MarketStatus: "Paused", Change: decimal.NewFromFloat(125.2)})

	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &expectedResult) {
		t.Fatalf("unexpected result: %+v", d)
	}
}

func testMarketService_GetMarketSummaries_ServerError(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getmarketsummaries",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.MarketService.GetMarketSummaries()
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetMarketSummaries_Error(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getmarketsummaries",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": []
                    }`),
	)

	d, err := c.MarketService.GetMarketSummaries()
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetMarketSummaries_EmptyResult(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getmarketsummaries",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": []
                    }`),
	)

	d, err := c.MarketService.GetMarketSummaries()
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetOrderBook_Success(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getorderbook?market=LTC_BTC&type=both&depth=20",
		httpmock.NewStringResponder(200, `
                {
                  "success": true,
                  "message": null,
                  "result": {
                    "buy": [
                      {
                        "quantity": 0.00000000,
                        "rate": 0.01500000
                      },
                      {
                        "quantity": 1.00000000,
                        "rate": 0.00750000
                      }
                    ],
                    "sell": [
                      {
                        "quantity": 0.00000000,
                        "rate": 0.00756150
                      },
                      {
                        "quantity": 0.00000000,
                        "rate": 0.00770000
                      },
                      {
                        "quantity": 0.00000000,
                        "rate": 0.01000000
                      }
                    ]
                  }
                } `),
	)

	d, err := c.MarketService.GetOrderBook("LTC", "BTC", "both", 20)
	expectedBuys := make([]tradesatoshi.OrderBookAggregate, 0)
	expectedSells := make([]tradesatoshi.OrderBookAggregate, 0)

	expectedBuys = append(expectedBuys, tradesatoshi.OrderBookAggregate{Quantity: decimal.NewFromFloat(0), Rate: decimal.NewFromFloat(0.015)})
	expectedBuys = append(expectedBuys, tradesatoshi.OrderBookAggregate{Quantity: decimal.NewFromFloat(1), Rate: decimal.NewFromFloat(0.0075)})

	expectedSells = append(expectedSells, tradesatoshi.OrderBookAggregate{Quantity: decimal.NewFromFloat(0), Rate: decimal.NewFromFloat(0.00756150)})
	expectedSells = append(expectedSells, tradesatoshi.OrderBookAggregate{Quantity: decimal.NewFromFloat(0), Rate: decimal.NewFromFloat(0.0077)})
	expectedSells = append(expectedSells, tradesatoshi.OrderBookAggregate{Quantity: decimal.NewFromFloat(0), Rate: decimal.NewFromFloat(0.01)})

	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &tradesatoshi.OrderBookType{Buy: expectedBuys, Sell: expectedSells}) {
		t.Fatalf("unexpected result: %+v", d)
	}
}

func testMarketService_GetOrderBook_ServerError(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getorderbook?market=LTC_BTC&type=both&depth=20",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.MarketService.GetOrderBook("LTC", "BTC", "both", 20)
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetOrderBook_Error(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getorderbook?market=LTC_BTC&type=both&depth=20",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": []
                    }`),
	)

	d, err := c.MarketService.GetOrderBook("LTC", "BTC", "both", 20)
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testMarketService_GetOrderBook_EmptyResult(t *testing.T) {
	c := http.NewMarketClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://tradesatoshi.com/api/public/getorderbook?market=LTC_BTC&type=both&depth=20",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": []
                    }`),
	)

	d, err := c.MarketService.GetOrderBook("LTC", "BTC", "both", 20)
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}
