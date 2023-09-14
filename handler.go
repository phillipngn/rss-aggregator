package main

import "net/http"

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	responseWithJson(w, 200, struct{}{})
}

func handleError(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 500, "Server error")
}
