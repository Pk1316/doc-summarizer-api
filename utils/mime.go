package utils

import (
	"mime"
	"path/filepath"
)

// DetectMimeType determines the MIME type of a file based on its extension
func DetectMimeType(filename string) string {
	ext := filepath.Ext(filename)
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		return "application/octet-stream"
	}
	return mimeType
}
