package router

import (
	"datagram-intelligence/config"
	"datagram-intelligence/internal/handler"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	chatHandler handler.ChatHandler,
	cfg *config.Config,
) *gin.Engine {
	router := gin.Default()

	// Initialize cors config
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowCredentials = true

	router.Use(cors.New(corsConfig))

	apiGroup := router.Group("/api")

	chatGroup := apiGroup.Group("")
	{
		chatGroup.POST("/chat", chatHandler.ChatQuestion)
	}

	apiGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	return router
}
