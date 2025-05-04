package migration

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // This is important for the postgres driver
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Up() {
	// Replace with your actual DB URL
	dbURL := "postgres://postgres:Awaisniaz@localhost:5432/user-management-system?sslmode=disable"

	// Create a new migration instance
	m, err := migrate.New("file://scripts/migrations", dbURL)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && err.Error() != "no change" {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migration completed successfully.")
}

func Down() {
	dbURL := "postgres://postgres:Awaisniaz@localhost:5432/user-management-system?sslmode=disable"
	migrationsPath := "file://scripts/migrations"

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	// Rollback one step (safe)
	if err := m.Steps(-1); err != nil && err.Error() != "no change" {
		log.Fatalf("Migration rollback failed: %v", err)
	}

	fmt.Println("Rolled back one migration step successfully.")
}
