FROM golang:1.22-alpine
ENV GOBIN /go/bin

# フォルダ構成をGoPathの通りに設置
WORKDIR /go/src/homeapi

RUN apk update && apk upgrade && \
    apk add bash gcc musl-dev git

RUN go mod init

RUN go get -u github.com/mattn/go-sqlite3 && \
    go get -v github.com/rubenv/sql-migrate/... && \
    go get -u github.com/labstack/echo/v4
    
RUN go install github.com/joho/godotenv/cmd/godotenv@latest && \
    go install github.com/golang/mock/mockgen@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/pilu/fresh@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest

RUN go mod vendor
