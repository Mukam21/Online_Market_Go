package main

import (
	"Online_market/pkg/database"
	"Online_market/pkg/models"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	database.ConnectDB()
	db := database.GetDB()
	database.AutoMigrate()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "orders",
		GroupID:     "order-consumer-group-v2",
		StartOffset: kafka.LastOffset, // чтобы читать только новые сообщения
	})

	log.Println("Kafka Order Consumer is running...")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Kafka read error:", err)
			continue
		}

		var order models.OrderRequest
		if err := json.Unmarshal(m.Value, &order); err != nil {
			log.Println("Invalid order format:", err)
			continue
		}

		log.Printf("Получен заказ: пользователь %d, товар %d, количество %d",
			order.UserID, order.ProductID, order.Quantity)

		var product models.Product
		if err := db.First(&product, order.ProductID).Error; err != nil {
			log.Printf("Товар с ID %d не найден", order.ProductID)
			continue
		}

		if product.Quantity < order.Quantity {
			log.Printf("Недостаточно товара %d. Осталось %d, запрошено %d",
				order.ProductID, product.Quantity, order.Quantity)
			continue
		}

		product.Quantity -= order.Quantity
		if err := db.Save(&product).Error; err != nil {
			log.Println("Ошибка при обновлении товара:", err)
			continue
		}

		newOrder := models.Order{
			UserID:    order.UserID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
		}
		if err := db.Create(&newOrder).Error; err != nil {
			log.Println("Ошибка при сохранении заказа:", err)
		} else {
			log.Printf("Заказ сохранен: пользователь %d заказал %d ед. товара %d",
				order.UserID, order.Quantity, order.ProductID)
		}
	}
}
