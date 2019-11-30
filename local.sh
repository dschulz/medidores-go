#!/usr/bin/env bash



export GO111MODULE=on
go build -o migrate ./cmd/migrate
go build -o medidores ./cmd/app


export DEBUG=true
export SERVER_PORT=3000
export SERVER_TIMEOUT_READ=5s
export SERVER_TIMEOUT_WRITE=10s
export SERVER_TIMEOUT_IDLE=15s

# Cambiar esto para apuntar a una base de datos local
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=medidores
export DB_PASS=medidores
export DB_NAME=medidores
export DB_SSLMODE=disable

echo "Migrate down"
./migrate down
echo "Migrate up"
./migrate up

./medidores



