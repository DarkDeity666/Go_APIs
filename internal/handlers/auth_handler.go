package handlers

import "github.com/gin-gonic/gin"

func Health(c *gin.Context){
	c.JSON(200,gin.H{
		"message": "Auth service is running ",
		"success": true,
	})
}