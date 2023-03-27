package utils

import (
	"github.com/weaviate/weaviate/entities/models"
)

func SimplifyResult(data map[string]models.JSONObject) []map[string]interface{} {
	get := data["Get"].(map[string]interface{})
	texts := get["Text"].([]interface{})

	simplifiedTexts := make([]map[string]interface{}, len(texts))

	for i, text := range texts {
		textMap := text.(map[string]interface{})
		additional := textMap["_additional"].(map[string]interface{})
		certainty := additional["certainty"].(float64)
		distance := additional["distance"].(float64)
		content := textMap["content"].(string)
		name := textMap["name"].(string)

		simplifiedText := map[string]interface{}{
			"certainty": certainty,
			"distance":  distance,
			"content":   content,
			"name":      name,
		}
		simplifiedTexts[i] = simplifiedText
	}

	return simplifiedTexts
}
