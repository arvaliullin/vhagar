package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/streadway/amqp"
)

func main() {
	// Подключение к RabbitMQ серверу
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

	// Объявление очереди (должно совпадать с producer)
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

	// Настройка качества обслуживания (QoS)
	// prefetchCount=1 - обрабатываем по одному сообщению за раз
	// prefetchSize=0 - без ограничения по размеру
	// global=false - применяется только к текущему каналу
	err = ch.Qos(
		1,     // prefetchCount
		0,     // prefetchSize
		false, // global
	)
	if err != nil {
		log.Fatalf("Не удалось установить QoS: %v", err)
	}

	// Регистрация потребителя
	msgs, err := ch.Consume(
		q.Name, // queue - имя очереди
		"",     // consumer - пустое имя (сервер сгенерирует автоматически)
		false,  // autoAck - ручное подтверждение получения
		false,  // exclusive - не эксклюзивный
		false,  // noLocal - разрешаем сообщения от этого же соединения
		false,  // noWait - ждем подтверждения
		nil,    // args - дополнительные аргументы
	)
	if err != nil {
		log.Fatalf("Не удалось зарегистрировать потребителя: %v", err)
	}

	log.Printf("Ожидание сообщений из очереди '%s'. Для выхода нажмите CTRL+C", q.Name)

	// Обработка сигналов для корректного завершения
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Обработка сообщений
	go func() {
		for d := range msgs {
			log.Printf(" [x] Получено сообщение: %s", string(d.Body))

			// Имитация обработки сообщения
			// В реальном приложении здесь будет бизнес-логика

			// Подтверждение получения сообщения
			// Если не подтвердить, сообщение вернется в очередь
			err := d.Ack(false)
			if err != nil {
				log.Printf("Ошибка подтверждения сообщения: %v", err)
			} else {
				log.Printf(" [✓] Сообщение обработано и подтверждено")
			}
		}
	}()

	// Ожидание сигнала завершения
	<-sigChan
	log.Println("\nПолучен сигнал завершения. Завершение работы...")
}
