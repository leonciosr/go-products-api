package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Error().Err(err).Msg("unable to load .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Error().Err(err)
	}

	err = db.Ping()
	if err != nil {
		log.Error().Err(err)
	}

	log.Info().Msg("connected to " + os.Getenv("POSTGRES_DB"))
	return db, nil
}
