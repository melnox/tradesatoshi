package tradesatoshi

// Client creates a connection to the services.
type Client interface {
	CurrencyService() CurrencyService
	MarketService() MarketService
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
