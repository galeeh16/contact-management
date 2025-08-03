package config

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_TIMEZONE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger:      logger.Default.LogMode(logger.Info), // kalo di prod dimatiin aja
		PrepareStmt: true,
	})

	if err != nil {
		return nil, errors.New("Gagal membuat koneksi ke database, err: " + err.Error())
	}

	// DB = db

	return db, nil
}
