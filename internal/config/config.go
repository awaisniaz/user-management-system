package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort   string
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    string
	RedisAddr string
	JWTSecret string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil{
		log.Fatalf("Failed to load config: %v", err)
	}

	AppConfig = Config{
		AppPort:   viper.GetString("APP_PORT"),
		DBHost:    viper.GetString("DB_HOST"),
		DBUser:    viper.GetString("DB_USER"),
		DBPass:    viper.GetString("DB_PASSWORD"),
		DBName:    viper.GetString("DB_NAME"),
		DBPort:    viper.GetString("DB_PORT"),
		RedisAddr: viper.GetString("REDIS_ADDR"),
		JWTSecret: viper.GetString("JWT_SECRET"),
	}
}