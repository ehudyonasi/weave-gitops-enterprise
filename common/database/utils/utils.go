package utils

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/weaveworks/wks/common/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB ref contains a pointer to the MCCP database
var DB *gorm.DB

// Open creates the SQLite database or connects to an existing database
func Open(dbURI string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to connect to database")
	}
	// Set the global database Ref
	DB = db
	return db, nil
}

// MigrateTables creates the database tables given a gorm.DB
func MigrateTables(db *gorm.DB) error {
	// Migrate the schema
	err := db.AutoMigrate(&models.Event{})
	if err != nil {
		return errors.New("failed to create Events table")
	}
	log.Info("created Events table")

	err = db.AutoMigrate(&models.Cluster{})
	if err != nil {
		return errors.New("failed to create Clusters table")
	}
	log.Info("created Clusters table")

	err = db.AutoMigrate(&models.GitRepository{})
	if err != nil {
		return errors.New("failed to create GitRepository table")
	}
	log.Info("created GitRepository table")

	err = db.AutoMigrate(&models.GitProvider{})
	if err != nil {
		return errors.New("failed to create GitProviders table")
	}
	log.Info("created GitProviders table")

	err = db.AutoMigrate(&models.Workspace{})
	if err != nil {
		return errors.New("failed to create Workspaces table")
	}
	log.Info("created Workspaces table")
	return nil
}