package routes

import (
	"github.com/DarkDeity666/Go_APIs/internal/handlers"
	"github.com/gin-gonic/gin"
)

func TestRoutes(router *gin.RouterGroup){
	auth:= router.Group("test")
	auth.GET("/health", handlers.Health)
}
func RegisterNewUser(router *gin.RouterGroup){
	auth:= router.Group("auth")
	auth.GET("/register", handlers.Register)
}