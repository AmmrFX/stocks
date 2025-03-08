package datastore_repo

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud.google.com/go/datastore"
)

// Account struct
type Account struct {
	InitialCredit float64          `datastore:"initialCredit,noindex"` // No need to index balances
	Companies     []CompanyDetails `datastore:"companies"`
	Stocks        []Stock          `datastore:"stocks"`
}

// CompanyDetails struct
type CompanyDetails struct {
	Title  string  `datastore:"title"`
	Market string  `datastore:"market"`
	Stock  string  `datastore:"stock"`
	Index  int64   `datastore:"index"`
	Value  float64 `datastore:"value,noindex"` // No need to index price values
}

// Broker struct
type Broker struct {
	ID       *datastore.Key `datastore:"__key__"` // Primary key (Datastore auto-generates)
	Name     string         `datastore:"name"`
	Age      int64          `datastore:"age"`
	Gender   string         `datastore:"gender"`
	UserName string         `datastore:"userName"`
	Email    string         `datastore:"email"`
	Password string         `datastore:"password,noindex"` // Do not index sensitive data
	Account  Account        `datastore:"account"`
}

// Stock struct
type Stock struct {
	ID          int64   `datastore:"id"`
	Symbol      string  `datastore:"symbol"`
	CompanyName string  `datastore:"company_name"`
	LatestPrice float64 `datastore:"latest_price,noindex"`
}

// Holding struct
type Holding struct {
	ID              *datastore.Key `datastore:"__key__"` // Auto-generated key
	BrokerID        int64          `datastore:"broker_id"`
	StockSymbol     string         `datastore:"stock_symbol"` // Replacing StockID with StockSymbol
	Quantity        int64          `datastore:"quantity"`
	BuyingPrice     float64        `datastore:"buying_price,noindex"`
	CurrentPrice    float64        `datastore:"current_price"`
	CompanyName     string         `datastore:"company_name"`             // title
	TotalInvestment float64        `datastore:"total_investment,noindex"` // AvgPrice * Quantity
	CurrentValue    float64        `datastore:"current_value,noindex"`    // Market price * Quantity
	Profit          float64        `datastore:"profit,noindex"`           // CurrentValue - TotalInvestment
	ProfitPercent   float64        `datastore:"profit_percent"`
	LastUpdated     time.Time      `datastore:"last_updated"` // Timestamp for tracking
}

// StockEntry struct
type StockEntry struct {
	Stock string `datastore:"stock"`
	Count int64  `datastore:"count"`
}

// ------------------------------------------------------------------------------------------------------------------------------------

func GetBroker(id int64) (*Broker, error) {
	ctx := context.Background()
	key := datastore.IDKey("Broker", id, nil)

	err := DSClient.Get(ctx, key, &Broker{})
	if err != nil {
		log.Printf("Using Datastore Key: %+v", key) // Log key structure
		return nil, err
	}
	return &Broker{}, nil
}

// ------------------------------------------------------------------------------------------------------------------------------------
func InsertBroker(broker *Broker) error {
	ctx := context.Background()
	key := datastore.IncompleteKey("Broker", nil)
	_, err := DSClient.Put(ctx, key, broker)
	if err != nil {
		return err
	}
	return nil
}

// ------------------------------------------------------------------------------------------------------------------------------------
func GetBrokerHoldings(id int64) (*[]Holding, error) {
	var holdings []Holding
	ctx := context.Background()
	query := datastore.NewQuery("Holding").FilterField("broker_id", "=", id)
	_, err := DSClient.GetAll(ctx, query, &holdings)
	if err != nil {
		log.Printf("Failed to fetch stocks for broker: %v", err)
		return nil, err
	}
	return &holdings, nil
}

// ------------------------------------------------------------------------------------------------------------------------------------

func InsertBrokerHolding(holding *Holding) error {
	ctx := context.Background()
	key := datastore.IncompleteKey("Holding", nil)
	_, err := DSClient.Put(ctx, key, &holding)
	if err != nil {
		return err
	}
	return nil
}

// ------------------------------------------------------------------------------------------------------------------------------------
func CheckBrokerStock(symbol string) error {
	ctx := context.Background()
	var holdings *[]Holding

	query := datastore.NewQuery("holding").FilterField("stock_symbol", "=", symbol)
	if _, err := DSClient.GetAll(ctx, query, holdings); err != nil {
		return err
	}
	if holdings == nil || len(*holdings) != 0 {
		return errors.New("there is already stock hold")
	}
	return nil
}

// GetBroker retrieves a broker by ID
// func GetBroker(id string) (*Broker, error) {
// 	ctx := context.Background()
// 	key := datastore.NameKey("Broker", id, nil)

// 	var broker Broker
// 	err := DSClient.Get(ctx, key, &broker)
// 	if err != nil {
// 		return nil, err
// 	}
// }
// ------------------------------------------------------------------------------------------------------------------------------------

// 	return &broker, nil
// }

// func InsertBroker(broker *Broker) error {
// 	context := context.Background()
// 	key := datastore.NameKey("Broker", broker.UserName, nil)

// 	_, err := DSClient.Put(context, key, broker)
// 	if err != nil {
// 		log.Printf("Failed to insert broker: %v", err)
// 		return err
// 	}

// 	log.Println("Broker inserted successfully!")
// 	return nil

// }
