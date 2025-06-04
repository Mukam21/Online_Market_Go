package routes

import (
	"Online_market/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", handlers.GetProduct)
			products.GET("/id", handlers.GetProducts)
			products.POST("", handlers.CreateProduct)
			products.PUT("/id", handlers.UpdateProduct)
			products.DELETE("/id", handlers.DeleteProduct)
		}

		api.POST("/register", handlers.User_Registr)
	}
}
