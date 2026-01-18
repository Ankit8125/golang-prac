package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// It's going to read those key-value pairs from .env and keep these pairs into a real OS-environment variables. Then we are going to use the OS package.
// godotenv.Load() reads .env and sets them into the process env
// os.getenv -> read those values

type Config struct {
	MongoURI string	
	MongoDBName string
	JWTSecret string
}

func checkEnvErrors (cfg Config) (Config, error) {
	if cfg.MongoURI == "" {
		return Config{}, fmt.Errorf("Missing mongo URI")
	}
	if cfg.MongoDBName == "" {
		return Config{}, fmt.Errorf("Missing mongo DB NAMe")
	}
	if cfg.JWTSecret == "" {
		return Config{}, fmt.Errorf("Missing mongo jwt")
	}
	return cfg, nil
}

func Load () (Config, error) {
	// Step 1 (Read above comments)
	_ = godotenv.Load()

	// Step 2
	cfg := Config {
		MongoURI: strings.TrimSpace(os.Getenv("MONGO_URI")),
		MongoDBName: strings.TrimSpace(os.Getenv("MONGO_DB_NAME")),
		JWTSecret: strings.TrimSpace(os.Getenv("JWT_SECRET")),
	}

	return checkEnvErrors(cfg)
}