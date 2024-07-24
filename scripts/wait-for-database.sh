#!/bin/sh

# O shell irá encerrar a execução do script quando um comando falhar
set -e

# Verifica se as variáveis de ambiente estão definidas
if [ -z "$POSTGRES_HOST" ] || [ -z "$POSTGRES_PORT" ]; then
  echo "❌ POSTGRES_HOST ou POSTGRES_PORT não estão definidas"
  exit 1
fi

while ! nc -z $POSTGRES_HOST $POSTGRES_PORT; do
  echo "🟡 Waiting for Postgres Database Startup ($POSTGRES_HOST:$POSTGRES_PORT) ..."
  sleep 5
done

echo "✅ Postgres Database Started Successfully ($POSTGRES_HOST:$POSTGRES_PORT)"

exec "$@"
