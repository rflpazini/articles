FROM golang:1.25-alpine AS build
WORKDIR /src

RUN apk --no-cache add ca-certificates git

COPY go.mod ./
COPY go.su[m] ./ 
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

ARG VERSION=dev
ARG COMMIT=unknown
ARG BUILD_TIME=unknown

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build \
    -ldflags="-s -w -X main.Version=${VERSION} -X main.CommitHash=${COMMIT} -X main.BuildTime=${BUILD_TIME}" \
    -trimpath \
    -o /out/app ./main.go

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /out/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]