package routes

import (
	"Online_market/pkg/handlers"
	"Online_market/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", handlers.GetProducts)
			products.GET("/:id", handlers.GetProduct)
			products.POST("", handlers.CreateProduct)
			products.PUT("/:id", handlers.UpdateProduct)
			products.DELETE("/:id", handlers.DeleteProduct)
			products.POST("/:id/purchase", handlers.PurchaseProduct)
			products.GET("/search", handlers.SearchProducts)

		}

		api.POST("/orders", handlers.CreateOrder)

		api.POST("/register", handlers.User_Registr)
		api.POST("/login", handlers.User_Login)

		auth := api.Group("/")
		auth.Use(middleware.JWTAuthMiddleware())
		{
			auth.GET("/user_profile", handlers.GetProfile)
		}
	}
}
