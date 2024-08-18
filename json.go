package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responwithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx level error: ", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	responWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func responWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed To Marshal JSON Response %v", err)
		w.WriteHeader(500)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	w.Write(data)

}
