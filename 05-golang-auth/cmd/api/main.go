package main

import (
	"go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main () {
	// Here we are creating the server and listening to the server
	router := httpserver.NewRouter()

	// Standard go type that runs a http server
	srv := &http.Server {
		Addr: ":5000",
		Handler: router,
		ReadHeaderTimeout: 5*time.Second,
	}

	log.Printf("API running on %s: ", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Printf("Server closed")
			return
		}
		log.Fatalf("Server error: %v", err)
	}
}
