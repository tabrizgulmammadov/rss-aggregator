package main

import (
	_ "github.com/tabrizgulmammadov/rss-aggregator/api"
	"github.com/tabrizgulmammadov/rss-aggregator/config"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/database"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/handler"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/middleware"
	platform "github.com/tabrizgulmammadov/rss-aggregator/internal/platform/database"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/routes"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/service"
	"log"
	"net/http"
	"time"
)

// @title						RSS Aggregator Api
// @version					    1.0
// @description				    This is an RSS Aggregator Api
// @termsOfService				http://swagger.io/terms/
// @BasePath					/v1
// @securityDefinitions.apiKey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	serverCfg := config.Get()

	// Initialize database with the new package
	dbConfig := platform.NewConfig(serverCfg.DBUrl)
	conn, err := platform.Initialize(dbConfig)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	dbQueries := database.New(conn)
	apiCfg := config.APIConfig{
		DB: dbQueries,
	}
	handlerAPIConfig := handler.NewAPIConfig(&apiCfg)
	middlewareAPIConfig := middleware.NewAPIConfig(&apiCfg)

	// Start RSS scraping in a goroutine
	go service.StartScraping(dbQueries, 10, time.Minute)

	// Initialize routes
	router := routes.SetupRoutes(handlerAPIConfig, middlewareAPIConfig)

	// Configure and start server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + serverCfg.Port,
	}

	log.Printf("Server starting on port %s\n", serverCfg.Port)
	log.Fatal(server.ListenAndServe())
}
