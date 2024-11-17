package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Todo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	todos := []Todo{}

	r := chi.NewRouter()

	r.Use((middleware.Logger))

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	r.Post("/todos", func(w http.ResponseWriter, r *http.Request) {
		todo := Todo{}

		err := json.NewDecoder(r.Body).Decode(&todo)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request body"))
			return
		}

		todos = append(todos, todo)
		w.WriteHeader(http.StatusCreated)
	})
}
