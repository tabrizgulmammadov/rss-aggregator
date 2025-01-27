package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/handler"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/middleware"
	"net/http"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(handlerAPIConfig *handler.APIConfig, middlewareAPIConfig *middleware.APIConfig) http.Handler {
	router := chi.NewRouter()

	// Set up CORS
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Health and error routes
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handler.HandlerReadiness)
	v1Router.Get("/error", handler.HandlerErr)

	// User routes
	v1Router.Get("/users", middlewareAPIConfig.AuthMiddleware(handlerAPIConfig.HandlerGetUser))
	v1Router.Post("/users", handlerAPIConfig.HandlerCreateUser)

	// Feed routes
	v1Router.Get("/feeds", handlerAPIConfig.HandlerGetFeeds)
	v1Router.Post("/feeds", middlewareAPIConfig.AuthMiddleware(handlerAPIConfig.HandlerCreateFeed))

	// Post routes
	v1Router.Get("/posts", middlewareAPIConfig.AuthMiddleware(handlerAPIConfig.HandlerGetPostsForUser))

	// Feed follows routes
	v1Router.Get("/feed-follows", middlewareAPIConfig.AuthMiddleware(handlerAPIConfig.HandlerGetFeedFollows))
	v1Router.Post("/feed-follows", middlewareAPIConfig.AuthMiddleware(handlerAPIConfig.HandlerCreateFeedFollow))
	v1Router.Delete("/feed-follows/{feedFollowID}", middlewareAPIConfig.AuthMiddleware(handlerAPIConfig.HandlerDeleteFeedFollow))

	// Mount Swagger and API version
	router.Mount("/swagger", httpSwagger.WrapHandler)
	router.Mount("/v1", v1Router)

	return router
}
