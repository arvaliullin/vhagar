.PHONY: run
run:
	- CGO_ENABLED=0 GOOS=linux GOPROXY=off go run -mod=readonly github.com/arvaliullin/vhagar/cmd/app

.PHONY: amqp-producer
amqp-producer:
	CGO_ENABLED=0 GOOS=linux GOPROXY=off go run -mod=readonly github.com/arvaliullin/vhagar/examples/amqp/producer

.PHONY: amqp-consumer
amqp-consumer:
	CGO_ENABLED=0 GOOS=linux GOPROXY=off go run -mod=readonly github.com/arvaliullin/vhagar/examples/amqp/consumer

.PHONY: database
database:
	DATABASE_URL="postgres://postgres:postgres@host.docker.internal:5432/postgres?sslmode=disable" \
	CGO_ENABLED=0 GOOS=linux GOPROXY=off go run -mod=readonly github.com/arvaliullin/vhagar/examples/database

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

.PHONY: lint
lint:
	- golangci-lint run ./...

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
