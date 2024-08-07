package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nochzato/ticketopia-user-service/internal/config"
)

func RunMigrations(config *config.Config) error {
	m, err := migrate.New(config.Database.MigrationPath, config.Database.URL)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
