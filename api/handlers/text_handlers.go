package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/szpnygo/VecTextSearch/config"
	"github.com/szpnygo/VecTextSearch/services"
	"github.com/szpnygo/VecTextSearch/utils"
)

func AddTextHandler(appConfig *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Name    string `json:"name"`
			Content string `json:"content"`
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check for duplicate content if not allowed
		if !appConfig.AllowDuplicateContent {
			response, err := services.FindExactText(appConfig, input.Content)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			if len(response.Data["Get"].(map[string]interface{})["Text"].([]interface{})) > 0 { // 修改这一行
				c.JSON(http.StatusBadRequest, gin.H{"error": "Duplicate content not allowed"})
				return
			}
		}

		id, err := services.AddText(appConfig, input.Name, input.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func SearchSimilarTextsHandler(appConfig *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Content string `json:"content"`
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response, err := services.SearchSimilarTexts(appConfig, input.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		simplifiedResult := utils.SimplifyResult(response.Data)

		c.JSON(http.StatusOK, simplifiedResult)
	}
}
