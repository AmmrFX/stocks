package types

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// ------------------------------------------------------------------------------------------------------------------------------------

func (c *Company) Manipulator(lastPercent float64, company string) (string, float64) {
	var message string
	initialCredit := 983.7
	noOfStocks := 3
	stockPrice := c.CompanyTab.PriceBar

	todayPercentage, err := parseTodayPercentage(stockPrice.ChangePercentage)
	if err != nil {
		fmt.Println("Invalid change percentage format:", err)
		return "", 0
	}

	profit := c.calculateProfit(float64(noOfStocks), initialCredit)
	InvestmentValue := profit + initialCredit

	if todayPercentage >= 0.5 {
		if todayPercentage-lastPercent > 0.5 {
			message = fmt.Sprintf(" Profit from stock "+company+": Total Change percentage is %.2f%% %s, Total Investment %f, profit %f value %s", todayPercentage, stockPrice.Status, InvestmentValue, profit, stockPrice.Value)
			return message, lastPercent
		}
	}

	if todayPercentage < -0.5 {
		message = fmt.Sprintf("Loss "+company+": Total Change percentage is %.2f%% %s, Total Investment %f,  profit %f value %s", todayPercentage, stockPrice.Status, InvestmentValue, profit, stockPrice.Value)
		return message, lastPercent
	}

	return "", 0
}

// ------------------------------------------------------------------------------------------------------------------------------------

func parseTodayPercentage(percentage string) (float64, error) {
	percentage = strings.TrimSuffix(percentage, "%")
	return strconv.ParseFloat(percentage, 64)
}

// ------------------------------------------------------------------------------------------------------------------------------------

func (c *Company) calculateProfit(noOfStocks, initialCredit float64) float64 {
	stockPrice, _ := strconv.ParseFloat(c.CompanyTab.PriceBar.Value, 64)
	profit := (noOfStocks * stockPrice) - initialCredit
	return profit
}

// ------------------------------------------------------------------------------------------------------------------------------------
func ExecuteTrade(order Order) {
	log.Printf("Executing %s order for %d shares of %s at $%.2f",
		order.OrderType, order.Quantity, order.Stock, order.Price)

	// TODO: Add database operations (update stock positions, send notification, etc.)
}
