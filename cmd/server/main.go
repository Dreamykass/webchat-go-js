package main

import (
	"log"
	"net/http"
)

const (
	Port = ":8080"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveStatic)
	mux.HandleFunc("/time", serveTime)

	log.Println("Listening...")
	_ = http.ListenAndServe(Port, mux)
}
