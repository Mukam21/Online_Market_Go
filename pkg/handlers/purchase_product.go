package handlers

import (
	"Online_market/pkg/database"
	"Online_market/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PurchaseProduct(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Quantity int `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&input); err != nil || input.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if product.Quantity < input.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough quantity in stock"})
		return
	}

	if product.Quantity == input.Quantity {
		database.DB.Delete(&product)
		c.JSON(http.StatusOK, gin.H{"message": "Product purchased and deleted"})
	} else {
		product.Quantity -= input.Quantity
		database.DB.Save(&product)
		c.JSON(http.StatusOK, gin.H{"message": "Product purchased", "remaining": product.Quantity})
	}
}
