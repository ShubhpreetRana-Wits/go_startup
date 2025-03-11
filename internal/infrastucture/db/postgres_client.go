package db

import (
	"fmt"

	"example.com/startup/internal/domain/entities"
	"example.com/startup/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresClient struct {
	DB *gorm.DB
}

func NewGormDBConfig(cfg *config.Config) (*PostgresClient, error) {

	// Construct the DSN (Data Source Name) for the PostgreSQL connection
	dsn := buildDSN(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)

	// Open the database connection
	db, err := openDatabaseConnection(dsn)
	if err != nil {
		// If an error occurs, return nil and the error
		return nil, err
	}
	runMigrations(db)

	// Return the GormDb instance with the established database connection
	return &PostgresClient{DB: db}, nil
}

func buildDSN(host, port, user, password, dbname, sslmode string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)
}

func openDatabaseConnection(dsn string) (*gorm.DB, error) {
	// Attempt to open the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// If an error occurs during the connection, return nil and the error
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	// Return the database connection if successful
	return db, nil
}

// Close closes the database connection
func (p *PostgresClient) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// runMigrations runs database migrations
func runMigrations(db *gorm.DB) error {
	// Auto-migrate the entities
	return db.AutoMigrate(&entities.GeneratedUrl{})
}
