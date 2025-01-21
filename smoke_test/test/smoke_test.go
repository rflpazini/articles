package tests

import (
	"net/http"
	"os"
	"testing"
)

type endpointTest struct {
	name           string
	path           string
	expectedStatus int
}

func TestEndpoints(t *testing.T) {
	baseURL := os.Getenv("TARGET_URL")
	if baseURL == "" {
		t.Fatal("TARGET_URL não configurada")
	}

	tests := []endpointTest{
		{
			name:           "Health Endpoint",
			path:           "/health",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Catalog Endpoint",
			path:           "/catalog",
			expectedStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			url := baseURL + test.path
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("Erro ao conectar ao endpoint %s: %v", test.name, err)
			}
			if resp.StatusCode != test.expectedStatus {
				t.Errorf("%s: Status esperado: %d, obtido: %d", test.name, test.expectedStatus, resp.StatusCode)
			}
		})
	}
}
