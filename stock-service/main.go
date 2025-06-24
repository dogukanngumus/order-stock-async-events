package main

import (
	"encoding/json"
	"log"
	"stock-service/messaging"
	"stock-service/models"
	"stock-service/service"
)

func main() {
	consumer, err := messaging.NewConsumer("amqp://guest:guest@localhost:5672/", "stock-update")
	if err != nil {
		log.Fatalf("Consumer oluşturulamadı: %v", err)
	}
	defer consumer.Close()

	err = consumer.Consume(func(body []byte) {
		log.Printf("📨 Mesaj alındı: %s", body)

		var order models.Order
		if err := json.Unmarshal(body, &order); err != nil {
			log.Printf("Mesaj parse edilemedi: %v", err)
			return
		}

		service.ProcessOrder(order)
	})

	if err != nil {
		log.Fatalf("Mesaj tüketim hatası: %v", err)
	}
}
