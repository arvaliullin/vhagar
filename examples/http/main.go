// Package main содержит пример HTTP-сервера с использованием библиотеки chi.
package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
)

const (
	defaultPort        = "8080"
	readHeaderTimeout  = 5 * time.Second
	readTimeout        = 10 * time.Second
	writeTimeout       = 10 * time.Second
	idleTimeout        = 120 * time.Second
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
		cfg.Port = defaultPort
	}

	r := chi.NewRouter()

	r.Get("/health", healthHandler)

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:     writeTimeout,
		IdleTimeout:       idleTimeout,
	}

	log.Printf("Сервер запущен на порту %s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	response := HealthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Ошибка кодирования ответа: %v", err)
	}
}
