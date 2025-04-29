#!/bin/bash
set -e

until getent hosts db; do
  sleep 2
done

until pg_isready -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME"; do
  sleep 2
done

make migrate

echo "Starting application..."
exec "$@"