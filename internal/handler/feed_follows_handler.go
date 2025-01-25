package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	_ "github.com/tabrizgulmammadov/rss-aggregator/api"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/database"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/models"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/utils"
	"log"
	"net/http"
	"time"
)

// CreateFeedFollowRequest represents the request body for creating a feed follow
type CreateFeedFollowRequest struct {
	FeedID uuid.UUID `json:"feed_id" example:"123e4567-e89b-12d3-a456-426614174000" format:"uuid"`
}

// HandlerCreateFeedFollow creates a new feed follow.
//
//	@Summary		Create a new feed follow
//	@Description	Allows a user to follow a specific feed
//	@Tags			Feed Follow
//	@Accept			json
//	@Produce		json
//	@Param			request	body	CreateFeedFollowRequest	true	"Feed Follow Request"
//	@Security		ApiKeyAuth
//	@Success		201	{object}	models.FeedFollow		"Successfully created feed follow"
//	@Failure		400	{object}	utils.JSONErrorResponse	"Error parsing request body or creating feed follow"
//	@Router			/feed-follows [post]
func (cfg *APIConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}
	log.Println(params)

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

// HandlerGetFeedFollows retrieves feed follows by user's ID.
//
//	@Summary		Get all feed follows for a user
//	@Description	Fetches a list of all feeds the user is following
//	@Tags			Feed Follow
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{array}		models.FeedFollow		"List of feed follows"
//	@Failure		400	{object}	utils.JSONErrorResponse	"Error fetching feed follows"
//	@Router			/feed-follows [get]
func (cfg *APIConfig) HandlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

// HandlerDeleteFeedFollow deletes feed follow by user's and feed follow's IDs.
//
//	@Summary		Delete a feed follow
//	@Description	Allows a user to unfollow a specific feed
//	@Tags			Feed Follow
//	@Param			feedFollowID	path	string	true	"ID of the feed follow to delete"	format(uuid)
//	@Security		ApiKeyAuth
//	@Success		200	{object}	utils.EmptyResponse		"Successfully deleted feed follow"
//	@Failure		400	{object}	utils.JSONErrorResponse	"Error parsing feed follow ID or deleting feed follow"
//	@Router			/feed-follows/{feedFollowID} [delete]
func (cfg *APIConfig) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't parse feed follow ID: %v", err))
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowID,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.EmptyResponse{})
}
