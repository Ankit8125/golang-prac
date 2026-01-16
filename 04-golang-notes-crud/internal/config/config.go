package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MONGOURI string
	MONGODB string
	ServerPort string
}

func Load() (Config, error) {
	// godotenv.Load() reads .env and sets them into the process env
	// os.getenv -> read those values

	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("Failed to load .env")
	}

	mongoURI, err := extractEnv("MONGO_URI")
	if err != nil {
		return Config{}, err
	}

	mongoDbName, err := extractEnv("MONGO_DB_NAME")
	if err != nil {
		return Config{}, err
	}

	port, err := extractEnv("PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
		MONGOURI: mongoURI,
		MONGODB: mongoDbName,
		ServerPort: port,
	}, nil

}

func extractEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("Missing req. env")
	}

	return val, nil
}