FROM golang:alpine
ENV GOBIN /go/bin

# フォルダ構成をGoPathの通りに設置
WORKDIR /go/src/homeapi

RUN apk update && apk upgrade
RUN apk add --virtual --no-cache bash gcc musl-dev git

RUN go get -u github.com/joho/godotenv && \
    go get -u github.com/mattn/go-sqlite3  && \
    go get -u github.com/derekparker/delve/cmd/dlv
RUN go get -u bitbucket.org/liamstask/goose/cmd/goose
RUN go get -u github.com/pilu/fresh 

RUN go mod init && go mod tidy && \
    go get -u github.com/labstack/echo/v4 && \
    go get -u github.com/swaggo/swag/cmd/swag

CMD swag init && \
    # dlv debug --headless --listen=:5050 --log --api-version=2
    fresh -c required/docker/runner.conf
