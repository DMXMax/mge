// Package storage provides shared game data structures and operations
package storage

import (
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDatabase initializes a SQLite database connection at the specified path.
// It creates the database directory if it doesn't exist and returns a configured
// GORM database instance with silent logging.
//
// Parameters:
//   - dbPath: The path to the SQLite database file
//
// Returns:
//   - *gorm.DB: The configured database connection
//   - error: Any error that occurred during initialization
func InitDatabase(dbPath string) (*gorm.DB, error) {
	// Ensure the database directory exists
	if err := os.MkdirAll(filepath.Dir(dbPath), 0755); err != nil {
		return nil, err
	}

	// Open SQLite database connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

