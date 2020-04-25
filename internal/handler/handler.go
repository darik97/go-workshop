package handler

import (
	"fmt"
	"net/http"
	"workshop/internal/api"
)

// Handler struct
type Handler struct {
	jokeClient api.Client
}

// NewHandler constructor for Handler
func NewHandler(jokeClient api.Client) *Handler {
	return &Handler{
		jokeClient: jokeClient,
	}
}

// Hello handler for getting joke
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	joke, err := h.jokeClient.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprint(w, joke.Joke)
}
