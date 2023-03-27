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
