package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	// create database
	psqlInfo := "host=localhost user=postgres password=admin sslmode=disable TimeZone=Asia/Kolkata"
	dbName := os.Getenv("DB_NAME")

	// Connect to the default database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the database exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbName).Scan(&exists)
	if err != nil {
		log.Println("Error checking database existence:", err)
	}

	// Create the database if it doesn't exist
	if !exists {
		_, err = db.Exec("CREATE DATABASE " + dbName)
		if err != nil {
			return nil, fmt.Errorf("error creating database: %w", err)
		}
	}

	// Connect to the new database with GORM
	dsn := fmt.Sprintf("%s dbname=%s port=5432", psqlInfo, dbName)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Run migrations
	err = gormDB.AutoMigrate(&User{}, &Tiger{}, &Sighting{})
	if err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return gormDB, nil
}
