package routes

import (
	"github.com/DarkDeity666/Go_APIs/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRouter(router *gin.RouterGroup){
	auth:= router.Group("auth")
	auth.GET("/health", handlers.Health)
}