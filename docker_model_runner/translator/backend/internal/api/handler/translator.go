package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rflpazini/articles/translator/internal/model"
	"github.com/rflpazini/articles/translator/internal/service"
)

type TranslatorHandler struct {
	service *service.TranslatorService
}

func NewTranslatorHandler(service *service.TranslatorService) *TranslatorHandler {
	return &TranslatorHandler{
		service: service,
	}
}

func (h *TranslatorHandler) Translate(c echo.Context) error {
	var req model.TranslateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	if req.Text == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "The 'text' field is required",
		})
	}

	if len(req.Text) > 5000 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Text too long (max 5000 characters)",
		})
	}

	translation, err := h.service.Translate(req.Text)
	if err != nil {
		c.Logger().Errorf("Translation error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error processing translation",
		})
	}

	return c.JSON(http.StatusOK, model.TranslateResponse{
		Translation: translation,
	})
}

func (h *TranslatorHandler) Health(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
