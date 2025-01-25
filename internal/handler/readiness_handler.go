package handler

import (
	_ "github.com/tabrizgulmammadov/rss-aggregator/api"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/utils"
	"net/http"
)

// HandlerReadiness indicates if the service is ready to handle requests.
//
//	@Summary		Readiness check
//	@Description	Endpoint to check if the service is ready to handle requests
//	@Tags			Health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.EmptyResponse	"Service is ready"
//	@Router			/healthz [get]
func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, utils.EmptyResponse{})
}
