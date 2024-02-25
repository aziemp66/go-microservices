FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o loggerApi ./cmd/api

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/loggerApi /app/loggerApi

ENTRYPOINT [ "/app/loggerApi" ] 