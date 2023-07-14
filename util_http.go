package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJSON(w http.ResponseWriter, v any) error {
	jsonBytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(jsonBytes); err != nil {
		return fmt.Errorf("failed to write result: %v", err)
	}

	return nil
}

func fail(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}
