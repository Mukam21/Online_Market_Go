package handlers

import (
	"Online_market/pkg/database"
	"Online_market/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchProducts(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query 'q' is required"})
		return
	}

	var products []models.Product
	searchPattern := query + "%" // для поиска "начинается с"

	if err := database.DB.Where("name ILIKE ?", searchPattern).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if len(products) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, products)
}
