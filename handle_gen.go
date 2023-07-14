package main

import (
	"encoding/hex"
	"log"
	"net/http"

	"github.com/somacorp/core/rnd"
)

// Generate a private key and optionally a set of public keys for the curves
// provided.

type GenRequest struct {
	Cur []string `json:"cur"`
}

type GenResponse struct {
	PrivateKey string      `json:"privateKey"`
	PublicKeys []PublicKey `json:"publicKeys"`
}

type PublicKey struct {
	Curve     string `json:"curve"`
	PublicKey string `json:"publicKey"`
}

func (x *GenRequest) Read(r *http.Request) error {
	x.Cur = r.URL.Query()["cur"]
	return nil
}

func (s *server) Gen(w http.ResponseWriter, r *http.Request) {

	var v GenRequest
	if err := v.Read(r); err != nil {
		log.Printf("failed to decode body: %v", err)
		fail(w, http.StatusBadRequest)
		return
	}

	privBytes, err := rnd.RandomBytes(32)
	if err != nil {
		log.Printf("failed to generate random key: %v", err)
		fail(w, http.StatusInternalServerError)
	}
	privHex := hex.EncodeToString(privBytes)

	keys := []PublicKey{}

	for _, curve := range v.Cur {
		pubBytes, err := createPublicKey(curve, privBytes)
		if err != nil {
			continue
		}
		pubHex := hex.EncodeToString(pubBytes)

		keys = append(keys, PublicKey{
			Curve:     curve,
			PublicKey: pubHex,
		})
	}

	res := GenResponse{
		PrivateKey: privHex,
		PublicKeys: keys,
	}

	if err := writeJSON(w, res); err != nil {
		log.Printf("failed to write JSON: %v", err)
		fail(w, http.StatusInternalServerError)
		return
	}
}
