package main

import (
	"log"
	"net/http"
	"ws/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel Listener")
	go handlers.ListenTothisWsChannel()

	log.Println("Starting web server at port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
