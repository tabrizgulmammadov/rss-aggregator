package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/database"
)

type APIConfig struct {
	DB *database.Queries
}

type ServerConfig struct {
	Port  string
	DBUrl string
}

var (
	serverConfig *ServerConfig
	once         sync.Once
)

func Get() *ServerConfig {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalln("No .env file found")
		}

		serverConfig = &ServerConfig{
			Port:  mustGetenv("PORT"),
			DBUrl: mustGetenv("DB_URL"),
		}
	})
	return serverConfig
}

func mustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s environment variable not set", key)
	}
	return value
}
