package handlers

import (
	ds "finance/datastore"
	"finance/internal/types"
	"log"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

// InsertBroker inserts a broker into Datastore
func InsertBroker(c *gin.Context) {
	var broker types.Broker
	if err := c.ShouldBindJSON(&broker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	dsBroker := ds.Broker{
		Name:     broker.Name,
		Age:      broker.Age,
		Gender:   broker.Gender,
		UserName: broker.UserName,
		Email:    broker.Email,
		Password: broker.Password,
		Account: ds.Account{
			InitialCredit: broker.Account.InitialCredit,
			Companies:     mapCompanies(broker.Account.Companies), // ✅ Correctly mapped
			Stocks:        mapStocks(broker.Account.Stocks),       // ✅ Correctly mapped
		},
	}

	err := ds.InsertBroker(&dsBroker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert broker %v"})
		log.Printf("Fetching Broker with ID: %v", err) // Log broker ID

		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Broker inserted successfully"})
}

// ------------------------------------------------------------------------------------------------------------------------------------

// GetBroker retrieves a broker by ID
func GetBroker(c *gin.Context) {
	id := c.Param("id")

	log.Printf("Fetching Broker with ID: %s", id) // Log broker ID

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	brokerID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid broker ID"})
		return
	}

	broker, err := ds.GetBroker(brokerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Broker not found"})
	}

	c.JSON(http.StatusOK, broker)
}

// ------------------------------------------------------------------------------------------------------------------------------------
func mapCompanies(companies []types.CompanyDetails) []ds.CompanyDetails {
	var dsCompanies []ds.CompanyDetails
	for _, company := range companies {
		dsCompanies = append(dsCompanies, ds.CompanyDetails{
			Title:  company.Title,
			Market: company.Market,
			Stock:  company.Stock,
			Index:  company.Index,
			Value:  company.Value,
		})
	}
	return dsCompanies
}

// ------------------------------------------------------------------------------------------------------------------------------------

// Helper function to map Stocks
func mapStocks(stocks []types.Stock) []ds.Stock {
	var dsStocks []ds.Stock
	for _, stock := range stocks {
		dsStocks = append(dsStocks, ds.Stock{
			ID:          stock.ID,
			Symbol:      stock.Symbol,
			CompanyName: stock.CompanyName,
			// LatestPrice: stock.LatestPrice,
		})
	}
	return dsStocks
}
