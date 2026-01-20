package main

import (
	"context"
	"go-auth/internal/app"
	"go-auth/internal/httpserver"
	"log"
	"net/http"
	"time"
)

func main () {

	// Step 1: Creating our root context
	// Root context: We use context.Background() whenever we are doing a startup or we need a the starting parent context.
	ctx := context.Background()
	
	// Step 2: Passing the root context in app.new; App.new()
	// 		  - Loads the env
	// 		  - Makes the DB connection
	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("Startup failed: %v", err)
	}

	defer func () {
		if err := a.Close(ctx); err != nil {
			log.Printf("Shutdown warning: %v", err)
		}
	} ()

	// Here we are creating the server and listening to the server
	router := httpserver.NewRouter(a) // We are passing this "a", because we are getting the config and DB

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
