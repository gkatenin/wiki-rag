package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// RequestPayload matches the JSON structure expected by the LLaMA server
type RequestPayload struct {
	Prompt   string `json:"prompt"`
	NPredict int    `json:"n_predict"`
}

// ResponsePayload matches the JSON response from the server
type ResponsePayload struct {
	Content string `json:"content"`
}

// ReaderLLM sends prompt to and receives response from a llama-server running
// LLaMA-based language model.
func ReaderLLM(p string, npred int) (string, error) {
	// Running llama-server locally in this scenario.
	const url = "http://localhost:8080/completion"

	payload := RequestPayload{
		Prompt:   "<|im_start|>user\n" + p + "<|im_end|>\n<|im_start|>assistant\n",
		NPredict: npred,
	}

	// Encode the payload into JSON.
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Create http client with a 2m timeout.
	client := http.Client{Timeout: 2 * time.Minute}

	// Send payload via POST request.
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read generated response.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Handle status codes from the server.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("server returned error: %s", string(body))
	}

	// Decode response payload as JSON.
	var response ResponsePayload
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}
	return response.Content, nil
}
