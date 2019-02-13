package tradesatoshi

import "github.com/shopspring/decimal"

// Client creates a connection to the services.
type Client interface {
	CurrencyService() CurrencyService
	MarketService() MarketService
	AccountService() AccountService
	OrderService() OrderService
}

type CurrencyService interface {
	GetCurrency(symbol CurrencySymbol) (*CurrencyResult, error)
	GetCurrencies() (*CurrenciesResult, error)
}

type MarketService interface {
	GetTicker(symbol MarketSymbol) (*TickerResult, error)
	GetMarketStatus(symbol MarketSymbol) (*MarketStatusResult, error)
	GetMarketHistory(symbol MarketSymbol, count int) (*MarketHistoryResult, error)
	GetMarketSummary(symbol MarketSymbol) (*MarketSummaryResult, error)
	GetMarketSummaries() (*MarketSummariesResult, error)
	GetOrderBook(symbol MarketSymbol, orderType string, count int) (*OrderBookResult, error)
}

type AccountService interface {
	GetBalance(symbol MarketSymbol) (*GetBalanceResponse, error)
	GetBalances() (*GetBalancesResponse, error)
	GenerateAddress(symbol MarketSymbol) (*GenerateAddressResponse, error)
	SubmitWithdraw(symbol MarketSymbol, address string, amount decimal.Decimal, paymentId string) (*SubmitWithdrawResponse, error)
	GetDeposits(symbol MarketSymbol, count int) (*GetDepositsResponse, error)
	GetWithdrawals(symbol MarketSymbol, count int) (*GetWithdrawalsResponse, error)
	SubmitTransfer(symbol MarketSymbol, username string, amount decimal.Decimal) (*SubmitTransferResponse, error)
	SubmitTip(symbol MarketSymbol, numberActiveUsers int, amount decimal.Decimal, reason string) (*SubmitTipResponse, error)
}

type OrderService interface {
	GetOrder(OrderID OrderID) (*OrderResponse, error)
	GetOrders(marketSymbol MarketSymbol, count int) (*OrdersResponse, error)
	SubmitOrder(marketSymbol MarketSymbol, orderType OrderType, amount decimal.Decimal, price decimal.Decimal) (*SubmitOrderResponse, error)
	CancelOrder(cancelOrderType CancelOrderType, orderID OrderID, marketSymbol MarketSymbol) (*SubmitOrderResponse, error)
	GetTradeHistory(marketSymbol MarketSymbol, count int, page int) (*TradeHistoryResponse, error)
}
