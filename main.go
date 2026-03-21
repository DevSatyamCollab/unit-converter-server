package main

import (
	"fmt"
	"log"
	"net/http"
	"unit-converter/handler"
)

func main() {
	port := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("POST /lengths", handler.LengthHandler)
	mux.HandleFunc("POST /weights", handler.WeightHandler)
	mux.HandleFunc("POST /temperatures", handler.TemperatureHanlder)

	svr := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	fmt.Printf("Starting a server at port %s\n", port)
	log.Fatal(svr.ListenAndServe())
}
