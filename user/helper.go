package main

import (
	"encoding/json"
	"net/http"
)

// respondError return error message
func respondError(w http.ResponseWriter, code int, msg string) {
	respondJSON(w, code, map[string]string{"message": msg})
}

// respondJSON write json response format
func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
