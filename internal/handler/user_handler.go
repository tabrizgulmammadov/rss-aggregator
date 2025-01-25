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

// HandlerCreateUser creates a new user.
//
//	@Summary		Create a user
//	@Description	Create a new user by providing their name
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			name	body		string	true	"User Name"
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	utils.JSONErrorResponse
//	@Router			/users [post]
func (apiCfg *APIConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not create user: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.DatabaseUserToUser(user))
}

// HandlerGetUser retrieves a user by ID.
//
//	@Summary		Get a user
//	@Description	Retrieve details of a user by their ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	models.User
//	@Failure		400	{object}	utils.JSONErrorResponse
//	@Router			/users [get]
func (apiCfg *APIConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, http.StatusOK, models.DatabaseUserToUser(user))
}

// HandlerGetPostsForUser retrieves posts for a user.
//
//	@Summary		Get posts for a user
//	@Description	Retrieve all posts for a specific user
//	@Tags			Posts
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{array}		models.Post
//	@Failure		400	{object}	utils.JSONErrorResponse
//	@Router			/posts [get]
func (apiCfg *APIConfig) HandlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not get posts: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.DatabasePostsToPosts(posts))
}
