package handlers

import "github.com/gin-gonic/gin"

func GetApi(c *gin.Context){

	c.JSON(200,gin.H{
		"Api": "Pong",
	})

}