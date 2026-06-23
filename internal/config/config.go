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

func Load() (*Config, error) {
	err:= godotenv.Load()
	if err != nil{
		return nil, err 
	}

	ctg := &Config{
		Port: os.Getenv("PORT"),
		MongoURI: os.Getenv("MONGO_URI"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		JWTRefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
		CloudinaryCloudName: os.Getenv("CLOUDINARY_CLOUD_NAME"),
		CloudinaryAPIKey: os.Getenv("CLOUDINARY_API_KEY"),
		CloudinaryAPISecret: os.Getenv("CLOUDINARY_API_SECRET"),
	}
	return ctg, nil
}

func MustLoad() *Config {
	cfg,err := Load()
	if err != nil || cfg == nil{
		panic(err)
	}
	return cfg 
}