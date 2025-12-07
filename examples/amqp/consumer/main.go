package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/streadway/amqp"
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

	err = ch.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		log.Fatalf("Не удалось установить QoS: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Не удалось зарегистрировать потребителя: %v", err)
	}

	log.Printf("Ожидание сообщений из очереди '%s'. Для выхода нажмите CTRL+C", q.Name)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		for d := range msgs {
			log.Printf(" [x] Получено сообщение: %s", string(d.Body))

			err := d.Ack(false)
			if err != nil {
				log.Printf("Ошибка подтверждения сообщения: %v", err)
			} else {
				log.Printf(" [✓] Сообщение обработано и подтверждено")
			}
		}
	}()

	<-sigChan
	log.Println("\nПолучен сигнал завершения. Завершение работы...")
}
