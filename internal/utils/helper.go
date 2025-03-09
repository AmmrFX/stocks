package utils

import (
	"encoding/json"
	"finance/internal/types"
	"fmt"
	"net/http"
	"strconv"
)

var (
	// pushoverAPIKey  string
	// pushoverUserKey string
	HostURL    string
	STOCKS_URL string
	COMPANY    string
)

// ------------------------------------------------------------------------------------------------------------------------------------

// func (c *types.Company) Manipulator(lastPercent float64, company string) (string, float64) {
// 	var message string
// 	initialCredit := 983.7
// 	noOfStocks := 3
// 	stockPrice := c.CompanyTab.PriceBar

// 	todayPercentage, err := parseTodayPercentage(stockPrice.ChangePercentage)
// 	if err != nil {
// 		fmt.Println("Invalid change percentage format:", err)
// 		return "", 0
// 	}

// 	profit := c.calculateProfit(float64(noOfStocks), initialCredit)
// 	InvestmentValue := profit + initialCredit

// 	if todayPercentage >= 0.5 {
// 		if todayPercentage-lastPercent > 0.5 {
// 			message = fmt.Sprintf(" Profit from stock "+company+": Total Change percentage is %.2f%% %s, Total Investment %f, profit %f value %s", todayPercentage, stockPrice.Status, InvestmentValue, profit, stockPrice.Value)
// 			return message, lastPercent
// 		}
// 	}

// 	if todayPercentage < -0.5 {
// 		message = fmt.Sprintf("Loss "+company+": Total Change percentage is %.2f%% %s, Total Investment %f,  profit %f value %s", todayPercentage, stockPrice.Status, InvestmentValue, profit, stockPrice.Value)
// 		return message, lastPercent
// 	}

// 	return "", 0
// }

// // ------------------------------------------------------------------------------------------------------------------------------------

// func parseTodayPercentage(percentage string) (float64, error) {
// 	percentage = strings.TrimSuffix(percentage, "%")
// 	return strconv.ParseFloat(percentage, 64)
// }

// // ------------------------------------------------------------------------------------------------------------------------------------

// func (c *Company) calculateProfit(noOfStocks, initialCredit float64) float64 {
// 	stockPrice, _ := strconv.ParseFloat(c.CompanyTab.PriceBar.Value, 64)
// 	profit := (noOfStocks * stockPrice) - initialCredit
// 	return profit
// }

// // ------------------------------------------------------------------------------------------------------------------------------------
// func ExecuteTrade(order Order) {
// 	log.Printf("Executing %s order for %d shares of %s at $%.2f",
// 		order.OrderType, order.Quantity, order.Stock, order.Price)

// 	// TODO: Add database operations (update stock positions, send notification, etc.)
// }

// ------------------------------------------------------------------------------------------------------------------------------------
// CalculateTotalInvestment computes the total cost of buying stocks.
func CalculateTotalInvestment(buyingPrice float64, quantity int64) float64 {
	return buyingPrice * float64(quantity)
}

// CalculateProfit computes the profit based on the current value and total investment.
func CalculateProfit(totalInvestment float64, quantity int64, currentPrice float64) float64 {
	currentValue := float64(quantity) * currentPrice
	return currentValue - totalInvestment
}

// ------------------------------------------------------------------------------------------------------------------------------------
func CalculateProfitPercent(profit float64, totalInvestment float64) float64 {
	if totalInvestment == 0 {
		return 0 // Avoid division by zero
	}
	return (profit / totalInvestment) * 100
}

// ------------------------------------------------------------------------------------------------------------------------------------
func GetStockCurrentValue(quantity int64, currentPrice float64) float64 {
	if quantity == 0 || currentPrice == 0 {
		return 0
	}
	return (float64(quantity) * currentPrice)
}

// ------------------------------------------------------------------------------------------------------------------------------------
func GetHoldingStocksInfo(stockSymbol string) (*types.Holding, error) {
	//message, lastperc, err := main.StockHandler(HostURL, STOCKS_URL, COMPANY, lastPercent)
	var holdingStock *types.Holding

	stock, err := GetStockBySymbol(stockSymbol)
	if err != nil {
		return nil, err
	}

	currentPrice, err := strconv.ParseFloat(stock.PriceBar.Value, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return nil, err
	}

	holdingStock = &types.Holding{
		CurrentPrice: currentPrice,
		CompanyName:  stock.CompanyName,
	}
	return holdingStock, nil

}

// ------------------------------------------------------------------------------------------------------------------------------------
func GetStockBySymbol(stockSymbol string) (*types.Stock, error) {
	companyStock := &types.Company{}
	var stock *types.Stock

	resp, err := http.Get(HostURL + STOCKS_URL + stockSymbol)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(companyStock); err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	stock = &types.Stock{
		Symbol:      companyStock.CompanyTab.Stock,
		Market:      companyStock.CompanyTab.Market,
		Index:       companyStock.CompanyTab.Index,
		CompanyName: companyStock.CompanyTab.Title,
		PriceBar:    types.StockPriceBar(companyStock.CompanyTab.PriceBar),
	}

	return stock, nil
}
