package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Response failed with 5XX status with msg: %v", msg)
	}

	type errResponse struct {
		Error string `json:"error"` // JSON reflect tag { error: "Message string" }
	}

	responseWithJson(w, code, errResponse{
		Error: msg,
	})
}
