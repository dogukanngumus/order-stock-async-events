package messaging

import (
	"fmt"

	"github.com/streadway/amqp"
)

type ConsumerInterface interface {
	Consume(handler func(body []byte)) error
	Close()
}

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewConsumer(amqpURL string, queueName string) (*Consumer, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq bağlantı hatası: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("kanal oluşturulamadı: %w", err)
	}

	_, err = ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("kuyruk tanımlanamadı: %w", err)
	}

	consumer := &Consumer{
		conn:    conn,
		channel: ch,
		queue:   queueName,
	}

	return consumer, nil
}

func (c *Consumer) Consume(handler func(body []byte)) error {
	msgs, err := c.channel.Consume(
		c.queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("mesaj tüketilemedi: %w", err)
	}

	// go routine oluşturuldu
	go func() {
		for msg := range msgs {
			handler(msg.Body)
		}
	}()

	fmt.Println("📡 Consumer aktif, mesajlar dinleniyor...")
	select {}
}

func (c *Consumer) Close() {
	c.channel.Close()
	c.conn.Close()
}
