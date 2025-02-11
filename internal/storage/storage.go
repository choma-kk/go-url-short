package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

// Storage — структура для работы с PostgreSQL
type Storage struct {
	db *pgx.Conn
}

// NewStorage — подключение к PostgreSQL
func NewStorage() (*Storage, error) {
	dbURL := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("✅ Connected to PostgreSQL")
	return &Storage{db: conn}, nil
}

// SaveURL — сохраняем ссылку в базе
func (s *Storage) SaveURL(shortID, originalURL string) error {
	_, err := s.db.Exec(context.Background(),
		"INSERT INTO urls (short_id, original_url) VALUES ($1, $2)", shortID, originalURL)
	return err
}

// GetURL — получаем оригинальный URL по short_id
func (s *Storage) GetURL(shortID string) (string, error) {
	var originalURL string
	err := s.db.QueryRow(context.Background(),
		"SELECT original_url FROM urls WHERE short_id = $1", shortID).Scan(&originalURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}
