build:
	docker-compose build
serve:
	docker-compose up
swag:
	docker-compose exec homeapi-app swag init
clean:
	docker rm `docker ps -a -q`
	docker rmi `docker images -q`
in:
	docker-compose exec homeapi-app  sh
