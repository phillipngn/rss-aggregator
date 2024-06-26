package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/phillipngn/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	decodeErr := decoder.Decode(&params)
	if decodeErr != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON:%s", decodeErr))
		return
	}

	newUser, createUserErr := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if createUserErr != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Could not create user:%s", createUserErr))
	}

	responseWithJson(w, http.StatusCreated, parseUser(newUser))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responseWithJson(w, http.StatusOK, parseUser(user))
}
