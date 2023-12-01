package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

var (
	malformedAuthHeaderError = errors.New("Malformed Authorization header.")
	notApiKeyError           = errors.New("Authorization header was not an API key")
	apiKeyNotCorrectLenError = errors.New("API key was not the corrrect length.")
)

func readParameters(r *http.Request, parameters interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(parameters); err != nil {
		return err
	}
	return nil
}

func readApikey(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return "", malformedAuthHeaderError
	}
	if fields[0] != "ApiKey" {
		return "", notApiKeyError
	}
	if len(fields[1]) != 64 {
		return "", apiKeyNotCorrectLenError
	}
	return fields[1], nil
}
