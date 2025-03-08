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
	var holding *types.Holding

	if err := c.ShouldBindJSON(&holding); err != nil { // Get the request and bind it to the types holding.
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

	// Check that this stock isn't already added
	//  if its already added it must update the value of quantity
	//  but with excuting orders or we can add same stock but with differnt quantity and buying price.
	if err := ds.CheckBrokerStock(holding.StockSymbol); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("stock is already hold by you: %v", err)})
		return
	}

	// valid all input

	// get all all the info stocks holding by this broker but view only some of our need.
	if holding, err = utils.GetHoldingStocksInfo(holding.StockSymbol); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("can't get stock info: %v", err)})
		return
	}

	holding.TotalInvestment = utils.CalculateTotalInvestment(holding.BuyingPrice, holding.Quantity)
	holding.CurrentValue = utils.GetStockCurrentValue(holding.Quantity, holding.CurrentPrice)
	holding.Profit = utils.CalculateProfit(holding.TotalInvestment, holding.Quantity, holding.CurrentValue)
	holding.ProfitPercent = utils.CalculateProfitPercent(holding.Profit, holding.TotalInvestment)

	dsholding := ds.Holding{
		StockSymbol:     holding.StockSymbol,
		Quantity:        holding.Quantity,
		BuyingPrice:     holding.BuyingPrice,
		CurrentValue:    holding.CurrentValue,
		TotalInvestment: holding.TotalInvestment,
		Profit:          holding.Profit,
		ProfitPercent:   holding.ProfitPercent,
		CompanyName:     holding.CompanyName,
		LastUpdated:     time.Now(),
	}

	if err := ds.InsertBrokerHolding(&dsholding); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to insert broker holding: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "broker holding inserted successfully!"})
}
