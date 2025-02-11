package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"

	"github.com/choma-kk/go-url-short/internal/handlers"
	"github.com/choma-kk/go-url-short/internal/storage"
)

func main() {
	// Загружаем переменные окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Подключаемся к PostgreSQL
	store, err := storage.NewStorage()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	// Создаем обработчики
	h := handlers.NewHandler(store)

	// Создаем роутер
	router := gin.Default()
	router.POST("/shorten", h.ShortenURL)
	router.GET("/:short_id", h.ResolveURL)

	// Запускаем сервер
	fmt.Println("✅ Server is running on :8080")
	router.Run(":8080")
}
