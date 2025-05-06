package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Request structures
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ModelRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Model response structures
type Choice struct {
	Message Message `json:"message"`
}

type ModelResponse struct {
	Choices []Choice `json:"choices"`
}

// API structures
type HelpRequest struct {
	Question string `json:"question"`
	Model    string `json:"model"`
}

type HelpResponse struct {
	Answer string `json:"answer"`
}

func main() {
	http.HandleFunc("/help", handleHelp)
	fmt.Println("Virtual assistant running on port 3000!")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleHelp(w http.ResponseWriter, r *http.Request) {
	// Check HTTP method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode request body
	var reqBody HelpRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Error processing request", http.StatusBadRequest)
		return
	}

	// Use default model if not specified
	model := reqBody.Model
	if model == "" {
		model = "ai/llama3.2:1B-Q8_0"
	}

	// Prepare request for Docker Model Runner
	modelReq := ModelRequest{
		Model: model,
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are an assistant specialized in LG brand appliances.",
			},
			{
				Role:    "user",
				Content: reqBody.Question,
			},
		},
	}

	// Convert to JSON
	reqJSON, err := json.Marshal(modelReq)
	if err != nil {
		http.Error(w, "Error preparing request", http.StatusInternalServerError)
		return
	}

	// Send request to Docker Model Runner
	resp, err := http.Post(
		"http://localhost:12434/engines/llama.cpp/v1/chat/completions",
		"application/json",
		bytes.NewBuffer(reqJSON),
	)
	if err != nil {
		http.Error(w, "Error querying the model", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading model response", http.StatusInternalServerError)
		return
	}

	var modelResp ModelResponse
	if err := json.Unmarshal(body, &modelResp); err != nil {
		http.Error(w, "Error processing model response", http.StatusInternalServerError)
		return
	}

	if len(modelResp.Choices) == 0 {
		http.Error(w, "Model returned no response", http.StatusInternalServerError)
		return
	}

	response := HelpResponse{
		Answer: modelResp.Choices[0].Message.Content,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
