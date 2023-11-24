package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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
	subRouter.Get("/readiness", handleEndpointReadiness)
	subRouter.Get("/err", handleEndpointErr)
	mainRouter.Mount("/v1", subRouter)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}
	server.ListenAndServe()
}
