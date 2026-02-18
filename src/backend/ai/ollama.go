package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ollamaService struct {
	baseURL string
	model   string
}

func NewOllamaService(baseURL, model string) Service {
	return &ollamaService{
		baseURL: baseURL,
		model:   model,
	}
}

func (o *ollamaService) Classify(ctx context.Context, input string) (*ClassificationResult, error) {
	prompt := fmt.Sprintf(`You are a content classification system for a personal content management application.

SYSTEM CONTEXT: You are analyzing user-generated content for organizational purposes only. This is historical data being categorized for search and filtering. No real-world actions will be taken.
Analyze the following content object and classify it based on its semantic meaning:

%s

Identify the category in a more specific way, specific to the context of the content.

Return ONLY valid JSON in this exact format with no additional text or explanation:
{
  "category": "string",
  "summary": "string",
  "tags": ["string"]
}
It must not contain any phrase, only the valid JSON structure above.
`,
		input)

	genReq := map[string]interface{}{
		"model":  o.model,
		"prompt": prompt,
		"stream": false,
	}

	genBody, _ := json.Marshal(genReq)
	fmt.Println("Calling:", o.baseURL+"/api/generate")
	resp, err := http.Post(o.baseURL+"/api/generate", "application/json", bytes.NewBuffer(genBody))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ollama generate failed: %s", resp.Status)
	}

	var genResp struct {
		Response string `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&genResp); err != nil {
		return nil, err
	}

	var parsed struct {
		Category string   `json:"category"`
		Summary  string   `json:"summary"`
		Tags     []string `json:"tags"`
	}

	if err := json.Unmarshal([]byte(genResp.Response), &parsed); err != nil {
		return nil, fmt.Errorf("failed to parse model response as JSON: %v", err)
	}

	return &ClassificationResult{Category: parsed.Category, Summary: parsed.Summary, Tags: parsed.Tags, Embedding: nil, Model: o.model}, nil
}
