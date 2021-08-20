FROM golang:alpine
ENV GOBIN /go/bin

# フォルダ構成をGoPathの通りに設置
WORKDIR /go/src/homeapi

RUN apk update && apk upgrade
RUN apk add bash gcc musl-dev git

RUN go get -u github.com/joho/godotenv && \
    go get -u github.com/mattn/go-sqlite3  && \
    go get -u github.com/go-delve/delve/cmd/dlv &&\
    go get -u github.com/pilu/fresh
RUN go get -u bitbucket.org/liamstask/goose/cmd/goose
RUN go get -u github.com/golang/mock/mockgen

RUN go mod init && go mod tidy && \
    go get -u github.com/labstack/echo/v4 && \
    go get -u github.com/swaggo/swag/cmd/swag
