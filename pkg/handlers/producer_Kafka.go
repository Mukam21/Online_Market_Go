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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order format"})
		return
	}

	msgBytes, _ := json.Marshal(order)
	err := kafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Value: msgBytes,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kafka send failed"})
		return
	}

	log.Printf("Order sent to Kafka: %v\n", order)
	c.JSON(http.StatusOK, gin.H{"message": "Order placed"})
}
