package models

// SummaryRequest represents the request body for file summarization
type SummaryRequest struct {
	FileData string `json:"file_data,omitempty"` // Base64 encoded file
	Prompt   string `json:"prompt,omitempty"`    // Custom prompt for analysis
	MimeType string `json:"mime_type,omitempty"` // Optional MIME type for the file
}

// SummaryResponse represents the response for file analysis
type SummaryResponse struct {
	Summary string `json:"summary"`
	Error   string `json:"error,omitempty"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status string `json:"status"`
}