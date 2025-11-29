package config

import (
	"os"
	"strconv"
)

// Config содержит все настройки приложения
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Kafka    KafkaConfig
}

// DatabaseConfig настройки подключения к PostgreSQL
type DatabaseConfig struct {
	URL string
}

// ServerConfig настройки HTTP сервера
type ServerConfig struct {
	Port int
}

// KafkaConfig настройки подключения к Kafka
type KafkaConfig struct {
	URL   string
	Topic string
}

// LoadConfig загружает конфигурацию из переменных окружения
func LoadConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			URL: getEnv("DB_URL", "postgres://user:password@localhost:5432/jp_ru_dict?sslmode=disable"),
		},
		Server: ServerConfig{
			Port: getEnvAsInt("PORT", 8080),
		},
		Kafka: KafkaConfig{
			URL:   getEnv("KAFKA_URL", ""),
			Topic: getEnv("KAFKA_TOPIC", "dict-events"),
		},
	}
}

// getEnv получает переменную окружения или возвращает значение по умолчанию
func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt получает переменную окружения как integer или возвращает значение по умолчанию
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
