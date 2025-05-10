package routes

import (
	"github.com/bdgomey/construction-site/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(router *gin.Engine) {
	// Get methods
	router.GET("/api", handlers.GetApi)


}

