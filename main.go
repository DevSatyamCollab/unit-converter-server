package main

import (
	"fmt"
	"log"
	"net/http"
	"unit-converter/internal/api"
	"unit-converter/internal/handler"
)

func main() {
	port := ":8080"

	mux := http.NewServeMux()

	// web  handler
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/length", handler.LengthHandler)
	mux.HandleFunc("/weight", handler.WeightHandler)
	mux.HandleFunc("/temperature", handler.TemperatureHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/static"))))
	// api (third party)
	mux.HandleFunc("POST /lengths", api.LengthHandler)
	mux.HandleFunc("POST /weights", api.WeightHandler)
	mux.HandleFunc("POST /temperatures", api.TemperatureHanlder)

	svr := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	fmt.Printf("Starting a server at port %s\n", port)
	log.Fatal(svr.ListenAndServe())
}
