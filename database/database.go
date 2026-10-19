package database

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect membuka koneksi GORM ke PostgreSQL memakai environment variable DATABASE_URL
func Connect() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, membaca dari OS environment")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, errors.New("environment variable DATABASE_URL belum diatur")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Info = tampilkan setiap SQL yang dibuat GORM (bagus untuk belajar).
		// Ganti ke logger.Warn jika terminal terlalu ramai.
		Logger:               logger.Default.LogMode(logger.Info),
		DisableAutomaticPing: true, // ping dilakukan manual di bawah, dengan batas waktu
	})
	if err != nil {
		return nil, err
	}
	// Ping untuk memastikan database benar-benar bisa dihubungi
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		sqlDB.Close()
		return nil, err
	}
	return db, nil
}
