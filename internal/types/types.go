package types

import "time"

type App struct {
	ROOT                string
	HostURL             string
	CompaniesAPI        string
	PricesAPI           string
	HistoricalDirectory string
	CompaniesDirectory  string
	OutputFile          string
	DataBase            []byte
	Country             string
	LastPrices          map[string]float64
	DataDownloaded      bool
	AbuQuir             string
}

type Account struct {
	InitialCredit float64          `json:"initial_credit"`
	Companies     []CompanyDetails `json:"companies"`
	Stocks        []Stock          `json:"stocks"` // ✅ Plural for consistency
}

// CompanyDetails struct
type CompanyDetails struct {
	Title  string  `json:"title"`
	Market string  `json:"market"`
	Stock  string  `json:"stock"`
	Index  int64   `json:"index"`
	Value  float64 `json:"value"`
}

// Broker struct
type Broker struct {
	Name     string  `json:"name"`
	Age      int64   `json:"age"` // ✅ Changed to int for correct data type
	Gender   string  `json:"gender"`
	UserName string  `json:"user_name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Account  Account `json:"account"`
}

// Stock struct
type Stock struct {
	ID          int64         `json:"id"`
	Symbol      string        `json:"symbol"`       // stock
	CompanyName string        `json:"company_name"` // title
	Market      string        `json:"market"`       // egx
	Index       int           `json:"index"`        // don know
	PriceBar    StockPriceBar `json:"priceBar"`
}

type StockPriceBar struct {
	Value            string `json:"value"`
	Change           string `json:"change"`
	ChangePercentage string `json:"changePercentage"`
	Open             string `json:"open"`
	Close            string `json:"close"`
	High             string `json:"high"`
	Low              string `json:"low"`
	HistoricalHigh   string `json:"historicalHigh"`
	HistoricalLow    string `json:"historicalLow"`
	Volume           string `json:"volume"`
	Turnover         string `json:"turnover"`
	Status           string `json:"status"`
	UpdatedAt        string `json:"updatedAt"`
	Currency         string `json:"currency"`
}

// Holding struct
type Holding struct {
	BrokerID        int64     `json:"broker_id"`
	StockSymbol     string    `json:"stock_symbol"` // Replacing StockID with StockSymbol
	Quantity        int64     `json:"quantity"`
	BuyingPrice     float64   `json:"buying_price"`
	CurrentPrice    float64   `json:"current_price"`
	CompanyName     string    `json:"company_name"` // title
	TotalInvestment float64   `json:"total_investment"`
	CurrentValue    float64   `json:"current_value"`
	Profit          float64   `json:"profit"`
	ProfitPercent   float64   `json:"profit_percent"`
	LastUpdated     time.Time `json:"last_updated"`
}

// ------------------------------------------------------------------------------------------------------------------------------------
// for the api
type Company struct {
	CompanyTab CompanyTab `json:"companyTab"`
}

type CompanyTab struct {
	Title    string   `json:"title"`
	Market   string   `json:"market"`
	Stock    string   `json:"stock"`
	Index    int      `json:"index"`
	PriceBar PriceBar `json:"priceBar"`
}

type PriceBar struct {
	Value            string `json:"value"`
	Change           string `json:"change"`
	ChangePercentage string `json:"changePercentage"`
	Open             string `json:"open"`
	Close            string `json:"close"`
	High             string `json:"high"`
	Low              string `json:"low"`
	HistoricalHigh   string `json:"historicalHigh"`
	HistoricalLow    string `json:"historicalLow"`
	Volume           string `json:"volume"`
	Turnover         string `json:"turnover"`
	Status           string `json:"status"`
	UpdatedAt        string `json:"updatedAt"`
	Currency         string `json:"currency"`
}

type Order struct {
	OrderID   string  `json:"order_id"`
	UserID    string  `json:"user_id"`
	Stock     string  `json:"stock"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	OrderType string  `json:"order_type"` // "BUY" or "SELL"
}
