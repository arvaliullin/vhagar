FROM registry.astralinux.ru/library/astra/ubi18-golang121:1.8.4 AS builder

RUN apt-get update && apt-get install -y --no-install-recommends \
    golang-github-go-chi-chi-dev \
    && rm -rf /var/lib/apt/lists/*

ENV GOPATH=/usr/share/gocode

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOPROXY=off go build -mod=readonly -o /app/http github.com/arvaliullin/vhagar/examples/http

FROM registry.astralinux.ru/library/astra/ubi18:1.8.4

RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/http /app/http

EXPOSE 8080

CMD ["/app/http"]

