package model

type TranslateRequest struct {
	Text string `json:"text"`
}

type TranslateResponse struct {
	Translation string `json:"translation"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Response struct {
	Choices []Choice `json:"choices"`
}
