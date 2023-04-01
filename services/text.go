package services

import (
	"sync"

	"github.com/google/uuid"
	"github.com/szpnygo/VecTextSearch/config"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate/entities/models"
)

var weaviateClient *weaviate.Client
var once sync.Once

func initWeaviateClient(config *config.AppConfig) {
	cfg := weaviate.Config{
		Host:   config.WeaviateURL,
		Scheme: "http",
	}
	weaviateClient = weaviate.New(cfg)
}

func AddText(appConfig *config.AppConfig, name, content string) (string, error) {
	once.Do(func() { initWeaviateClient(appConfig) })
	embedding, err := getEmbedding(content, appConfig.OpenAIKey)
	if err != nil {
		return "", err
	}

	float32Embedding := make([]float32, len(embedding))
	for i, v := range embedding {
		float32Embedding[i] = float32(v)
	}

	id := uuid.New().String()
	dataSchema := map[string]interface{}{
		"name":    name,
		"content": content,
	}

	err = addVector(weaviateClient, id, dataSchema, float32Embedding)
	if err != nil {
		return "", err
	}

	return id, nil
}

func SearchSimilarTexts(appConfig *config.AppConfig, content string) (*models.GraphQLResponse, error) {
	once.Do(func() { initWeaviateClient(appConfig) })
	embedding, err := getEmbedding(content, appConfig.OpenAIKey)
	if err != nil {
		return nil, err
	}

	float32Embedding := make([]float32, len(embedding))
	for i, v := range embedding {
		float32Embedding[i] = float32(v)
	}

	response, err := searchVectors(weaviateClient, float32Embedding)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func FindExactText(appConfig *config.AppConfig, content string) (*models.GraphQLResponse, error) {
	once.Do(func() { initWeaviateClient(appConfig) })
	response, err := findTextByContent(weaviateClient, content)
	if err != nil {
		return nil, err
	}

	return response, nil
}
