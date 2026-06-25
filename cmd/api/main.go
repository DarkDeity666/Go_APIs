package main

import (
	"log"
	"github.com/DarkDeity666/Go_APIs/internal/config"
	"github.com/DarkDeity666/Go_APIs/internal/database"
	"github.com/DarkDeity666/Go_APIs/internal/routes"
	"github.com/gin-gonic/gin"
) 

func main(){
	//config load
	cfg := config.MustLoad()

	//mongodb connection function call for connection
	mongoClient , err := database.ConnectionMongo(cfg.MongoURI)
	if err != nil{
		log.Fatal(err)
	}
	_ = mongoClient
	//redis clinet connection function call
	// redisClient , err := database.ConnectionRedis(cfg.RedisAddr)

	//gin engine creation
	router := gin.Default()

	//Api grouping
	api:= router.Group("/api/v1")

	//register router
	routes.RegisterAuthRouter(api)

	//server start
	router.Run(":" + cfg.Port)
}