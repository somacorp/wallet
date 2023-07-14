package main

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
)

// Sign a message using the private key identified by the key id with the
// provided curve and encoding method.

type SignRequest struct {
	Kid string `json:"kid"`
	Cur string `json:"cur"`
	Enc string `json:"enc"`
	Msg string `json:"msg"`
}

type SignResponse struct {
	Kid string `json:"kid"`
	Sig string `json:"sig"`
	Cur string `json:"cur"`
	Enc string `json:"enc"`
}

func (s *server) Sign(w http.ResponseWriter, r *http.Request) {

	var v SignRequest
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		log.Printf("failed to decode sign body: %v", err)
		fail(w, http.StatusInternalServerError)
		return
	}

	msg := []byte(v.Msg)

	if !find(s.conf.Constraints, func(item Constraint) bool {
		kid := contains(item.KeyIds, v.Kid)
		cur := contains(item.Curves, v.Cur)
		enc := contains(item.Encodings, v.Enc)
		val := validate(item.Validation, msg) == nil

		return kid && cur && enc && val
	}) {
		log.Printf("forbidden %s", []string{v.Kid, v.Cur, v.Enc, string(msg)})
		fail(w, http.StatusForbidden)
		return
	}

	encodedData, err := encode(v.Enc, msg)
	if err != nil {
		log.Printf("failed to encode message: %v", err)
		fail(w, http.StatusInternalServerError)
		return
	}

	privBytes, err := s.keys.getPrivBytes(v.Kid)
	if err != nil {
		log.Printf("could not get priv bytes: %v", err)
		fail(w, http.StatusBadRequest)
		return
	}

	sigBytes, err := createSignature(v.Cur, encodedData, privBytes)
	if err != nil {
		log.Printf("failed to create signature: %v", err)
		fail(w, http.StatusInternalServerError)
		return
	}

	sigHex := hex.EncodeToString(sigBytes)

	res := SignResponse{
		Kid: v.Kid,
		Sig: sigHex,
		Cur: v.Cur,
		Enc: v.Enc,
	}

	if err := writeJSON(w, res); err != nil {
		log.Printf("failed to write JSON: %v", err)
		fail(w, http.StatusInternalServerError)
		return
	}
}
