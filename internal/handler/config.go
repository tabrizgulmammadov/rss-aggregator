// internal/handler/config.go
package handler

import (
	"github.com/tabrizgulmammadov/rss-aggregator/config"
)

type APIConfig struct {
	*config.APIConfig
}

func NewAPIConfig(cfg *config.APIConfig) *APIConfig {
	return &APIConfig{APIConfig: cfg}
}
