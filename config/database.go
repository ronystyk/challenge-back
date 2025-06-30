package config

import (
	"challenge/models"
	"log"
	"os"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	if err := db.AutoMigrate(&models.Stock{}); err != nil {
		log.Fatal("Failed to migrate DB:", err)
	}

	return db
}

func InitTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}

	// Migraciones de prueba
	if err := db.AutoMigrate(&models.Stock{}); err != nil {
		t.Fatalf("Failed to migrate test DB: %v", err)
	}

	return db
}
