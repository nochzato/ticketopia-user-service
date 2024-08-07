package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/nochzato/ticketopia-user-service/internal/config"
	db "github.com/nochzato/ticketopia-user-service/internal/db/migration"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot read config file: %s", err)
	}

	_, err = pgx.Connect(context.Background(), config.Database.URL)
	if err != nil {
		log.Fatalf("cannot connect to the db: %s", err)
	}

	if err = db.RunMigrations(config); err != nil {
		log.Fatalf("failed to run migrations: %s", err)
	}
}
