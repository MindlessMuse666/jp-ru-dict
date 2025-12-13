package db

import (
	"database/sql"
	"fmt"
	"jp-ru-dict/backend/internal/config"
	"time"

	_ "github.com/lib/pq"
)

// Connect устанавливает подключение к PostgreSQL
func Connect(cfg config.DBConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	var db *sql.DB
	var err error

	// Пытаемся подключиться с ретраями
	for range 10 {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}

		// Проверяем подключение
		if err = db.Ping(); err != nil {
			db.Close()
			time.Sleep(2 * time.Second)
			continue
		}

		// Настраиваем пул соединений
		db.SetMaxOpenConns(25)
		db.SetMaxIdleConns(5)
		db.SetConnMaxLifetime(5 * time.Minute)

		return db, nil
	}

	return nil, fmt.Errorf("не удалось подключиться к базе данных после 10 попыток: %v", err)
}
