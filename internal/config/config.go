package config 

import (
	"os"
	"github.com/joho/godotenv"
)
type Config struct{
	Port string
	MongoURI string
	RedisAddr string
	JWTSecret string 
	JWTRefreshSecret string

	CloudinaryCloudName string
	CloudinaryAPIKey string
	CloudinaryAPISecret string
}

func Load()(*config,error){

}