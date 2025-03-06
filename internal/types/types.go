package types

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
	Index  int64     `json:"index"`
	Value  float64 `json:"value"`
}

// Broker struct
type Broker struct {
	Name     string  `json:"name"`
	Age      int64     `json:"age"` // ✅ Changed to int for correct data type
	Gender   string  `json:"gender"`
	UserName string  `json:"user_name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Account  Account `json:"account"`
}

// Stock struct
type Stock struct {
	ID          int64   `json:"id"`
	Symbol      string  `json:"symbol"`
	CompanyName string  `json:"company_name"`
	LatestPrice float64 `json:"latest_price"`
}

// Holding struct
type Holding struct {
	ID       int64   `json:"id"`
	BrokerID int64   `json:"broker_id"`
	StockID  int64   `json:"stock_id"`
	Quantity int     `json:"quantity"`
	AvgPrice float64 `json:"avg_price"`
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
