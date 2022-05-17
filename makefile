start:
	@docker-compose --env-file ./.env up

build:
	@docker-compose --env-file ./.env build

dev:
	@docker-compose --env-file ./.env -f compose.dev.yaml up --build

run:
	@docker exec -i --tty http-api-1 /bin/sh

test:
	@docker exec -i --tty http-api-1 curl --request GET --url http://localhost:3000
