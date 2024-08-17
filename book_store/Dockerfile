# syntax=docker/dockerfile:1
# check=skip=none;error=true

FROM golang:1.23-alpine AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0  \
    GOOS=linux  \
    go build -a -installsuffix cgo -o bookstore-api /app/cmd/bookstore/main.go

FROM scratch
WORKDIR /app

ENV ENV=production

COPY --from=build /app/bookstore-api .
COPY --from=build /app/config/dev.json /config/production.json
EXPOSE 8080
CMD ["./bookstore-api"]

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
            CMD curl --fail http://localhost:8080/healthcheck/ping || exit