FROM registry.astralinux.ru/library/astra/ubi18-golang121:1.8.4 AS builder

RUN apt-get update && apt-get install -y --no-install-recommends \
    golang-github-streadway-amqp-dev \
    && rm -rf /var/lib/apt/lists/*

ENV GOPATH=/usr/share/gocode

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOPROXY=off go build -mod=readonly -o /app/consumer github.com/arvaliullin/vhagar/examples/amqp/consumer

FROM registry.astralinux.ru/library/astra/ubi18:1.8.4

WORKDIR /app

COPY --from=builder /app/consumer /app/consumer

CMD ["/app/consumer"]
