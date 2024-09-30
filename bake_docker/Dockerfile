# syntax=docker/dockerfile:1
FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main .

FROM scratch
WORKDIR /app

COPY --from=builder /app/main .

ARG HELLO_MSG="Hello, Docker!"
ENV HELLO_MSG=${HELLO_MSG}

CMD ["./main"]



