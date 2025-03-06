package handlers

import (
	"finance/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExcuteOrder(c *gin.Context) {
	var order types.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

}
