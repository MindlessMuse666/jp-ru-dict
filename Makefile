.PHONY: run build down logs migrate test clean reset-db help frontend-install frontend-dev frontend-build logs-frontend 

.DEFAULT_GOAL := help

# Запуск всех сервисов
run:
	docker-compose up --build

# Запуск в фоновом режиме
up:
	docker-compose up --build -d

# Сборка образов
build:
	docker-compose build

# Остановка всех сервисов
down:
	docker-compose down

# Остановка и удаление томов
down-volumes:
	docker-compose down -v

# Просмотр логов
logs:
	docker-compose logs -f

# Выполнение миграций
migrate:
	docker-compose run --rm backend ./main --migrate

# Пересоздание базы данных
reset-db:
	docker-compose down -v
	docker volume prune -f
	docker-compose up -d postgres
	sleep 5
	docker-compose run --rm backend ./main --migrate
	echo "База данных пересоздана"

# Очистка
clean:
	docker-compose down -v
	docker system prune -f --volumes

# Проверка статуса сервисов
status:
	docker-compose ps

# Подключение к PostgreSQL через psql
psql:
	docker-compose exec postgres psql -U app_user -d jp_ru_dict

# Открыть pgAdmin
pgadmin:
	open http://localhost:5050

# Установка зависимостей фронтенда
frontend-install:
	cd frontend && npm install

# Запуск фронтенда в режиме разработки
frontend-dev:
	cd frontend && npm run dev

# Сборка фронтенда
frontend-build:
	cd frontend && npm run build

# Просмотр логов фронтенда
logs-frontend:
	docker-compose logs -f frontend

# Справка
help:
	@echo "Доступные команды:"
	@echo "  make run           - Запуск всех сервисов"
	@echo "  make up            - Запуск в фоновом режиме"
	@echo "  make build         - Сборка образов"
	@echo "  make down          - Остановка сервисов"
	@echo "  make logs          - Просмотр логов"
	@echo "  make migrate       - Выполнить миграции"
	@echo "  make reset-db      - Полное пересоздание базы данных"
	@echo "  make clean         - Полная очистка"
	@echo "  make status        - Статус сервисов"
	@echo "  make psql          - Подключение к PostgreSQL"
	@echo "  make pgadmin       - Открыть pgAdmin"
	@echo "  make frontend-install - Установка зависимостей фронтенда"
	@echo "  make frontend-dev     - Запуск фронтенда в режиме разработки"
	@echo "  make frontend-build   - Сборка фронтенда"
	@echo "  make logs-frontend    - Просмотр логов фронтенда"
	@echo "  make help          - Показать эту справку"