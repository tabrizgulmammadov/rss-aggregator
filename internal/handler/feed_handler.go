package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/tabrizgulmammadov/rss-aggregator/api"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/database"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/models"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/utils"
	"net/http"
	"time"
)

// HandlerCreateFeed creates a new feed.
//
//	@Summary		Create a new feed
//	@Description	Allows a user to create a new feed with a name and URL
//	@Tags			Feed
//	@Accept			json
//	@Produce		json
//	@Param			feed	body	models.FeedRequest	true	"Feed information"
//	@Security		ApiKeyAuth
//	@Success		200	{object}	models.Feed				"Successfully created feed"
//	@Failure		400	{object}	utils.JSONErrorResponse	"Error parsing request body or creating feed"
//	@Router			/feeds [post]
func (apiCfg *APIConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)

	params := models.FeedRequest{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusCreated, fmt.Sprintf("Could not create feed: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.DatabaseFeedToFeed(feed))
}

// HandlerGetFeeds retrieves feeds.
//
//	@Summary		Get all feeds
//	@Description	Retrieves a list of all available feeds
//	@Tags			Feed
//	@Produce		json
//	@Success		200	{array}		models.Feed				"List of feeds"
//	@Failure		400	{object}	utils.JSONErrorResponse	"Error fetching feeds"
//	@Router			/feeds [get]
func (apiCfg *APIConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		utils.RespondWithError(w, http.StatusCreated, fmt.Sprintf("Could not get feeds: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.DatabaseFeedsToFeeds(feeds))
}
