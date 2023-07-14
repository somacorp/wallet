package main

import (
	"log"
	"net/http"
)

// Get all available key ids.

type IdsResponse struct {
	Ids []string `json:"ids"`
}

func (s *server) Ids(w http.ResponseWriter, r *http.Request) {

	ids := []string{}

	for id := range s.keys.privById {
		ids = append(ids, id)
	}

	res := IdsResponse{
		Ids: ids,
	}

	if err := writeJSON(w, res); err != nil {
		log.Printf("failed to write JSON: %v", err)
		fail(w, http.StatusInternalServerError)
		return
	}
}
