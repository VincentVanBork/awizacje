package models

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"os"
)

func Migrate(db *sql.DB) error {
	driver, startupErr := postgres.WithInstance(db, &postgres.Config{})
	if startupErr != nil {
		return startupErr
	}
	m, setupErr := migrate.NewWithDatabaseInstance(
		os.Getenv("MIGRATIONS_PATH"),
		os.Getenv("POSTGRES_DB"), driver)
	if setupErr != nil {
		return setupErr
	}
	migrateError := m.Up()
	if migrateError != nil {
		return migrateError
	}
	return nil
}
