package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/avearmin/go-blog-aggregator/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	dbURL := os.Getenv("CONN_STRING")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	dbQueries := database.New(db)

	apiConfig := apiConfig{DB: dbQueries}

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
	subRouter.Post("/users", apiConfig.handlePostUser)
	subRouter.Get("/users", apiConfig.middlewareAuth(handleGetUser))
	subRouter.Get("/feeds", apiConfig.handleGetAllFeeds)
	subRouter.Post("/feeds", apiConfig.middlewareAuth(apiConfig.handlePostFeed))
	mainRouter.Mount("/v1", subRouter)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}
	server.ListenAndServe()
}
