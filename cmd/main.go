package main

import (
	"Online_market/pkg/database"
	"Online_market/pkg/handlers"
	"Online_market/pkg/routes"
	"context"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

type OrderRequest struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func main() {
	// Подключение к базе
	database.ConnectDB()

	// Инициализация Kafka producer (если нужен)
	handlers.InitKafkaWriter()

	// Создаем и запускаем consumer в отдельной горутине
	go func() {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   "orders",
			GroupID: "order-consumer-group",
		})

		log.Println("🛒 Order Consumer started")

		for {
			m, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Fatal("Failed to read message:", err)
			}

			var order OrderRequest
			if err := json.Unmarshal(m.Value, &order); err != nil {
				log.Println("Invalid order format")
				continue
			}

			log.Printf("📥 New Order: User %d buys %d of Product %d\n",
				order.UserID, order.Quantity, order.ProductID)
		}
	}()

	// Запускаем HTTP сервер Gin
	router := gin.Default()
	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
