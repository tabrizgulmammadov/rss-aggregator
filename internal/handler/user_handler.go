package handler

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/database"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/models"
	"github.com/tabrizgulmammadov/rss-aggregator/internal/utils"
	"net/http"
	"time"
)

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

func (apiCfg *APIConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, http.StatusOK, models.DatabaseUserToUser(user))
}

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
