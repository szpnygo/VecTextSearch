package config

import (
	"errors"
	"os"
	"strconv"
)

// AppConfig contains the application's configuration.
type AppConfig struct {
	APIPort               int
	WeaviateURL           string
	OpenAIKey             string
	AllowDuplicateContent bool
}

// LoadConfig loads the configuration from environment variables.
func LoadConfig() (*AppConfig, error) {
	apiPort, err := strconv.Atoi(os.Getenv("VECTEXTSEARCH_API_PORT"))
	if err != nil {
		return nil, err
	}

	weaviateURL := os.Getenv("VECTEXTSEARCH_WEAVIATE_URL")
	if weaviateURL == "" {
		return nil, errors.New("VECTEXTSEARCH_WEAVIATE_URL not set")
	}

	openAIKey := os.Getenv("VECTEXTSEARCH_OPENAI_KEY")
	if openAIKey == "" {
		return nil, errors.New("VECTEXTSEARCH_OPENAI_KEY not set")
	}

	allowDuplicateContent := true // default value
	if os.Getenv("VECTEXTSEARCH_ALLOW_DUPLICATE_CONTENT") == "false" {
		allowDuplicateContent = false
	}

	return &AppConfig{
		APIPort:               apiPort,
		WeaviateURL:           weaviateURL,
		OpenAIKey:             openAIKey,
		AllowDuplicateContent: allowDuplicateContent,
	}, nil
}
