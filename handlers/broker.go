package handlers

import (
	ds "finance/datastore"
	"finance/internal/types"
	"log"

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
	err := ds.InsertBroker(&broker)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert broker"})
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

	broker, err := ds.GetBroker(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Broker not found"})
	}
	c.JSON(http.StatusOK, broker)
}

// ------------------------------------------------------------------------------------------------------------------------------------
func ExcuteOrder(c *gin.Context) {
	var order types.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

}
