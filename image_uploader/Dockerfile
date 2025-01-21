# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/image-uploader

FROM alpine:3.21
WORKDIR /app

COPY --from=builder /app/server /app/server
RUN mkdir -p /app/uploads

EXPOSE 8080
CMD ["/app/server"]