package main

import (
	"fmt"
	"log"
	"notes-api/internal/config"
	"notes-api/internal/db"
	"notes-api/internal/server"
)

// This is out entry point
func main () {

	// Firstly, we need to get our env variables. 
	// Step 1: Load our config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error")
	}

	// Step 2: Connect to DB
	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("DB error: %v", err)
	}

	// Step 3: Disconnect if you find something is wrong
	defer func ()  {
		if err := db.Disconnect(client); err != nil {
			log.Printf("Mongo Disconnected %v", err)
		}
	} ()

	router := server.NewRouter(database)
	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Server failed")
	}

}