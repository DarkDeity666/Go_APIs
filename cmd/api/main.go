package main

import(
	"fmt"
	"github.com/DarkDeity666/Go_APIs/internal/config"
	"github.com/gin-gonic/gin"
) 

func main(){
	cfg := config.MustLoad()
	router := gin.Default()
	fmt.Println(cfg)

	
}