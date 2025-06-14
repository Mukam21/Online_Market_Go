package handlers

import (
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

type OrderRequest struct {
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}

func CreateOrder(c *gin.Context) {
	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	msgBytes, _ := json.Marshal(req)
	err := kafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("order"),
		Value: msgBytes,
	})

	if err != nil {
		log.Println("Failed to send order:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order placed"})
}
