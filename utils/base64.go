package utils

import (
	"encoding/base64"
	"strings"
)

// ParseBase64Data decodes a base64 encoded string and extracts the MIME type if present
func ParseBase64Data(base64Data string) ([]byte, string, error) {
	var mimeType string
	data := base64Data

	if strings.Contains(base64Data, ";base64,") {
		parts := strings.Split(base64Data, ";base64,")
		if len(parts) > 1 && strings.HasPrefix(parts[0], "data:") {
			mimeType = strings.TrimPrefix(parts[0], "data:")
		}
		data = parts[1]
	}

	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, "", err
	}

	return decoded, mimeType, nil
}
