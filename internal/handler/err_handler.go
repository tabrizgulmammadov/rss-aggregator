package handler

import (
	"github.com/tabrizgulmammadov/rss-aggregator/internal/utils"
	"net/http"
)

func ErrHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusBadRequest, "Something went wrong")
}
