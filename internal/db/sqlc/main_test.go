package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/nochzato/ticketopia-user-service/internal/config"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../../..")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	conn, err := pgx.Connect(context.Background(), config.Database.URL)
	if err != nil {
		log.Fatalf("cannot connect to the db: %s", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
