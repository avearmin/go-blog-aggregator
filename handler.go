package main

import (
	"net/http"
	"time"

	"github.com/avearmin/go-blog-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
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

func handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, dbUserToJSONUser(user))
}

func (cfg apiConfig) handlePostFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := parameters{}
	if err := readParameters(r, &params); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		Userid:    user.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	payload := [2]any{dbFeedToJSONFeed(feed), dbFeedFollowToJSONFeedFollow(feedFollow)}
	respondWithJSON(w, http.StatusOK, payload)
}

func (cfg apiConfig) handleGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, dbFeedsToJSONFeeds(feeds))
}

func (cfg apiConfig) handleGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetAllFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, dbFeedFollowsToJSONFeedFollows(feedFollows))
}

func (cfg apiConfig) handlePostFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID string `json:"feed_id"`
	}
	params := parameters{}
	if err := readParameters(r, &params); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	feedID, err := uuid.Parse(params.FeedID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    feedID,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, dbFeedFollowToJSONFeedFollow(feedFollow))
}

func (cfg apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID := chi.URLParam(r, "feedFollowID")
	feedFollowUUID, err := uuid.Parse(feedFollowID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = cfg.DB.DeleteFeedFollow(r.Context(), feedFollowUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	type payload struct{}
	respondWithJSON(w, http.StatusOK, payload{})
}

func (cfg apiConfig) handleGetPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		PostsLimit int `json:"limit"`
	}
	params := parameters{}
	if err := readParameters(r, &params); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		Userid: user.ID,
		Limit:  int32(params.PostsLimit),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, dbPostsToJSONPosts(posts))
}
