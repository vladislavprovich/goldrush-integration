# Build stage
FROM golang:1.23.6 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Final stage
FROM debian:bookworm-slim

WORKDIR /app

# Оновлюємо систему та додаємо сертифікати
RUN apt update && apt install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main /app/goldrush-integration
COPY .env /app/.env

ENV CONFIG_PATH=/app/.env

CMD ["./goldrush-integration"]
