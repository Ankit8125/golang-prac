package app

import (
	"context"
	"fmt"
	"go-auth/internal/config"
	"go-auth/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Config config.Config
	MongoClient *mongo.Client
	DB *mongo.Database
}

// Joining all the existing pieces together (mongo.go, config.go, main.go, etc)
func New (ctx context.Context) (*App, error) {
	// Step 1: Loading the .env variables
	cfg, err := config.Load() 
	if err != nil {
		return nil, err
	}

	// Step 2: Make the DB connection
	mongoClnt, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Config: cfg,
		MongoClient: mongoClnt.Client,
		DB: mongoClnt.DB,
	}, nil
}

func (a *App) Close (ctx context.Context) error {
	if a.MongoClient != nil {
		return nil
	}

	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("Mongo Disconnected Failed: %w", err)
	}

	return nil
}