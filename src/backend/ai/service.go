package ai

import "context"

type ClassificationResult struct {
	Category  string
	Summary   string
	Tags      []string
	Embedding []float64
	Model     string
}

type Service interface {
	Classify(ctx context.Context, input string) (*ClassificationResult, error)
}
