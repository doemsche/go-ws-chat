package main

import (
	"go-ws-chat/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("staring channel listener")
	go handlers.ListenToWsChannel()

	log.Println("starting webserver on port 8080")
	_ = http.ListenAndServe(":8080", mux)
}
