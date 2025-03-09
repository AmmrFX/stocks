package handlers

import (
	ds "finance/datastore"
	"finance/internal/types"
	"finance/internal/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ------------------------------------------------------------------------------------------------------------------------------------
func GetBrokerHolding(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Fetching Broker with ID: %s", id) // Log broker ID

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	brokerID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid broker ID"})
		return
	}

	_, err = ds.GetBroker(brokerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Broker not found"})
	}

	holding, err := ds.GetBrokerHoldings(brokerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Broker not found"})
		return
	}

	c.JSON(http.StatusOK, holding)
}

// ------------------------------------------------------------------------------------------------------------------------------------
func InsertHolding(c *gin.Context) {
	var holdings *[]types.Holding

	if err := c.ShouldBindJSON(&holdings); err != nil { // Get the request and bind it to the types holding.
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	brokerID, err := strconv.ParseInt(c.Param("id"), 10, 64) // Get the broker id and convert it to check with it.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid broker ID"})
		return
	}

	if _, err := ds.GetBroker(brokerID); err != nil { // Check if this broker with this id is registered.
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("broker id is invalid: %v", err)})
		return
	}

	var dsHoldings []ds.Holding
	var failedStocks []string
	var successedStocks []string

	for _, holding := range *holdings {

		// Check that this stock isn't already added
		//  if its already added it must update the value of quantity
		//  but with excuting orders or we can add same stock but with differnt quantity and buying price.
		if err := ds.CheckBrokerStock(holding.StockSymbol); err != nil {
			failedStocks = append(failedStocks, fmt.Sprintf("%s (already held)", holding.StockSymbol))
			continue
		}

		// valid all input

		// get all all the info stocks holding by this broker but view only some of our need.
		stockInfo, err := utils.GetHoldingStocksInfo(holding.StockSymbol)
		if err != nil {
			failedStocks = append(failedStocks, fmt.Sprintf("%s (stock info error: %v)", holding.StockSymbol, err))
			continue
		}

		holding.CompanyName = stockInfo.CompanyName
		holding.CurrentPrice = stockInfo.CurrentPrice

		holding.TotalInvestment = utils.CalculateTotalInvestment(holding.BuyingPrice, holding.Quantity)
		holding.CurrentValue = utils.GetStockCurrentValue(holding.Quantity, holding.CurrentPrice)
		holding.Profit = utils.CalculateProfit(holding.TotalInvestment, holding.Quantity, holding.CurrentValue)
		holding.ProfitPercent = utils.CalculateProfitPercent(holding.Profit, holding.TotalInvestment)

		dsHoldings = append(dsHoldings, ds.Holding{
			BrokerID:        brokerID,
			StockSymbol:     holding.StockSymbol,
			Quantity:        holding.Quantity,
			BuyingPrice:     holding.BuyingPrice,
			CurrentValue:    holding.CurrentValue,
			TotalInvestment: holding.TotalInvestment,
			Profit:          holding.Profit,
			ProfitPercent:   holding.ProfitPercent,
			CurrentPrice:    holding.CurrentPrice,
			CompanyName:     holding.CompanyName,
			LastUpdated:     time.Now(),
		})
		successedStocks = append(successedStocks, fmt.Sprintf("%s (stock inserted successfully:)", holding.StockSymbol))
	}

	// valid all input

	// Insert only valid holdings
	if len(dsHoldings) > 0 {
		if err := ds.InsertBrokerHoldings(dsHoldings); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to insert holdings: %v", err)})
			return
		}

	}

	// Return response based on success/failures
	if len(successedStocks) > 0 {
		c.JSON(http.StatusPartialContent, gin.H{
			"success":       successedStocks,
			"failed_stocks": failedStocks,
		})
		return
	}
	

	c.JSON(http.StatusExpectationFailed, gin.H{"error": "all stocks wasn't inserted!"})
}
