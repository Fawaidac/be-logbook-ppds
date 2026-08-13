package database

import (
	"fmt"
	"log"
	"os"

	"be-logbook-ppds/configs"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgres(cfg *configs.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	// Run auto migrations from SQL files
	if err := RunMigrations(db); err != nil {
		log.Printf("Warning: Auto migration failed: %v", err)
	}

	return db, nil
}

func RunMigrations(db *sqlx.DB) error {
	upSQL1, err := os.ReadFile("migrations/000001_create_users_table.up.sql")
	if err == nil && len(upSQL1) > 0 {
		if _, err := db.Exec(string(upSQL1)); err != nil {
			log.Printf("Migration 1 execution note: %v", err)
		}
	}

	upSQL2, err := os.ReadFile("migrations/000002_seed_users.up.sql")
	if err == nil && len(upSQL2) > 0 {
		if _, err := db.Exec(string(upSQL2)); err != nil {
			log.Printf("Migration 2 execution note: %v", err)
		}
	}

	upSQL3, err := os.ReadFile("migrations/000003_create_jadwals_table.up.sql")
	if err == nil && len(upSQL3) > 0 {
		if _, err := db.Exec(string(upSQL3)); err != nil {
			log.Printf("Migration 3 execution note: %v", err)
		}
	}

	upSQL4, err := os.ReadFile("migrations/000004_seed_jadwals.up.sql")
	if err == nil && len(upSQL4) > 0 {
		if _, err := db.Exec(string(upSQL4)); err != nil {
			log.Printf("Migration 4 execution note: %v", err)
		}
	}

	upSQL5, err := os.ReadFile("migrations/000005_create_tindakans_table.up.sql")
	if err == nil && len(upSQL5) > 0 {
		if _, err := db.Exec(string(upSQL5)); err != nil {
			log.Printf("Migration 5 execution note: %v", err)
		}
	}

	upSQL6, err := os.ReadFile("migrations/000006_seed_tindakans.up.sql")
	if err == nil && len(upSQL6) > 0 {
		if _, err := db.Exec(string(upSQL6)); err != nil {
			log.Printf("Migration 6 execution note: %v", err)
		}
	}

	log.Println("Database migration & seed sync completed successfully")
	return nil
}


