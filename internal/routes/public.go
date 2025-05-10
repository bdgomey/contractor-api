package routes

import (
	"github.com/bdgomey/construction-site/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(router *gin.Engine) {
	// Get methods
	router.GET("/", handlers.GetHome)
	router.GET("/home", handlers.GetHome)
	router.GET("/about", handlers.GetAbout)

}
