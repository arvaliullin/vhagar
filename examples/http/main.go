package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type Config struct {
	Port string `yaml:"port"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	cfg := Config{
		Port: os.Getenv("PORT"),
	}
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	r := chi.NewRouter()

	r.Get("/health", healthHandler)

	log.Printf("Сервер запущен на порту %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Ошибка кодирования ответа: %v", err)
	}
}
