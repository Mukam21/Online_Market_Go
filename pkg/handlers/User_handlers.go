package handlers

import (
	"Online_market/pkg/database"
	"Online_market/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func User_Registr(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}
