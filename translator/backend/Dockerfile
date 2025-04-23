FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o translator-service ./cmd

FROM alpine:3.21
RUN apk add --no-cache  ca-certificates

WORKDIR /root/

COPY --from=builder /app/translator-service .

EXPOSE 3000

CMD ["./translator-service"]