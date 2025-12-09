// Package main содержит пример производителя сообщений в очередь RabbitMQ.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

const (
	messageDelay = 500 * time.Millisecond
)

func main() {
	amqpURL := os.Getenv("AMQP_URL")
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@127.0.0.1:5672/"
	}
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatalf("Не удалось подключиться к RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Не удалось открыть канал: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"example_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Не удалось объявить очередь: %v", err)
	}

	log.Printf("Очередь '%s' объявлена. Готов к отправке сообщений...", q.Name)

	for i := 1; i <= 10; i++ {
		body := fmt.Sprintf("Сообщение номер %d", i)

		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(body),
				Timestamp:    time.Now(),
			})
		if err != nil {
			log.Fatalf("Не удалось отправить сообщение: %v", err)
		}

		log.Printf(" [x] Отправлено: %s", body)
		time.Sleep(messageDelay)
	}

	log.Println("Все сообщения отправлены!")
}
