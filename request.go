package main

import (
	"encoding/json"
	"net/http"
)

func readRequest(r *http.Request, parameters interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(parameters); err != nil {
		return err
	}
	return nil
}
