package main

import (
	"net/http"
)

func handleError(w http.ResponseWriter, r *http.Request) {
	responwithError(w, 400, "Something went wrong")
}
