package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/szpnygo/VecTextSearch/config"
	"github.com/weaviate/weaviate-go-client/v4/weaviate"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"
)

var appConfig *config.AppConfig

type OpenAIResponse struct {
	Data  []OpenAIData `json:"data"`
	Model string       `json:"model"`
}

type OpenAIData struct {
	Embedding []float64 `json:"embedding"`
}

func getEmbedding(text, apiKey string) ([]float64, error) {
	url := "https://api.openai.com/v1/embeddings"
	model := "text-embedding-ada-002"

	requestBody, err := json.Marshal(map[string]string{
		"input": text,
		"model": model,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("OpenAI API returned an error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var openAIResponse OpenAIResponse
	err = json.Unmarshal(body, &openAIResponse)
	if err != nil {
		return nil, err
	}

	if len(openAIResponse.Data) == 0 {
		return nil, errors.New("OpenAI API returned no data")
	}

	return openAIResponse.Data[0].Embedding, nil
}

// 定义一个全局变量用于存储 Weaviate 客户端实例
var weaviateClient *weaviate.Client

// 初始化 Weaviate 客户端
func initWeaviateClient(config *config.AppConfig) {
	cfg := weaviate.Config{
		Host:   "127.0.0.1:8888",
		Scheme: "http",
	}
	weaviateClient = weaviate.New(cfg)
}

// 添加一个向量到 Weaviate 数据库
func addVector(id string, dataSchema map[string]interface{}, vector []float32) error {
	_, err := weaviateClient.Data().Creator().
		WithClassName("Text").
		WithID(id).
		WithProperties(dataSchema).
		WithVector(vector).
		Do(context.Background())

	return err
}

// 根据相似向量搜索 Weaviate 数据库
func searchVectors(vector []float32) (interface{}, error) {
	className := "Text"
	name := graphql.Field{Name: "name"}
	content := graphql.Field{Name: "content"}
	_additional := graphql.Field{
		Name: "_additional",
		Fields: []graphql.Field{
			{Name: "certainty"},
			{Name: "distance"},
		},
	}

	nearVector := weaviateClient.GraphQL().NearVectorArgBuilder().
		WithVector(vector)

	ctx := context.Background()
	result, err := weaviateClient.GraphQL().Get().
		WithClassName(className).
		WithFields(name, content, _additional).
		WithNearVector(nearVector).
		Do(ctx)

	return result, err
}

// 添加文本处理函数
func addTextHandler(c *gin.Context) {
	var input struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	embedding, err := getEmbedding(input.Content, appConfig.OpenAIKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	float32Embedding := make([]float32, len(embedding))
	for i, v := range embedding {
		float32Embedding[i] = float32(v)
	}

	id := uuid.New().String()
	dataSchema := map[string]interface{}{
		"name":    input.Name,
		"content": input.Content,
	}

	err = addVector(id, dataSchema, float32Embedding)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// 搜索相似文本处理函数
func searchSimilarTextsHandler(c *gin.Context) {
	var input struct {
		Content string `json:"content"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	embedding, err := getEmbedding(input.Content, appConfig.OpenAIKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	float32Embedding := make([]float32, len(embedding))
	for i, v := range embedding {
		float32Embedding[i] = float32(v)
	}

	result, err := searchVectors(float32Embedding)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// 启动 API 服务器
func StartServer(config *config.AppConfig) {
	appConfig = config
	initWeaviateClient(config)

	router := gin.Default()

	router.POST("/add-text", addTextHandler)
	router.POST("/search-similar-texts", searchSimilarTextsHandler)

	router.Run(fmt.Sprintf(":%d", config.APIPort))
}
