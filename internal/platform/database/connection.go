package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Config struct {
	URL        string
	MaxRetries int
	RetryDelay time.Duration
}

// NewConfig creates a new database configuration
func NewConfig(url string) *Config {
	return &Config{
		URL:        url,
		MaxRetries: 5,
		RetryDelay: 5 * time.Second,
	}
}

// ConnectWithRetry attempts to connect to the database with retries
func ConnectWithRetry(cfg *Config) (*sql.DB, error) {
	var (
		db  *sql.DB
		err error
	)

	for i := 0; i < cfg.MaxRetries; i++ {
		db, err = sql.Open("postgres", cfg.URL)
		if err != nil {
			log.Printf("Failed to open database connection: %v", err)
			time.Sleep(cfg.RetryDelay)
			continue
		}

		err = db.Ping()
		if err == nil {
			log.Printf("Successfully connected to database on attempt %d", i+1)
			return db, nil
		}

		log.Printf("Failed to ping database on attempt %d: %v", i+1, err)
		db.Close()
		time.Sleep(cfg.RetryDelay)
	}

	return nil, fmt.Errorf("failed to connect to database after %d attempts", cfg.MaxRetries)
}

// Initialize handles the complete database initialization process
func Initialize(cfg *Config) (*sql.DB, error) {
	log.Println("Starting database initialization...")

	// Connect to database with retry mechanism
	db, err := ConnectWithRetry(cfg)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %v", err)
	}

	// Run migrations using the same connection
	if err := RunMigrations(db); err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}

	return db, nil
}
