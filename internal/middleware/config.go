package middleware

import (
	"github.com/tabrizgulmammadov/rss-aggregator/config"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/database"
)

type APIConfig struct {
	DB *database.Queries
}

func NewAPIConfig(cfg *config.APIConfig) *APIConfig {
	return &APIConfig{DB: cfg.DB}
}
