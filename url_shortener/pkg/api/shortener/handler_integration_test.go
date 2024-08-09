//go:build integration

package shortener

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	tcRedis "github.com/testcontainers/testcontainers-go/modules/redis"
)

func TestIntegration_RegisterRoutes(t *testing.T) {
	ctx := context.Background()

	redisContainer, err := tcRedis.Run(ctx,
		"redis:alpine",
		tcRedis.WithSnapshotting(10, 1),
		tcRedis.WithLogLevel(tcRedis.LogLevelVerbose),
	)

	if err != nil {
		log.Fatalf("failed to start container: %s", err)
	}

	// Clean up the container
	defer func() {
		if err := redisContainer.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	// Set up REDIS client pointing to TestContainer
	endpoint, _ := redisContainer.Host(ctx)
	port, _ := redisContainer.MappedPort(ctx, "6379")

	client := redis.NewClient(&redis.Options{
		Addr:     endpoint + ":" + port.Port(),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Set up Echo and routes
	e := echo.New()
	jsonAPI := jsoniter.ConfigCompatibleWithStandardLibrary
	RegisterRoutes(e, client, jsonAPI)

	tests := []struct {
		name           string
		method         string
		url            string
		body           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "POST /v1/shortener - Create short URL",
			method:         http.MethodPost,
			url:            "/v1/shortener",
			body:           `{"url":"https://example.com"}`,
			expectedStatus: http.StatusCreated,
			expectedBody:   `"url":"https://example.com"`,
		},
		{
			name:           "GET /v1/shortener - Retrieve short URL",
			method:         http.MethodGet,
			url:            "/v1/shortener?url=someShortURL",
			expectedStatus: http.StatusNotFound,
			expectedBody:   "", // add expected body if needed
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.method == http.MethodPost {
				req = httptest.NewRequest(tt.method, tt.url, bytes.NewBufferString(tt.body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			} else {
				req = httptest.NewRequest(tt.method, tt.url, nil)
			}

			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedBody != "" {
				assert.Contains(t, rec.Body.String(), tt.expectedBody)
			}
		})
	}
}
