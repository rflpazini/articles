# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o smoke-test  ./cmd

FROM scratch
WORKDIR /app

COPY --from=builder /app/smoke-test .
CMD ["./smoke-test"]