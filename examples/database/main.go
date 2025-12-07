package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

// connectDB устанавливает соединение с базой данных PostgreSQL.
func connectDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия соединения: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка проверки соединения: %w", err)
	}

	return db, nil
}

// createTable создает таблицу для примера, если она не существует.
func createTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := db.Exec(query)
	return err
}

// insertUser вставляет нового пользователя в базу данных.
func insertUser(db *sql.DB, name, email string) (int, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	var id int
	err := db.QueryRow(query, name, email).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("ошибка вставки пользователя: %w", err)
	}
	return id, nil
}

// getAllUsers получает всех пользователей из базы данных.
func getAllUsers(db *sql.DB) error {
	rows, err := db.Query("SELECT id, name, email, created_at FROM users ORDER BY id")
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	fmt.Println("\nВсе пользователи в базе данных:")
	fmt.Println("ID | Имя           | Email              | Создан")
	fmt.Println("---|---------------|--------------------|-------------------")

	for rows.Next() {
		var id int
		var name, email string
		var createdAt time.Time
		if err := rows.Scan(&id, &name, &email, &createdAt); err != nil {
			return fmt.Errorf("ошибка сканирования строки: %w", err)
		}
		fmt.Printf("%-3d| %-13s | %-18s | %s\n", id, name, email, createdAt.Format("2006-01-02 15:04:05"))
	}

	return rows.Err()
}

func main() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable"
	}

	db, err := connectDB(connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close()

	log.Println("Успешно подключено к базе данных PostgreSQL")

	if err := createTable(db); err != nil {
		log.Fatalf("Не удалось создать таблицу: %v", err)
	}
	log.Println("Таблица 'users' готова к использованию")

	users := []struct {
		name  string
		email string
	}{
		{"Иван Иванов", "ivan@example.com"},
		{"Мария Петрова", "maria@example.com"},
		{"Алексей Сидоров", "alex@example.com"},
	}

	for _, u := range users {
		id, err := insertUser(db, u.name, u.email)
		if err != nil {
			log.Printf("Ошибка при вставке пользователя %s: %v", u.name, err)
			continue
		}
		log.Printf("Добавлен пользователь: ID=%d, Имя=%s, Email=%s", id, u.name, u.email)
	}

	if err := getAllUsers(db); err != nil {
		log.Fatalf("Ошибка при получении пользователей: %v", err)
	}

	log.Println("\nПример работы с базой данных завершен успешно!")
}
