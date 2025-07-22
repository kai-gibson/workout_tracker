package db

import (
	"fmt"
  "database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDB(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Couldn't get db instance: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)

	if err != nil {
		return fmt.Errorf("error getting migration instance: %w", err)
	}

  err = m.Up()
  if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration error: %w", err)
  }

	return nil
}
