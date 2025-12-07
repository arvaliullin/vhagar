package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// Подключение к RabbitMQ серверу
	// По умолчанию: amqp://guest:guest@localhost:5672/
	amqpURL := os.Getenv("AMQP_URL")
	if amqpURL == "" {
		amqpURL = "amqp://guest:guest@localhost:5672/"
	}
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Fatalf("Не удалось подключиться к RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Открытие канала
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Не удалось открыть канал: %v", err)
	}
	defer ch.Close()

	// Объявление очереди
	// durable=true - очередь будет сохраняться после перезапуска сервера
	// autoDelete=false - очередь не будет удаляться автоматически
	// exclusive=false - очередь доступна для всех соединений
	// noWait=false - ждем подтверждения от сервера
	q, err := ch.QueueDeclare(
		"example_queue", // имя очереди
		true,            // durable - устойчивая
		false,           // autoDelete - не удаляется автоматически
		false,           // exclusive - не эксклюзивная
		false,           // noWait - ждем подтверждения
		nil,             // arguments - дополнительные аргументы
	)
	if err != nil {
		log.Fatalf("Не удалось объявить очередь: %v", err)
	}

	log.Printf("Очередь '%s' объявлена. Готов к отправке сообщений...", q.Name)

	// Отправка нескольких сообщений
	for i := 1; i <= 10; i++ {
		body := fmt.Sprintf("Сообщение номер %d", i)

		err = ch.Publish(
			"",     // exchange - используем default exchange
			q.Name, // routing key - имя очереди
			false,  // mandatory - не обязательное
			false,  // immediate - не немедленное
			amqp.Publishing{
				DeliveryMode: amqp.Persistent, // Persistent - сообщение будет сохранено на диск
				ContentType:  "text/plain",
				Body:         []byte(body),
				Timestamp:    time.Now(),
			})
		if err != nil {
			log.Fatalf("Не удалось отправить сообщение: %v", err)
		}

		log.Printf(" [x] Отправлено: %s", body)
		time.Sleep(500 * time.Millisecond) // Небольшая задержка между сообщениями
	}

	log.Println("Все сообщения отправлены!")
}

