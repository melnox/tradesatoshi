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

func TestAccountService(t *testing.T) {
	t.Run("BalanceOK", testAccountService_GetBalance_Success)
	t.Run("BalanceServerError", testAccountService_GetBalance_ServerError)
	t.Run("BalanceError", testAccountService_GetBalance_Error)
	t.Run("BalanceEmptyResult", testAccountService_GetBalance_EmptyResult)
	t.Run("BalancesOK", testAccountService_GetBalances_Success)
	t.Run("BalancesServerError", testAccountService_GetBalances_ServerError)
	t.Run("BalancesError", testAccountService_GetBalances_Error)
	t.Run("BalancesEmptyResult", testAccountService_GetBalances_EmptyResult)
	t.Run("GenerateAddressOK", testAccountService_GenerateAddress_Success)
	t.Run("GenerateAddressServerError", testAccountService_GenerateAddress_ServerError)
	t.Run("GenerateAddressError", testAccountService_GenerateAddress_Error)
	t.Run("GenerateAddressEmptyResult", testAccountService_GenerateAddress_EmptyResult)
	t.Run("SubmitWithdrawOK", testAccountService_SubmitWithdraw_Success)
	t.Run("SubmitWithdrawServerError", testAccountService_SubmitWithdraw_ServerError)
	t.Run("SubmitWithdrawError", testAccountService_SubmitWithdraw_Error)
	t.Run("SubmitWithdrawEmptyResult", testAccountService_SubmitWithdraw_EmptyResult)
	t.Run("DepositsOK", testAccountService_GetDeposits_Success)
	t.Run("DepositsServerError", testAccountService_GetDeposits_ServerError)
	t.Run("DepositsError", testAccountService_GetDeposits_Error)
	t.Run("DepositsEmptyResult", testAccountService_GetDeposits_EmptyResult)
	t.Run("WithdrawalsOK", testAccountService_GetWithdrawals_Success)
	t.Run("WithdrawalsServerError", testAccountService_GetWithdrawals_ServerError)
	t.Run("WithdrawalsError", testAccountService_GetWithdrawals_Error)
	t.Run("WithdrawalsEmptyResult", testAccountService_GetWithdrawals_EmptyResult)
	t.Run("SubmitTransferOK", testAccountService_SubmitTransfer_Success)
	t.Run("SubmitTransferServerError", testAccountService_SubmitTransfer_ServerError)
	t.Run("SubmitTransferError", testAccountService_SubmitTransfer_Error)
	t.Run("SubmitTransferEmptyResult", testAccountService_SubmitTransfer_EmptyResult)
	t.Run("SubmitTipOK", testAccountService_SubmitTip_Success)
	t.Run("SubmitTipServerError", testAccountService_SubmitTip_ServerError)
	t.Run("SubmitTipError", testAccountService_SubmitTip_Error)
	t.Run("SubmitTipEmptyResult", testAccountService_SubmitTip_EmptyResult)
}

func testAccountService_GetBalance_Success(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getbalance",
		httpmock.NewStringResponder(200, `
                   {  
					   "success":true,
					   "message":null,
					   "result": {  
							 "Currency": "BTC",
							 "CurrencyLong": "Bitcoin",
							 "Avaliable": 49.00000000,
							 "Total": 53.00000000,
							 "HeldForTrades": 1.00000000,
							 "Unconfirmed": 1.00000000,
							 "PendingWithdraw": 1.00000000,
							 "Address": "1HB5XMLmzFVj8ALj6mfBsbifRoD4miY36v"
						  }
					}`),
	)

	d, err := c.AccountService.GetBalance("BTC")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &tradesatoshi.CurrencyBalance{Currency: "BTC", CurrencyLong: "Bitcoin",
		Avaliable: decimal.NewFromFloat(49),  Total:decimal.NewFromFloat(53),
		HeldForTrades:decimal.NewFromFloat(1), Unconfirmed: decimal.NewFromFloat(1),
		PendingWithdraw:decimal.NewFromFloat(1), Address:"1HB5XMLmzFVj8ALj6mfBsbifRoD4miY36v"} ) {
		t.Fatalf("unexpected result: %#v", d)
	}
}

func testAccountService_GetBalance_ServerError(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getbalance",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.AccountService.GetBalance("BTC")
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_GetBalance_Error(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getbalance",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": {}
                    }`),
	)

	_, err := c.AccountService.GetBalance("BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}

func testAccountService_GetBalance_EmptyResult(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getbalance",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": {}
                    }`),
	)

	d, err := c.AccountService.GetBalance("BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_GetBalances_Success(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getbalances",
		httpmock.NewStringResponder(200, `
                   {  
					   "success":true,
					   "message":null,
					   "result":[{  
							 "Currency": "BTC",
							 "CurrencyLong": "Bitcoin",
							 "Avaliable": 49.00000000,
							 "Total": 53.00000000,
							 "HeldForTrades": 1.00000000,
							 "Unconfirmed": 1.00000000,
							 "PendingWithdraw": 1.00000000,
							 "Address": "1HB5XMLmzFVj8ALj6mfBsbifRoD4miY36v"
						  },{  
							 "Currency": "LTC",
							 "CurrencyLong": "Litecoin",
							 "Avaliable": 49.00000000,
							 "Total": 53.00000000,
							 "HeldForTrades": 1.00000000,
							 "Unconfirmed": 1.00000000,
							 "PendingWithdraw": 1.00000000,
							 "Address": "3KBUuGko4H5ke7EVsq9B7PLK1c5Askdd7y"
						  }
					   ]
					}`),
	)

	d, err := c.AccountService.GetBalances()
	expectedResult := make([]tradesatoshi.CurrencyBalance, 0)
	expectedResult = append(expectedResult, tradesatoshi.CurrencyBalance{Currency: "BTC", CurrencyLong: "Bitcoin",
		Avaliable: decimal.NewFromFloat(49),  Total:decimal.NewFromFloat(53),
		HeldForTrades:decimal.NewFromFloat(1), Unconfirmed: decimal.NewFromFloat(1),
		PendingWithdraw:decimal.NewFromFloat(1), Address:"1HB5XMLmzFVj8ALj6mfBsbifRoD4miY36v"})
	expectedResult = append(expectedResult, tradesatoshi.CurrencyBalance{Currency: "LTC", CurrencyLong: "Litecoin",
		Avaliable: decimal.NewFromFloat(49),  Total:decimal.NewFromFloat(53),
		HeldForTrades:decimal.NewFromFloat(1), Unconfirmed: decimal.NewFromFloat(1),
		PendingWithdraw:decimal.NewFromFloat(1), Address:"3KBUuGko4H5ke7EVsq9B7PLK1c5Askdd7y"})
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &expectedResult) {
		t.Fatalf("unexpected result: %+v", d)
	}
}

func testAccountService_GetBalances_ServerError(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getbalances",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.AccountService.GetBalances()
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_GetBalances_Error(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getbalances",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": []
                    }`),
	)

	_, err := c.AccountService.GetBalances()
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}

func testAccountService_GetBalances_EmptyResult(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getbalances",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": []
                    }`),
	)

	d, err := c.AccountService.GetBalances()
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_GenerateAddress_Success(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/generateaddress",
		httpmock.NewStringResponder(200, `
                   {  
					   "success":true,
					   "message":null,
					   "result": {  
							 "Currency": "LTC",
							 "Address": "3KBUuGko4H5ke7EVsq9B7PLK1c5Askdd7y",
							 "PaymentId":"8e93b95650b473c859693771864432dbd10a192629d4883aa0c89857b1cd67fd"
						  }
					}`),
	)

	d, err := c.AccountService.GenerateAddress("BTC")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &tradesatoshi.AddressResult{Currency: "LTC",
	Address:"1HB5XMLmzFVj8ALj6mfBsbifRoD4miY36v",
	PaymentId:"8e93b95650b473c859693771864432dbd10a192629d4883aa0c89857b1cd67fd"} ) {
		t.Fatalf("unexpected result: %#v", d)
	}
}

func testAccountService_GenerateAddress_ServerError(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/generateaddress",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.AccountService.GenerateAddress("BTC")
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_GenerateAddress_Error(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/generateaddress",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": {}
                    }`),
	)

	_, err := c.AccountService.GenerateAddress("BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}

func testAccountService_GenerateAddress_EmptyResult(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/generateaddress",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": {}
                    }`),
	)

	d, err := c.AccountService.GenerateAddress("BTC")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_SubmitWithdraw_Success(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/submitwithdraw",
		httpmock.NewStringResponder(200, `
                   {  
					   "success":true,
					   "message":null,
					   "result": {
							 "WithdrawalId": 546474
					   }
					}`),
	)

	d, err := c.AccountService.SubmitWithdraw("BTC", "", decimal.NewFromFloat(1), "")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &tradesatoshi.WithdrawalResult{ID:546474} ) {
		t.Fatalf("unexpected result: %#v", d)
	}
}

func testAccountService_SubmitWithdraw_ServerError(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/submitwithdraw",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.AccountService.SubmitWithdraw("BTC", "", decimal.NewFromFloat(1), "")
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_SubmitWithdraw_Error(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/submitwithdraw",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": {}
                    }`),
	)

	_, err := c.AccountService.SubmitWithdraw("BTC", "", decimal.NewFromFloat(1), "")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}

func testAccountService_SubmitWithdraw_EmptyResult(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/submitwithdraw",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": {}
                    }`),
	)

	d, err := c.AccountService.SubmitWithdraw("BTC", "", decimal.NewFromFloat(1), "")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_GetDeposits_Success(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getdeposits",
		httpmock.NewStringResponder(200, `
                   {  
					   "success":true,
					   "message":null,
					   "result": [{  
							 "Id": "436436",
							 "Currency": "BTC",
							 "CurrencyLong": "Bitcoin",
							 "Amount": 100.00000000,
							 "Status": "Unconfirmed",
							 "Txid": "9281eacaad58335b884adc24be884c00200a4fc17b2e05c72e255976223de187",
							 "Confirmations": 0,
							 "Timestamp": "2015-12-07T20:04:05.3947572"
						  },{  
							 "Id": "436437",
							 "Currency": "BTC",
							 "CurrencyLong": "Bitcoin",
							 "Amount": 100.00000000,
							 "Status": "Confirmed",
							 "Txid": "6ddbaca454c97ba4e8a87a1cb49fa5ceace80b89eaced84b46a8f52c2b8c8ca3",
							 "Confirmations": 12,
							 "Timestamp": "2015-12-07T20:04:05.3947572"
						  }
					   ]
					}`),
	)

	d, err := c.AccountService.GetDeposits("all", 20)
	expectedResult := make([]tradesatoshi.DepositResult, 0)

	t1, err := time.Parse(tradesatoshi.TIME_FORMAT, "2015-12-07T20:04:05.3947572")
	if err != nil {
		t.Fatal(err)
	}

	t2, err := time.Parse(tradesatoshi.TIME_FORMAT, "2015-12-07T20:04:05.3947572")
	if err != nil {
		t.Fatal(err)
	}

	expectedResult = append(expectedResult, tradesatoshi.DepositResult{ID: 436436, Currency: "BTC", CurrencyLong: "Bitcoin",
		Amount:decimal.NewFromFloat(100), Status:"Unconfirmed", Txid:"9281eacaad58335b884adc24be884c00200a4fc17b2e05c72e255976223de187",
		Confirmations:0, Timestamp:tradesatoshi.Time{t1}})
	expectedResult = append(expectedResult, tradesatoshi.DepositResult{ID: 436437, Currency: "BTC", CurrencyLong: "Bitcoin",
		Amount:decimal.NewFromFloat(100), Status:"Confirmed", Txid:"6ddbaca454c97ba4e8a87a1cb49fa5ceace80b89eaced84b46a8f52c2b8c8ca3",
		Confirmations:12, Timestamp:tradesatoshi.Time{t2}})
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &expectedResult) {
		t.Fatalf("unexpected result: %+v", d)
	}
}

func testAccountService_GetDeposits_ServerError(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getdeposits",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.AccountService.GetBalances()
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_GetDeposits_Error(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getdeposits",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": []
                    }`),
	)

	_, err := c.AccountService.GetBalances()
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}

func testAccountService_GetDeposits_EmptyResult(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getdeposits",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": []
                    }`),
	)

	d, err := c.AccountService.GetBalances()
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_GetWithdrawals_Success(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getwithdrawals",
		httpmock.NewStringResponder(200, `
                   {  
					   "success":true,
					   "message":null,
					   "result": [{  
							 "Id": "436436",
							 "Currency": "BTC",
							 "CurrencyLong": "Bitcoin",
							 "Amount": 100.00000000,
							 "Fee": 0.00040000,
							 "Address": "3KBUuGko4H5ke7EVsq9B7PLK1c5Askdd7y"
							 "Status": "Unconfirmed",
							 "Txid": null,
							 "Confirmations": 0,
							 "Timestamp": "2015-12-07T20:04:05.3947572",
							 "IsApi" : false
						  },{  
							 "Id": "436437",
							 "Currency": "BTC",
							 "CurrencyLong": "Bitcoin",
							 "Amount": 100.00000000,
							 "Fee": 0.00040000,
							 "Address": "3KBUuGko4H5ke7EVsq9B7PLK1c5Askdd7y",
							 "Status": "Complete",
							 "Txid": "9281eacaad58335b884adc24be884c00200a4fc17b2e05c72e255976223de187",
							 "Confirmations": 12,
							 "Timestamp": "2015-12-07T20:04:05.3947572",
							 "IsApi" : false
						  }
					   ]
					}`),
	)

	d, err := c.AccountService.GetWithdrawals("all", 20)
	expectedResult := make([]tradesatoshi.WithdrawalResult, 0)

	t1, err := time.Parse(tradesatoshi.TIME_FORMAT, "2015-12-07T20:04:05.3947572")
	if err != nil {
		t.Fatal(err)
	}

	t2, err := time.Parse(tradesatoshi.TIME_FORMAT, "2015-12-07T20:04:05.3947572")
	if err != nil {
		t.Fatal(err)
	}

	expectedResult = append(expectedResult, tradesatoshi.WithdrawalResult{ID: 436436, Currency: "BTC", CurrencyLong: "Bitcoin",
		Amount:decimal.NewFromFloat(100), Fee:decimal.NewFromFloat(0.0004), Address:"3KBUuGko4H5ke7EVsq9B7PLK1c5Askdd7y",
		Status:"Unconfirmed", Txid:"", Confirmations:0, Timestamp:tradesatoshi.Time{t1}, IsApi:false})
	expectedResult = append(expectedResult, tradesatoshi.WithdrawalResult{ID: 436437, Currency: "BTC", CurrencyLong: "Bitcoin",
		Amount:decimal.NewFromFloat(100), Fee:decimal.NewFromFloat(0.0004), Address:"3KBUuGko4H5ke7EVsq9B7PLK1c5Askdd7y",
		Status:"Complete", Txid:"9281eacaad58335b884adc24be884c00200a4fc17b2e05c72e255976223de187", Confirmations:0,
		Timestamp:tradesatoshi.Time{t2}, IsApi:false})
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &expectedResult) {
		t.Fatalf("unexpected result: %+v", d)
	}
}

func testAccountService_GetWithdrawals_ServerError(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getwithdrawals",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.AccountService.GetBalances()
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_GetWithdrawals_Error(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getwithdrawals",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": []
                    }`),
	)

	_, err := c.AccountService.GetBalances()
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}

func testAccountService_GetWithdrawals_EmptyResult(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/getwithdrawals",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": []
                    }`),
	)

	d, err := c.AccountService.GetBalances()
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_SubmitTransfer_Success(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/submittransfer",
		httpmock.NewStringResponder(200, `
                   {
                    "success": true,
                    "message": null,
                    "result": {
                            "data": "Successfully transfered 251.00000000 DOGE to XYZ"
                        }
                }`),
	)

	d, err := c.AccountService.SubmitTransfer("DOGE", "XYZ", decimal.NewFromFloat(251))
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &tradesatoshi.SubmitTransferResult{Data:"Successfully transfered 251.00000000 DOGE to XYZ"} ) {
		t.Fatalf("unexpected result: %#v", d)
	}
}

func testAccountService_SubmitTransfer_ServerError(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/submittransfer",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.AccountService.SubmitTransfer("DOGE", "XYZ", decimal.NewFromFloat(251))
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_SubmitTransfer_Error(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/submittransfer",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": {}
                    }`),
	)

	_, err := c.AccountService.SubmitTransfer("DOGE", "XYZ", decimal.NewFromFloat(251))
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}

func testAccountService_SubmitTransfer_EmptyResult(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/submittransfer",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": {}
                    }`),
	)

	d, err := c.AccountService.SubmitTransfer("DOGE", "XYZ", decimal.NewFromFloat(251))
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_SubmitTip_Success(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/SubmitTip",
		httpmock.NewStringResponder(200, `
                   {
                    "success": true,
                    "message": null,
                    "result": {
                            "data": "TipBot: xyz tipped the last 10 users 100.00000000 BOLI each., Reason: Happy trading...!!!"
                        }
                }`),
	)

	d, err := c.AccountService.SubmitTip("BOLI", 10, decimal.NewFromFloat(100), "Happy trading...!!!")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(d, &tradesatoshi.CurrencyBalance{Currency: "BTC", CurrencyLong: "Bitcoin",
		Avaliable: decimal.NewFromFloat(49),  Total:decimal.NewFromFloat(53),
		HeldForTrades:decimal.NewFromFloat(1), Unconfirmed: decimal.NewFromFloat(1),
		PendingWithdraw:decimal.NewFromFloat(1), Address:"1HB5XMLmzFVj8ALj6mfBsbifRoD4miY36v"} ) {
		t.Fatalf("unexpected result: %#v", d)
	}
}

func testAccountService_SubmitTip_ServerError(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/SubmitTip",
		httpmock.NewStringResponder(500, ``),
	)

	d, err := c.AccountService.SubmitTip("BOLI", 10, decimal.NewFromFloat(100), "Happy trading...!!!")
	if err == nil {
		t.Fatal("did not fail")
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}

func testAccountService_SubmitTip_Error(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/SubmitTip",
		httpmock.NewStringResponder(200, `
                   {
                      "success": false,
                      "message": "Some message",
                      "result": {}
                    }`),
	)

	_, err := c.AccountService.SubmitTip("BOLI", 10, decimal.NewFromFloat(100), "Happy trading...!!!")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. Message: Some message" {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}

func testAccountService_SubmitTip_EmptyResult(t *testing.T) {
	c := http.NewAccountClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://tradesatoshi.com/api/private/SubmitTip",
		httpmock.NewStringResponder(200, `
                   {
                      "success": true,
                      "message": null,
                      "result": {}
                    }`),
	)

	d, err := c.AccountService.SubmitTip("BOLI", 10, decimal.NewFromFloat(100), "Happy trading...!!!")
	if err == nil {
		t.Fatal(err)
	} else if err.Error() != "Request failed. No result received" {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	if d != nil {
		t.Fatal("should not get a result")
	}
}
