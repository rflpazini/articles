package monitoring

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests processed, labeled by path and method.",
		},
		[]string{"path", "method"},
	)

	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
)

func Init() {
	prometheus.MustRegister(HttpRequestsTotal, HttpRequestDuration)
}

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()

		// Incrementa o contador de requisições
		HttpRequestsTotal.WithLabelValues(c.Path(), c.Request().Method).Inc()

		// Executa o handler original
		err := next(c)

		// Calcula a duração e a grava no histograma
		duration := time.Since(start).Seconds()
		HttpRequestDuration.WithLabelValues(c.Path(), c.Request().Method).Observe(duration)

		return err
	}
}

func Route(e *echo.Echo) {
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}
