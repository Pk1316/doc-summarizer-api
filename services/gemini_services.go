package services

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)


type GeminiService struct {
	apiKey string
}

// NewGeminiService creates a new Gemini service
func NewGeminiService(apiKey string) *GeminiService {
	return &GeminiService{
		apiKey: "AIzaSyARYidtIcJGpcSGpn6QsqDMn6ydMJcVwVA",
	}
}

// ProcessFile sends a file to the Gemini API for processing
func (s *GeminiService) ProcessFile(fileBytes []byte, mimeType string, prompt string) (string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  "AIzaSyARYidtIcJGpcSGpn6QsqDMn6ydMJcVwVA",
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create Gemini client: %v", err)
	}
	//defer client.Close()

	parts := []*genai.Part{
		&genai.Part{
			InlineData: &genai.Blob{
				MIMEType: mimeType,
				Data:     fileBytes,
			},
		},
		genai.NewPartFromText(prompt),
	}
	contents := []*genai.Content{
	{
		Parts: parts,
		Role:  "user",
	},
}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-1.5-flash", // Using flash model for faster responses
		contents,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %v", err)
	}

	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
	return "", fmt.Errorf("no response from Gemini")
}
	summary := result.Candidates[0].Content.Parts[0].Text

	return summary, nil
}