#!/bin/sh

set -e

echo "Запуск приложения..."

# Ожидаем доступности PostgreSQL
echo "Ожидание PostgreSQL..."
while ! nc -z postgres 5432; do
  sleep 0.1
done
echo "PostgreSQL доступен"

# Ожидаем доступности Kafka
echo "Ожидание Kafka..."
while ! nc -z kafka 9092; do
  sleep 0.1
done
echo "Kafka доступен"

# Выполняем миграции, если указан флаг
if [ "$1" = "--migrate" ]; then
  echo "Выполнение миграций..."
  ./main --migrate
  exit 0
fi

# Запускаем сервер
echo "Запуск сервера..."
exec ./main