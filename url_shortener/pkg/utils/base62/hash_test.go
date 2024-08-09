package base62

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeURL(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Test short URL",
			input: "https://example.com",
		},
		{
			name:  "Test long URL",
			input: "https://www.example.com/this/is/a/very/long/url/with/lots/of/characters/and/special?chars=%20&%3D#fragment",
		},
		{
			name:  "Test URL with query parameters",
			input: "https://example.com/search?q=golang",
		},
		{
			name:  "Test URL with fragment",
			input: "https://example.com/path#section",
		},
		{
			name:  "Test URL with port",
			input: "https://example.com:8080",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EncodeURL(tt.input)

			// Assert that the result has the correct length
			assert.Equal(t, keyLength, len(got), "Encoded URL should have length %d", keyLength)

			// Assert that the result only contains characters from the charset
			for _, char := range got {
				assert.Contains(t, charset, string(char), "Encoded URL contains invalid character %c", char)
			}
		})
	}
}
