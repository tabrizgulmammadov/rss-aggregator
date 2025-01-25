package handler

import (
	_ "github.com/tabrizgulmammadov/rss-aggregator/api"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/utils"
	"net/http"
)

// HandlerErr handles unexpected errors.
//
//	@Summary		Handle errors
//	@Description	Return a generic error message for unexpected errors
//	@Tags			Error
//	@Accept			json
//	@Produce		json
//	@Success		400	{object}	utils.JSONErrorResponse
//	@Router			/error [get]
func HandlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusBadRequest, "Something went wrong")
}
