package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rflpazini/articles/translator/internal/config"
	"github.com/rflpazini/articles/translator/internal/model"
)

type TranslatorService struct {
	config *config.Config
	client *http.Client
}

func NewTranslatorService(cfg *config.Config) *TranslatorService {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &TranslatorService{
		config: cfg,
		client: client,
	}
}

func (s *TranslatorService) Translate(text string) (string, error) {
	modelReq := model.Request{
		Model: s.config.ModelName,
		Messages: []model.Message{
			{
				Role:    "system",
				Content: "You are an expert in translating legal language into plain and simple Portuguese. Always provide direct, human-like translations without any AI-like responses.",
			},
			{
				Role: "user",
				Content: fmt.Sprintf("Translate this legal text to simple, everyday Portuguese that anyone can UNDERSTAND: %s ."+
					"Provide ONLY the translation without any introductory phrases, explanations of what you're doing, or conclusions without using double quotes. Be simple and HUMAN, also try to use daily language to be simple as possible", text),
			},
		},
	}

	reqJSON, err := json.Marshal(modelReq)
	if err != nil {
		return "", fmt.Errorf("error serializing request: %w", err)
	}

	resp, err := s.client.Post(
		s.config.ModelEndpoint,
		"application/json",
		bytes.NewBuffer(reqJSON),
	)
	if err != nil {
		return "", fmt.Errorf("error communicating with the model: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("model service error (code %d): %s", resp.StatusCode, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading model response: %w", err)
	}

	var modelResp model.Response
	if err := json.Unmarshal(body, &modelResp); err != nil {
		return "", fmt.Errorf("error processing model response: %w", err)
	}

	if len(modelResp.Choices) == 0 {
		return "", fmt.Errorf("model returned no responses")
	}

	return modelResp.Choices[0].Message.Content, nil
}
