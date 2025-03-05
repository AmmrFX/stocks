package datastore

type Account struct {
	InitialCredit float64          `datastore:"initialCredit"`
	Companies     []CompanyDetails `datastore:"companies"`
	Stocks        []StockEntry     `datastore:"stocks"` // ✅ Use slice instead of map
}

// CompanyDetails struct (Fixed)
type CompanyDetails struct {
	Title  string  `datastore:"title"`
	Market string  `datastore:"market"`
	Stock  string  `datastore:"stock"`
	Index  int     `datastore:"index"`
	Value  float64 `datastore:"priceBar"` // Fixed: "value" is now "Value"
}

// Broker struct
type Broker struct {
	Name     string  `datastore:"name"`
	Age      string  `datastore:"age"`
	Gender   string  `datastore:"gender"`
	UserName string  `datastore:"userName"`
	Email    string  `datastore:"email"`
	Password string  `datastore:"password"`
	Account  Account `datastore:"account"`
}
type StockEntry struct {
	Stock string `datastore:"stock"`
	Count int    `datastore:"count"`
}

// // GetBroker retrieves a broker by ID
// func GetBroker(id string) (*Broker, error) {
// 	ctx := context.Background()
// 	key := datastore.NameKey("Broker", id, nil)

// 	var broker Broker
// 	err := DSClient.Get(ctx, key, &broker)
// 	if err != nil {
// 		return nil, err
// 	}

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
