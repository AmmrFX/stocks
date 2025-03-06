package main

import (
	datastore_repo "finance/datastore"
	"finance/handlers"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gregdel/pushover"
	"github.com/joho/godotenv"
)

var (
	pushoverAPIKey  string
	pushoverUserKey string
	HostURL         string
	STOCKS_URL      string
	COMPANY         string
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}
	os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	pushoverAPIKey = os.Getenv("PUSHOVER_API_KEY")
	pushoverUserKey = os.Getenv("PUSHOVER_USER_KEY")
	HostURL = os.Getenv("HOST_URL")
	STOCKS_URL = os.Getenv("STOCK")
	COMPANY = os.Getenv("Mobco")
	STOCKS_URL += COMPANY
}

func main() {
	datastore_repo.InitDatastoreClient()

	// Setup Gin Router
	r := gin.Default()

	// Define API Routes
	r.POST("/broker", handlers.InsertBroker)
	r.GET("/broker/:id", handlers.GetBroker)
	r.POST("/order/excute_order", handlers.ExcuteOrder)
	//r.GET("/broker/:id/dashboard",handlers.BrokerFinance)
	// r.GET("/broker/:id/transactions",handlers.BrokerTransactoins)
	// r.GET("/broker/:id/orders",handlers.BrokerOrders)
	// Start Server
	log.Println("Server running on port 8080...")
	r.Run(":8080")
}

// func main() {
// 	var lastPercent float64
// 	for {
// 		message, lastperc, err := StockHandler(HostURL, STOCKS_URL, COMPANY, lastPercent)
// 		lastPercent = lastperc
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 			// Handle the error as needed
// 		} else {
// 			if message != "" {
// 				pushover2(message)
// 			}
// 		}
// 		time.Sleep(5 * time.Minute)
// 	}
// }

// func main() {
// 	datastore.InitDatastoreClient() // Make sure to initialize DSClient

// 	broker := &datastore.Broker{
// 		Name:     "John Doe",
// 		Age:      "35",
// 		Gender:   "Male",
// 		UserName: "john_doe",
// 		Email:    "johndoe@example.com",
// 		Password: "securepassword",
// 		Account: datastore.Account{
// 			InitialCredit: 10000.50,
// 			Companies: []datastore.CompanyDetails{
// 				{Title: "Apple", Market: "NASDAQ", Stock: "AAPL", Index: 1, Value: 150.75},
// 				{Title: "Tesla", Market: "NASDAQ", Stock: "TSLA", Index: 2, Value: 800.30},
// 			},

// 			Stocks: []datastore.StockEntry{
// 				{Stock: "MSFT", Count: 15},
// 				{Stock: "AMZN", Count: 5},
// 			},
// 		},
// 	}

// 	err := datastore.InsertBroker(broker)
// 	if err != nil {
// 		log.Fatalf("Error inserting broker: %v", err)
// 	}
// }

func pushover2(message1 string) error {
	app := pushover.New(pushoverAPIKey)
	recipient := pushover.NewRecipient(pushoverUserKey)

	message := &pushover.Message{
		Message: message1,
	}

	_, err := app.SendMessage(message, recipient)
	if err != nil {
		return err
	}
	return nil
}

// stocks := []string{"abou-kir-fertilizers"}
// scrapeURL := "https://sa.investing.com/equities/"

// c := colly.NewCollector(colly.AllowedDomains("sa.investing.com"))

// c.OnRequest(func(r *colly.Request) {
// 	r.Headers.Set("Accept-Language", "en-US;q=0.9")
// 	fmt.Printf("Visiting %s\n", r.URL)
// })

// c.OnHTML(".text-base.font-bold.leading-6.md\\:text-xl.md\\:leading-7.rtl\\:force-ltr", func(e *colly.HTMLElement) {
// 	data := strings.TrimSpace(e.Text)
// 	fmt.Println(data)
// 	priceChangeText := strings.TrimSpace(data) // Replace with actual scraped data
// 	priceChangeText = strings.ReplaceAll(priceChangeText, "+", "")
// 	priceChangeText = strings.ReplaceAll(priceChangeText, "%", "")
// 	priceChange, err := strconv.ParseFloat(priceChangeText, 64)
// 	if err != nil {
// 		fmt.Println("Error parsing price change:", err)
// 	}

// 	if priceChange > 0.5 || priceChange < -0.5 {
// 		message := fmt.Sprintf("%s: Price change: %.2f%%", stocks, priceChange)
// 		go pushover2(message)
// 	}

// 	// Sleep for 2 minutes before checking again
// 	time.Sleep(2 * time.Minute)
// })

// for _, stock := range stocks {
// 	url := scrapeURL + stock
// 	err := c.Visit(url)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }
