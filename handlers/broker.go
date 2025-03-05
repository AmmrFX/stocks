package handlers

import (
	"context"
	ds "finance/datastore"
	"finance/internal/types"
	"log"

	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

// InsertBroker inserts a broker into Datastore
func InsertBroker(c *gin.Context) {
	ctx := context.Background()
	var broker types.Broker

	if err := c.ShouldBindJSON(&broker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	key := datastore.IncompleteKey("Broker", nil)
	_, err := ds.DSClient.Put(ctx, key, &broker)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert broker"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Broker inserted successfully"})
}

// GetBroker retrieves a broker by ID
func GetBroker(c *gin.Context) {
	ctx := context.Background()
	id := c.Param("id")
	log.Printf("Fetching Broker with ID: %s", id) // Log broker ID
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	key := datastore.NameKey("Broker", id, nil)
	var broker ds.Broker

	err := ds.DSClient.Get(ctx, key, &broker)
	if err != nil {
		log.Printf("Using Datastore Key: %+v", key) // Log key structure
		c.JSON(http.StatusNotFound, gin.H{"error": "Broker not found"})
		return
	}

	c.JSON(http.StatusOK, broker)
}
