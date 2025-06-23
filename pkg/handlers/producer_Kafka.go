package handlers

import (
	"Online_market/pkg/models"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/segmentio/kafka-go"

	"github.com/gin-gonic/gin"
)

var kafkaWriter *kafka.Writer

func InitKafkaWriter() {
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "orders",
		Balancer: &kafka.LeastBytes{},
	}
}

func CreateOrder(c *gin.Context) {
	var order models.OrderRequest
	if err := c.ShouldBindJSON(&order); err != nil {
		log.Println("Ошибка привязки JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order format"})
		return
	}

	// Проверка на нули
	if order.UserID == 0 || order.ProductID == 0 || order.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные значения заказа"})
		return
	}

	msgBytes, err := json.Marshal(order)
	if err != nil {
		log.Println("Ошибка маршалинга заказа:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	log.Printf("Отправка Kafka сообщения: %s", string(msgBytes))

	err = kafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Value: msgBytes,
		})

	if err != nil {
		log.Println("Ошибка Kafka:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kafka send failed"})
		return
	}

	log.Printf("Order sent to Kafka: %+v", order)
	c.JSON(http.StatusOK, gin.H{"message": "Order placed"})
}
