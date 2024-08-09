package base62

import (
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
	"net/url"
)

const (
	charset   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	keyLength = 6
)

func EncodeURL(value string) string {
	// Parse the URL (if you need to use components of it)
	parsedURL, _ := url.Parse(value)

	urlHash := sha256.Sum256([]byte(parsedURL.String()))

	// Use the base62 to seed the random number generator
	seed := binary.LittleEndian.Uint64(urlHash[:8])
	seededRand := rand.New(rand.NewSource(int64(seed)))

	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(shortKey)
}
