package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/DarkDeity666/Go_APIs/internal/models"
)

func Register(c *gin.Context){
	var req models.RegisterRequest

	if err:= c.ShouldBindJSON(&req)
	err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"message":"Invlaid request body",
		})
		return 
	}
	c.JSON(http.StatusOK, gin.H{
		"success":true,
		"data": req,
	})
}
