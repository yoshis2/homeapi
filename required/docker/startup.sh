#!/bin/bash

if [ ! -e "go.mod" ]; then
    go mod init
    go get -u github.com/mattn/go-sqlite3 
    go get -u github.com/labstack/echo/v4

    go install github.com/joho/godotenv/cmd/godotenv@latest
    go install github.com/golang/mock/mockgen@latest
    go install github.com/go-delve/delve/cmd/dlv@latest
    go install github.com/rubenv/sql-migrate/...@latest
    go install github.com/pilu/fresh@latest
    go install github.com/swaggo/swag/cmd/swag@latest
fi

if [ ! -e "vendor" ]; then
    go mod tidy
    go mod vendor
fi

if [ ! -e "required/logs/access.log" ]; then
    mkdir -p required/logs
    touch required/logs/access.log
    echo "created access log file"
fi

swag init
fresh -c required/docker/runner.conf

# dlv debug --headless --listen=:5050 --log --api-version=2
