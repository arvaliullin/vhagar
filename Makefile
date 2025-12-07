.PHONY: run amqp-producer amqp-consumer database up down ps logs

run:
	- go run github.com/arvaliullin/vhagar/cmd/app

amqp-producer:
	go run github.com/arvaliullin/vhagar/examples/amqp/producer

amqp-consumer:
	go run github.com/arvaliullin/vhagar/examples/amqp/consumer

database:
	go run github.com/arvaliullin/vhagar/examples/database

up:
	docker compose -f deployments/docker-compose.yaml up --build -d

down:
	docker compose -f deployments/docker-compose.yaml down -v

ps:
	docker compose -f deployments/docker-compose.yaml ps

logs:
	docker compose -f deployments/docker-compose.yaml logs -f

.PHONY: prune
prune: down
	- docker image prune -f
	- docker container prune -f
	- docker volume prune -f
	- docker network prune -f
	- docker system prune -a --volumes -f
