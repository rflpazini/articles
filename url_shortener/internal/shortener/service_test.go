package shortener

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockRepository struct {
	*Repository // Embedding the Repository struct
	mock.Mock
}

func (m *MockRepository) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) Set(ctx context.Context, u *URLInfo) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *MockRepository) GetAll(ctx context.Context) (map[string]string, error) {
	args := m.Called(ctx)
	return args.Get(0).(map[string]string), args.Error(1)
}

func TestService_Upsert(t *testing.T) {
	tests := []struct {
		name           string
		input          *URLInfo
		mockGetResult  string
		mockGetError   error
		mockSetError   error
		expectedStatus int
	}{
		{
			name: "successful upsert",
			input: &URLInfo{
				Url: "https://example.com",
			},
			mockGetResult:  "",
			mockGetError:   nil,
			mockSetError:   nil,
			expectedStatus: http.StatusCreated,
		},
		{
			name: "upsert with existing URL",
			input: &URLInfo{
				Url: "https://example.com",
			},
			mockGetResult:  `{"Url": "https://example.com"}`,
			mockGetError:   nil,
			mockSetError:   nil,
			expectedStatus: http.StatusCreated,
		},
		{
			name: "upsert with repository set error",
			input: &URLInfo{
				Url: "https://example.com",
			},
			mockGetResult:  "",
			mockGetError:   nil,
			mockSetError:   errors.New("redis error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mockRepo := new(MockRepository)
			mockRepo.On("Get", mock.Anything, mock.Anything).Return(tt.mockGetResult, tt.mockGetError)
			mockRepo.On("Set", mock.Anything, mock.Anything).Return(tt.mockSetError)

			jsonAPI := jsoniter.ConfigCompatibleWithStandardLibrary
			service := Service{
				Repository: mockRepo,
				Json:       jsonAPI,
			}

			err := service.Upsert(c)
			if tt.mockSetError != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedStatus, rec.Code)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_Get(t *testing.T) {
	tests := []struct {
		name           string
		urlQuery       string
		mockGetResult  string
		mockGetError   error
		expectedStatus int
		expectedURL    string
	}{
		{
			name:           "successful get",
			urlQuery:       "abc123",
			mockGetResult:  `{"Url": "https://example.com"}`,
			mockGetError:   nil,
			expectedStatus: http.StatusFound,
			expectedURL:    "https://example.com",
		},
		{
			name:           "URL not found",
			urlQuery:       "nonexistent",
			mockGetResult:  "",
			mockGetError:   redis.Nil,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "repository error",
			urlQuery:       "error",
			mockGetResult:  "",
			mockGetError:   errors.New("redis error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/v1/shortener?url="+tt.urlQuery, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mockRepo := new(MockRepository)
			mockRepo.On("Get", mock.Anything, tt.urlQuery).Return(tt.mockGetResult, tt.mockGetError)

			jsonAPI := jsoniter.ConfigCompatibleWithStandardLibrary
			service := Service{
				Repository: mockRepo,
				Json:       jsonAPI,
			}

			err := service.Get(c)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedURL != "" {
				assert.Equal(t, tt.expectedURL, rec.Header().Get("Location"))
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestService_GetAll(t *testing.T) {
	tests := []struct {
		name            string
		mockGetAllData  map[string]string
		mockGetAllError error
		expectedStatus  int
	}{
		{
			name: "successful get all",
			mockGetAllData: map[string]string{
				"abc123": `{"Url": "https://example.com"}`,
				"def456": `{"Url": "https://anotherexample.com"}`,
			},
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/v1/shortener", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mockRepo := new(MockRepository)
			mockRepo.On("GetAll", mock.Anything).Return(tt.mockGetAllData, tt.mockGetAllError)

			jsonAPI := jsoniter.ConfigCompatibleWithStandardLibrary
			service := Service{
				Repository: mockRepo,
				Json:       jsonAPI,
			}

			err := service.Get(c)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, rec.Code)

			mockRepo.AssertExpectations(t)
		})
	}
}
