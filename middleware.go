package main

import (
	"net/http"

	"github.com/avearmin/go-blog-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		handler(w, r, user)
	})
}
