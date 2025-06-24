package main

import (
	"encoding/json"
	"log"
	"stock-service/messaging"
	"stock-service/models"
	"stock-service/service"
	"time"
)

func connectWithRetry(amqpURL, queue string, maxRetries int, retryDelay time.Duration) (*messaging.Consumer, error) {
	var consumer *messaging.Consumer
	var err error

	for i := 0; i < maxRetries; i++ {
		consumer, err = messaging.NewConsumer(amqpURL, queue)
		if err == nil {
			return consumer, nil
		}
		log.Printf("Consumer oluşturulamadı: %v. %d. deneme, %v sonra tekrar denenecek.", err, i+1, retryDelay)
		time.Sleep(retryDelay)
	}
	return nil, err
}

func handleMessage(body []byte) {
	log.Printf("📨 Mesaj alındı: %s", body)

	var order models.Order
	if err := json.Unmarshal(body, &order); err != nil {
		log.Printf("Mesaj parse edilemedi: %v", err)
		return
	}

	service.ProcessOrder(order)
}

func main() {
	const (
		amqpURL    = "amqp://guest:guest@rabbitmq:5672/"
		queueName  = "stock-update"
		maxRetries = 10
		retryDelay = 3 * time.Second
	)

	consumer, err := connectWithRetry(amqpURL, queueName, maxRetries, retryDelay)
	if err != nil {
		log.Fatalf("RabbitMQ'ya bağlanılamadı, servis kapanıyor: %v", err)
	}
	defer consumer.Close()

	if err := consumer.Consume(handleMessage); err != nil {
		log.Fatalf("Mesaj tüketim hatası: %v", err)
	}
}
