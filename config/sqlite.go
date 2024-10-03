package config

import (
	"os"

	"github.com/marciopriebe/crud-go.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	// Check if the database file existis
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		// Create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}

		defer file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening error: %w", err)
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	// Return the DB
	return db, nil
}
