package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/JuanPO17/myfirstgo/internal/config"
	"github.com/JuanPO17/myfirstgo/internal/model"
)

// Connect initializes the database connection based on the provided configuration.
func Connect(cfg *config.Config) (*gorm.DB, error) {
	var dialector gorm.Dialector
	switch cfg.DBDriver {
	case "sqlite":
		dialector = sqlite.Open(cfg.DSN)
	case "postgres":
		dialector = postgres.Open(cfg.DSN)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.DBDriver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connection successful.")
	return db, nil
}

// Migrate runs the auto-migration for the database models.
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.Task{})
}
