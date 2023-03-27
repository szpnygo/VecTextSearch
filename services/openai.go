package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

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
