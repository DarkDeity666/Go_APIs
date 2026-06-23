package main

import(
	"github.com/DarkDeity666/Go_APIs/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/DarkDeity666/Go_APIs/internal/routes"
) 

func main(){
	//config load
	cfg := config.MustLoad()

	//gin engine creation
	router := gin.Default()

	//Api grouping
	api:= router.Group("/api/v1")

	//register router
	routes.RegisterAuthRouter(api)

	//server start
	router.Run(":" + cfg.Port)
}