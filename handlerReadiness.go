package main

import (
	"net/http"
)

func handleReady(w http.ResponseWriter, r *http.Request) {
	responWithJSON(w, 200, struct{}{})
}
