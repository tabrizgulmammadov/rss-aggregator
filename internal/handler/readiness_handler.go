package handler

import (
	"github.com/tabrizgulmammadov/rss-aggregator/internal/utils"
	"net/http"
)

func ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, struct{}{})
}
