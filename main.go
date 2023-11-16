package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")

	mainRouter := chi.NewRouter()

	corsOptions := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
	}
	mainRouter.Use(cors.Handler(corsOptions))

	subRouter := chi.NewRouter()
	mainRouter.Mount("/v1", subRouter)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}
	server.ListenAndServe()
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(status)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code/100 == 5 {
		log.Printf("Responding with status code %d: %s", code, msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}
