
ARG BUILD_IMAGE=registry.astralinux.ru/library/astra/ubi18-golang121:1.8.4
ARG RUNTIME_IMAGE=registry.astralinux.ru/library/astra/ubi18:1.8.4

FROM ${BUILD_IMAGE} AS builder

RUN apt-get update && apt-get install -y --no-install-recommends \
    golang-github-streadway-amqp-dev \
    && rm -rf /var/lib/apt/lists/*

ENV GOPATH=/usr/share/gocode

WORKDIR /app

COPY go.mod go.sum* ./
COPY examples/amqp/producer/ ./examples/amqp/producer/

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/producer ./examples/amqp/producer

FROM ${RUNTIME_IMAGE}

WORKDIR /app

COPY --from=builder /app/producer /app/producer

CMD ["/app/producer"]
