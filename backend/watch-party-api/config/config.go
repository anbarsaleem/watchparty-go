package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config structure to hold env variables
type Config struct {
	DatabaseURL string
	ServerPort string
	JWTSecret string
}

// reads from env file and system environment and returns a config
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env variables instead")
	}

	// Read env variables into config
	config := &Config {
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/watch_party_db?sslmode=disable"),
		ServerPort: getEnv("PORT", "8080"),
		JWTSecret: getEnv("JWT_SECRET", "defaultSecret"),
	}

	return config
}

// helper func to get env with fall back value
func getEnv(key, fallback string) string {

	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback

}

