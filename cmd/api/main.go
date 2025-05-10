package main

import (
	"github.com/bdgomey/construction-site/internal/routes"
	"github.com/gin-gonic/gin"
)

func main(){

	r := gin.Default()
	
	routes.RegisterPublicRoutes(r)
	routes.RegisterApiRoutes(r)
	
	r.Run()

}