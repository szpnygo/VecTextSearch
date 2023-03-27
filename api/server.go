package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/szpnygo/VecTextSearch/api/handlers"
	"github.com/szpnygo/VecTextSearch/config"
)

func StartServer(appConfig *config.AppConfig) {
	router := gin.Default()

	router.POST("/add-text", handlers.AddTextHandler(appConfig))
	router.POST("/search-similar-texts", handlers.SearchSimilarTextsHandler(appConfig))

	router.Run(fmt.Sprintf(":%d", appConfig.APIPort))
}
