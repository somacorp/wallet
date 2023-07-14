package main

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
)

// Create an Elliptic Curve Diffie-Hellman shared secret between the private
// key identified by the key id and a public key.

type ECDHRequest struct {
	Kid string `json:"kid"`
	Pub string `json:"pub"`
	Cur string `json:"cur"`
}

type ECDHResponse struct {
	Sec string `json:"sec"`
	Cur string `json:"cur"`
}

func (s *server) ECDH(w http.ResponseWriter, r *http.Request) {

	var v ECDHRequest
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		log.Printf("failed to decode ecdh body: %v", err)
		fail(w, http.StatusInternalServerError)
		return
	}

	privBytes, err := s.keys.getPrivBytes(v.Kid)
	if err != nil {
		log.Printf("could not get priv bytes: %v", err)
		fail(w, http.StatusBadRequest)
		return
	}

	pubBytes, err := hex.DecodeString(v.Pub)
	if err != nil {
		log.Printf("failed to decode pub: %v", err)
		fail(w, http.StatusInternalServerError)
		return
	}

	secBytes, err := createECDHSharedSecret(v.Cur, privBytes, pubBytes)
	if err != nil {
		log.Println("ecdh failed:", err)
		fail(w, http.StatusInternalServerError)
		return
	}

	secHex := hex.EncodeToString(secBytes)

	res := ECDHResponse{
		Sec: secHex,
		Cur: v.Cur,
	}

	if err := writeJSON(w, res); err != nil {
		log.Printf("failed to write JSON: %v", err)
		fail(w, http.StatusInternalServerError)
		return
	}
}
