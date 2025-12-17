.PHONY: swagger run build

# Генерация Swagger документации
swagger:
	swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal

# Запуск сервера
run:
	go run cmd/server/main.go

# Сборка
build:
	go build -o bin/server cmd/server/main.go
