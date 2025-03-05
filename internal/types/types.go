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

type Broker struct {
	name     string
	age      string
	gender   string
	userName string	
	email    string
	password string
	Account  Account
}

type Account struct {
	initialCredit float64
	Companies     []CompanyTab
	Stocks        map[CompanyTab]int
}

type Company struct {
	CompanyTab CompanyTab `json:"companyTab"`
}
type CompanyDetails struct {
	Title    string   `json:"title"`
	Market   string   `json:"market"`
	Stock    string   `json:"stock"`
	Index    int      `json:"index"`
	PriceBar PriceBar `json:"priceBar"`
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
