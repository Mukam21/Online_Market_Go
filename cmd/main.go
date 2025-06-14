package main

import (
	"Online_market/pkg/database"
	"Online_market/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	router := gin.Default()

	routes.SetupRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
