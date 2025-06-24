package main

import (
	"log"
	"stock-service/db"
	"stock-service/inbox"
	"stock-service/messaging"
	"stock-service/models"
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

func main() {
	const (
		amqpURL    = "amqp://guest:guest@rabbitmq:5672/"
		queueName  = "stock-update"
		maxRetries = 10
		retryDelay = 3 * time.Second
		dsn        = "host=postgres user=user password=password dbname=orderdb port=5432 sslmode=disable"
	)

	dbConn := db.ConnectWithRetry(dsn, maxRetries, retryDelay)
	if err := dbConn.AutoMigrate(&models.InboxMessage{}); err != nil {
		log.Fatalf("📛 Inbox tablosu migrate edilemedi: %v", err)
	}

	consumer, err := connectWithRetry(amqpURL, queueName, maxRetries, retryDelay)
	if err != nil {
		log.Fatalf("RabbitMQ'ya bağlanılamadı, servis kapanıyor: %v", err)
	}
	defer consumer.Close()

	if err := consumer.Consume(func(body []byte) {
		inbox.HandleInboxMessage(dbConn, body)
	}); err != nil {
		log.Fatalf("Mesaj tüketim hatası: %v", err)
	}
}
