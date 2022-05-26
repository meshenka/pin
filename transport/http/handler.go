package http

import (
	"encoding/json"
	"net/http"
)

// Generator produce a valid pin.
type Generator interface {
	Generate() string
}

func get(g Generator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"pin": g.Generate()})
	}
}

func heartbeat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")

		w.WriteHeader(http.StatusOK)
	}
}
