# syntax=docker/dockerfile:1

# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o url-shortener ./cmd

# Final stage with Scratch
FROM scratch
WORKDIR /app
COPY --from=builder /app/url-shortener .
CMD ["./url-shortener"]