package database

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations executes all pending migrations from the specified directory.
// Returns the number of applied migrations or an error.
func RunMigrations(db *sql.DB, migrationsPath string) (int, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return 0, fmt.Errorf("failed to create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationsPath),
		"postgres",
		driver,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to create migrate instance: %w", err)
	}

	// Get current version before migration
	versionBefore, _, _ := m.Version()

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return 0, fmt.Errorf("failed to run migrations: %w", err)
	}

	// Get version after migration
	versionAfter, _, _ := m.Version()

	appliedCount := int(versionAfter - versionBefore)
	if err == migrate.ErrNoChange {
		appliedCount = 0
	}

	return appliedCount, nil
}

// GetMigrationVersion returns the current migration version.
func GetMigrationVersion(db *sql.DB) (uint, bool, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return 0, false, fmt.Errorf("failed to create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return 0, false, fmt.Errorf("failed to create migrate instance: %w", err)
	}

	return m.Version()
}
