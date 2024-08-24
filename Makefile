.PHONY: build
build:
	docker-compose build
	docker-compose run homeapi-app go mod tidy
	docker-compose run homeapi-app go mod vendor

.PHONY: serve
serve:
	docker-compose up

.PHONY: swag
swag:
	docker-compose exec homeapi-app swag init

.PHONY: generate
generate:
	docker-compose exec homeapi-app go generate ./...

.PHONY: clean
clean:
	docker rm `docker ps -a -q`
	docker rmi `docker images -q`
	docker image prune

.PHONY: in
in:
	docker-compose exec homeapi-app sh

.PHONY: tidy
tidy:
	docker-compose exec homeapi-app go mod tidy
	docker-compose exec homeapi-app go mod vendor

.PHONY: build-tidy
build-tidy:
	docker-compose run homeapi-app go mod tidy
	docker-compose run homeapi-app go mod vendor

.PHONY: force-in
force-in:
	docker-compose run homeapi-app sh

.PHONY: test
test:
	docker-compose exec homeapi-app go test ./...
