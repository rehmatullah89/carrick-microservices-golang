package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type _DBConfigStruct struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type _RedisConfigStruct struct {
	Host     string
	Port     string
	Password string
	DB       int
	Prefix   string
}

type _RabbitMQConfigStruct struct {
	Host     string
	Port     string
	User     string
	Password string
}

type AppConfigStruct struct {
	DB       _DBConfigStruct
	Redis    _RedisConfigStruct
	RabbitMQ _RabbitMQConfigStruct

	CacheTTL int
}

var AppConfig AppConfigStruct

func init() {
	err := godotenv.Load()
	if err != nil && os.Getenv("DB_HOST") == "" {
		log.Fatal("Error loading .env file")
	}

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		redisDB = 0
	}

	redisPrefix := os.Getenv("REDIS_PREFIX")
	if redisPrefix == "" {
		redisPrefix = "api_"
	}

	CacheTTL, err := strconv.Atoi(os.Getenv("CACHE_TTL"))
	if err != nil {
		CacheTTL = 0
	}

	AppConfig = AppConfigStruct{
		DB: _DBConfigStruct{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Redis: _RedisConfigStruct{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       redisDB,
			Prefix:   redisPrefix,
		},
		RabbitMQ: _RabbitMQConfigStruct{
			Host:     os.Getenv("RABBITMQ_HOST"),
			Port:     os.Getenv("RABBITMQ_PORT"),
			User:     os.Getenv("RABBITMQ_USER"),
			Password: os.Getenv("RABBITMQ_PASSWORD"),
		},
		CacheTTL: CacheTTL,
	}
}
