package models

import (
	"context"
	"os"
	"log"

	"watch-party-api/config"

	"github.com/jackc/pgx/v5/pgxpool"

)

var DB *pgxpool.Pool

//  ConnectDB initializes the PostgreSQL connection using pgx (psql driver for golang)
func ConnectDB() {
	cfg := config.LoadConfig()

	dbURL := cfg.DatabaseURL
	
	// parse config makes a config out of a db connection string
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %s", err)
	}

	// create a new connection pool (method used to keep db connections open so they can be reused by others)
	DB, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s", err)
	}

	log.Println("Connected to PostgreSQL successfully")

}


