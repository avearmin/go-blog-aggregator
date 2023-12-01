package main

import (
	"net/http"
	"time"

	"github.com/avearmin/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func handleEndpointReadiness(w http.ResponseWriter, r *http.Request) {
	type payload struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, http.StatusOK, payload{
		Status: "ok",
	})
}

func handleEndpointErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

func (cfg apiConfig) handlePostUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}
	if err := readParameters(r, &params); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, dbUserToJSONUser(user))
}

func (cfg apiConfig) handleGetCurrentUser(w http.ResponseWriter, r *http.Request) {
	apikey, err := readApikey(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	user, err := cfg.DB.GetUserByApikey(r.Context(), apikey)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, dbUserToJSONUser(user))
}
