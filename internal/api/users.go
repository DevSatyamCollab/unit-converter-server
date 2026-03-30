package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Converter struct {
	Value    float64 `json:"value"`
	Ans      float64 `json:"ans"`
	FromUnit string  `json:"fromUnit"`
	ToUnit   string  `json:"toUnit"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func errorResponse(w http.ResponseWriter, status int, msg string) {
	response(w, status, ErrorResponse{
		Error: msg,
	})
}

func response(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling json: %s\n", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
