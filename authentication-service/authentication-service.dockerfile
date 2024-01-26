FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o authApp ./cmd/api

# Path: config/Dockerfile.backend

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/authApp /app/authApp

ENTRYPOINT [ "/app/authApp" ] 