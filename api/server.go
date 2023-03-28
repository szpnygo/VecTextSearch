package api

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/szpnygo/VecTextSearch/api/handlers"
	"github.com/szpnygo/VecTextSearch/config"
)

func StartServer(appConfig *config.AppConfig) {
	router := gin.Default()

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "Token"}
	router.Use(cors.New(config))

	router.POST("/add-text", handlers.AddTextHandler(appConfig))
	router.POST("/search-similar-texts", handlers.SearchSimilarTextsHandler(appConfig))

	router.Run(fmt.Sprintf(":%d", appConfig.APIPort))
}
