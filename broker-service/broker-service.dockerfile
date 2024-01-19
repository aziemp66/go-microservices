FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# Path: config/Dockerfile.backend

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/brokerApp /app/brokerApp

ENTRYPOINT [ "/app/brokerApp" ] 