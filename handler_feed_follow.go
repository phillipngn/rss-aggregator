package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/phillipngn/rss-aggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Could not parse feed id")
		return
	}

	feedFollows, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Could not create feed follow")
		return
	}

	responseWithJson(w, http.StatusOK, parseFeedFollow(feedFollows))
}

func (cfg *apiConfig) handlerFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdString := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdString)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Could not parse feed follow id")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Could not delete feed follow")
		return
	}

	responseWithJson(w, http.StatusOK, struct{}{})
}

func (cfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.ListFeedFollowsByUser(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Could not get feed follows")
		return
	}
	responseWithJson(w, http.StatusOK, parseFeedFollows(feedFollows))
}
