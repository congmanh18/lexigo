package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "lexigo/model/china"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	var err error
	log.Println("🔍 DSN:", dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	if os.Getenv("MIGRATE") == "true" {
		// AutoMigrate (for dev only, use Goose in prod)
		if err := DB.AutoMigrate(&model.ChinaEntry{}); err != nil {
			log.Fatal("❌ Migration error:", err)
		}
	}

	log.Println("✅ Connected to PostgreSQL with GORM")
}
