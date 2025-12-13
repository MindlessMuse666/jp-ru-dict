#!/bin/sh

set -e

echo "Ожидание PostgreSQL..."

# Ждем, пока PostgreSQL будет готов принимать подключения
until PGPASSWORD=password psql -h postgres -U app_user -d jp_ru_dict -c '\q' 2>/dev/null; do
  echo "PostgreSQL еще не доступен, ждем..."
  sleep 2
done

echo "PostgreSQL доступен"

echo "Выполнение миграций..."
# Запускаем миграции
./main --migrate

echo "Запуск сервера..."
# Запускаем приложение
exec ./main
