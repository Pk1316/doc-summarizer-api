package config

import "os"


type Config struct {
	APIKey string
	Port   string
}

func LoadConfig() *Config {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {

		apiKey = os.Getenv("GEMINI_API_KEY")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		APIKey: apiKey,
		Port:   port,
	}
}
