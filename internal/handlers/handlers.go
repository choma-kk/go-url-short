package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"

	"github.com/choma-kk/go-url-short/internal/storage"
)

// Handler — структура с хранилищем ссылок
type Handler struct {
	store *storage.Storage
}

// NewHandler — конструктор для Handler
func NewHandler(store *storage.Storage) *Handler {
	return &Handler{store: store}
}

// ShortenURL — создание короткой ссылки
func (h *Handler) ShortenURL(c *gin.Context) {
	var req struct {
		URL string `json:"url"`
	}

	// Читаем JSON-запрос
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Генерируем уникальный short_id
	shortID, _ := shortid.Generate()

	// Сохраняем в PostgreSQL
	if err := h.store.SaveURL(shortID, req.URL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL"})
		return
	}

	// Возвращаем ответ с коротким URL
	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:8080/" + shortID,
	})
}

// ResolveURL — перенаправление по короткому ID
func (h *Handler) ResolveURL(c *gin.Context) {
	shortID := c.Param("short_id")

	// Получаем оригинальный URL из БД
	originalURL, err := h.store.GetURL(shortID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Short URL not found"})
		return
	}

	// Перенаправляем пользователя
	c.Redirect(http.StatusFound, originalURL)
}
