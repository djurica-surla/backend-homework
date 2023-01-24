package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "modernc.org/sqlite"
)

const (
	sqliteDriver = "sqlite"
	// migrations table name (homework_schema_migrations).
	sqliteMigrationsTable = "homework"
)

var (
	ErrFailedConnection = errors.New("database connection failed")
	ErrDriver           = errors.New("database migration driver creation failed")
	ErrReadMigration    = errors.New("database migration reading files failed")
	ErrMigration        = errors.New("database migration failed")
)

// Repository is the concrete implementations that allows to use sqlite as store.
type Repository struct {
	Instance *sql.DB
}

// Configuration for creating a new sqlite instance.
type Config struct {
	DSN string
}

// New creates a database client wrapper.
func NewRepository(
	ctx context.Context,
	cfg Config,
) (*Repository, error) {
	instance, err := sql.Open(sqliteDriver, cfg.DSN)
	if err != nil {
		return nil, ErrFailedConnection
	}

	err = instance.Ping()
	if err != nil {
		return nil, ErrFailedConnection
	}

	log.Println("Connection successful!")
	return &Repository{Instance: instance}, nil
}

// Migrate makes sure database migrations are up to date.
func (db Repository) Migrate(path string) error {
	driver, err := sqlite.WithInstance(db.Instance, &sqlite.Config{
		MigrationsTable: fmt.Sprintf("%s_%s", sqliteMigrationsTable, sqlite.DefaultMigrationsTable),
	})
	if err != nil {
		return fmt.Errorf("%s: %w", ErrDriver, err)
	}

	// Read migration files.
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", path), sqliteDriver, driver)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrReadMigration, err)
	}

	// Perform database migration.
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("%s: %w", ErrMigration, err)
	} else if err == migrate.ErrNoChange {
		v, _, _ := m.Version()
		log.Printf("postgres migrations up to date, version: %d", v)
	} else if err == nil {
		v, _, _ := m.Version()
		log.Printf("postgres database updated, version: %d", v)
	}

	return nil
}
