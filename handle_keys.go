package main

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
)

// Get the public key for each (key id, curve) pair.

type KeysRequest struct {
	Ids []struct {
		Kid string `json:"kid"`
		Cur string `json:"cur"`
	} `json:"ids"`
}

type KeysResponse struct {
	Keys []Key `json:"keys"`
}

type Key struct {
	Kid string `json:"kid"`
	Pub string `json:"pub"`
	Cur string `json:"cur"`
}

func (s *server) Keys(w http.ResponseWriter, r *http.Request) {

	var v KeysRequest
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		log.Printf("failed to decode keys body: %v", err)
		fail(w, http.StatusInternalServerError)
		return
	}

	keys := []Key{}

	for _, data := range v.Ids {
		privHex, ok := s.keys.privById[data.Kid]
		if !ok {
			continue
		}

		privBytes, err := hex.DecodeString(privHex)
		if err != nil {
			log.Printf("failed to decode privHex: %v", err)
			continue
		}

		pubBytes, err := createPublicKey(data.Cur, privBytes)
		if err != nil {
			log.Printf("failed to get pub key from curve '%s': %v", data.Cur, err)
			continue
		}

		pubHex := hex.EncodeToString(pubBytes)

		keys = append(keys, Key{
			Kid: data.Kid,
			Pub: pubHex,
			Cur: data.Cur,
		})
	}

	res := KeysResponse{
		Keys: keys,
	}

	if err := writeJSON(w, res); err != nil {
		log.Printf("failed to write JSON: %v", err)
		fail(w, http.StatusInternalServerError)
		return
	}
}
