package utils

import (
	"sync"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

var json jsoniter.API
var mtx = &sync.Mutex{}

func NewJsoniterInstance() jsoniter.API {
	if json == nil {
		mtx.Lock()
		defer mtx.Unlock()
		if json == nil {
			json = jsoniter.ConfigCompatibleWithStandardLibrary
		}
	}
	return json
}

func RespondWithJSON(c echo.Context, code int, payload interface{}) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(code)

	js := NewJsoniterInstance()
	return js.NewEncoder(c.Response()).Encode(payload)
}

func RespondWithError(c echo.Context, code int, message string) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(code)

	js := NewJsoniterInstance()
	return js.NewEncoder(c.Response()).Encode(map[string]string{"error": message})
}
