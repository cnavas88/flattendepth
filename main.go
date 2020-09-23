package main

import (
	"log"
	"net/http"
	"time"

	"./server"
)

func main() {
	router := server.NewRouter()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening...")
	server.ListenAndServe()
}
