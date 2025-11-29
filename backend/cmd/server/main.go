package main

import (
	"database/sql"
	"log"

	"jp-ru-dict/internal/config"
	// "jp-ru-dict/internal/repository"
	// "jp-ru-dict/internal/handlers"
	// "jp-ru-dict/internal/kafka"
	// "jp-ru-dict/internal/service"

	_ "github.com/lib/pq"
)

// main является точкой входа в приложение
// Инициализирует все зависимости и запускает HTTP сервер
func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Подключение к базе данных
	db, err := sql.Open("postgres", cfg.Database.URL)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}
	defer db.Close()

	// Проверка подключения к БД
	if err := db.Ping(); err != nil {
		log.Fatal("Ошибка проверки подключения к базе данных:", err)
	}
	log.Println("Успешное подключение к базе данных")

	// TODO: Инициализация Kafka продюсера

	// TODO: Инициализация слоев приложения

	// TODO: Запуск сервера
	// log.Printf("Сервер запускается на порту :%d", cfg.Server.Port)

}
