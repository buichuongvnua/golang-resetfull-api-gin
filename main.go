package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	configuration "example.com/configs"
	"example.com/pkg/kafkas"
	"example.com/schemas"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

func init() {
	godotenv.Load()
	env := os.Getenv("APP_ENV")
	configuration.AppConfig = configuration.New(env)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		productTopic := "producer-product-table-testing"
		publisher, _ := kafkas.NewPublisher([]string{"localhost:8098"})
		product := map[string]string{
			"name": "Hello",
		}
		marshalProduct, _ := json.Marshal(product)
		publisher.Publish(context.Background(), productTopic, "product", marshalProduct)

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}

	productConsumer := kafkas.Consumer[schemas.Product]{
		Dialer: dialer,
		Topic:  "producer-product-table-testing",
	}
	productConsumer.CreateConnection([]string{"localhost:8098"}, "producer-product-table-testing_group")
	go productConsumer.Read(schemas.Product{}, func(product schemas.Product, err error) {
		fmt.Println("Consumer collect")
	})
	go productConsumer.Read(schemas.Product{}, func(product schemas.Product, err error) {
		fmt.Println("Consumer collect")
	})
	r.Run()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
