package main

import (
	"fmt"
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
	fmt.Print("Start Success!")
	server.ListenAndServe()
}
