package main

import "net/http"

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
