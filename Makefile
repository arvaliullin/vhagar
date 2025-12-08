.PHONY: run
run:
	- go run github.com/arvaliullin/vhagar/cmd/app

.PHONY: amqp-producer
amqp-producer:
	go run github.com/arvaliullin/vhagar/examples/amqp/producer

.PHONY: amqp-consumer
amqp-consumer:
	go run github.com/arvaliullin/vhagar/examples/amqp/consumer

.PHONY: database
database:
	go run github.com/arvaliullin/vhagar/examples/database

.PHONY: up
up:
	docker compose -f deployments/docker-compose.yaml up --build -d

.PHONY: down
down:
	docker compose -f deployments/docker-compose.yaml down -v

.PHONY: ps
ps:
	docker compose -f deployments/docker-compose.yaml ps

.PHONY: logs
logs:
	docker compose -f deployments/docker-compose.yaml logs -f

.PHONY: prune
prune: down
	- docker image prune -f
	- docker container prune -f
	- docker volume prune -f
	- docker network prune -f
	- docker system prune -a --volumes -f

.PHONY: generate-mocks
generate-mocks:
	go generate ./...

.PHONY: fmt
fmt:
	- go fmt ./...

.PHONY: test
test:
	- go test ./...

.PHONY: build-devimage
build-devimage:
	mkdir -p bin
	docker build -f .devcontainer/Dockerfile -t vhagar-devimage .
	docker save -o bin/devimage.tar vhagar-devimage

.PHONY: archive-source
archive-source:
	mkdir -p bin
	git archive --format=tar.gz --output=bin/source.tar.gz HEAD
