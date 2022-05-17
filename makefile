start:
	@docker-compose --env-file ./.env up

build:
	@docker-compose --env-file ./.env build

dev:
	docker
