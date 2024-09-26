# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o users-api ./cmd/service

FROM scratch
WORKDIR /app
COPY --from=builder /app/users-api .
EXPOSE 8080
CMD ["./users-api"]