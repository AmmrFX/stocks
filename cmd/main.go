package main

import (
	"fmt"
	"os"
	"time"

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

	pushoverAPIKey = os.Getenv("PUSHOVER_API_KEY")
	pushoverUserKey = os.Getenv("PUSHOVER_USER_KEY")
	HostURL = os.Getenv("HOST_URL")
	STOCKS_URL = os.Getenv("STOCK")
	COMPANY = os.Getenv("Mobco")
	STOCKS_URL += COMPANY
}

func main() {
	var lastPercent float64
	for {
		message, lastperc, err := StockHandler(HostURL, STOCKS_URL, COMPANY, lastPercent)
		lastPercent = lastperc
		if err != nil {
			fmt.Println("Error:", err)
			// Handle the error as needed
		} else {
			if message != "" {
				pushover2(message)

			}
		}

		time.Sleep(5 * time.Minute)
	}
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
