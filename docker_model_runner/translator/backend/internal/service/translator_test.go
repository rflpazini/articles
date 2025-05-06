package service_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rflpazini/articles/translator/internal/config"
	"github.com/rflpazini/articles/translator/internal/model"
	"github.com/rflpazini/articles/translator/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestTranslatorService_Translate(t *testing.T) {
	tests := []struct {
		name           string
		inputText      string
		serverResponse model.Response
		serverStatus   int
		serverError    bool
		expected       string
		expectError    bool
	}{
		{
			name:      "successful translation",
			inputText: "The party of the first part shall hereafter be known as the party of the first part.",
			serverResponse: model.Response{
				Choices: []model.Choice{
					{
						Message: model.Message{
							Content: "A primeira parte será conhecida como a primeira parte.",
						},
					},
				},
			},
			serverStatus: http.StatusOK,
			expected:     "A primeira parte será conhecida como a primeira parte.",
			expectError:  false,
		},
		{
			name:         "server error",
			inputText:    "Some legal text",
			serverStatus: http.StatusInternalServerError,
			expectError:  true,
		},
		{
			name:        "timeout error",
			inputText:   "Some legal text",
			serverError: true,
			expectError: true,
		},
		{
			name:      "empty response choices",
			inputText: "Some legal text",
			serverResponse: model.Response{
				Choices: []model.Choice{},
			},
			serverStatus: http.StatusOK,
			expectError:  true,
		},
		{
			name:         "malformed response",
			inputText:    "Some legal text",
			serverStatus: http.StatusOK,
			expectError:  true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "POST", r.Method)
				assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

				if tc.serverError {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				w.WriteHeader(tc.serverStatus)

				if tc.serverStatus == http.StatusOK {
					respBody, _ := json.Marshal(tc.serverResponse)
					_, err := w.Write(respBody)
					assert.NoError(t, err)
				} else {
					_, err := w.Write([]byte("Error from translation service"))
					assert.NoError(t, err)
				}
			}))
			defer mockServer.Close()

			cfg := &config.Config{
				ModelName:     "test-model",
				ModelEndpoint: mockServer.URL,
			}

			translatorService := service.NewTranslatorService(cfg)
			result, err := translatorService.Translate(tc.inputText)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
