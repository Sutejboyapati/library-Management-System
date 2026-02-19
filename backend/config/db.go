package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// InitDB connects to MySQL using your .env settings.
func InitDB() error {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "library_db"
	}

	log.Printf("DB: host=%s user=%s database=%s", host, user, name)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, pass, host, name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("database connection: %w", err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err := db.Ping(); err != nil {
		log.Println("MySQL connection failed. Check:")
		log.Println("  1) MySQL server is running")
		log.Println("  2) .env has correct DB_USER, DB_PASSWORD, DB_NAME")
		log.Println("  3) Database exists (e.g. CREATE DATABASE library_db;)")
		return fmt.Errorf("database ping: %w", err)
	}

	log.Println("Connected to database")
	DB = db

	// Auto-create tables if they don't exist
	if err := initSchema(db); err != nil {
		return fmt.Errorf("init schema: %w", err)
	}

	return nil
}

// initSchema creates tables if they don't exist.
func initSchema(db *sql.DB) error {
	stmt := []string{
		`CREATE TABLE IF NOT EXISTS Users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(50) NOT NULL DEFAULT 'user',
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS Books (
			id INT AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(500) NOT NULL,
			author VARCHAR(255) NOT NULL,
			isbn VARCHAR(50),
			genre VARCHAR(100),
			language VARCHAR(100),
			shelf_number VARCHAR(50),
			available_copies INT NOT NULL DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS borrowingrecords (
			id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			book_id INT NOT NULL,
			borrowed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			due_date TIMESTAMP NOT NULL,
			returned_at TIMESTAMP NULL,
			FOREIGN KEY (user_id) REFERENCES Users(id),
			FOREIGN KEY (book_id) REFERENCES Books(id)
		)`,
	}
	for _, s := range stmt {
		if _, err := db.Exec(s); err != nil {
			return fmt.Errorf("create table: %w", err)
		}
	}
	log.Println("Schema ready")
	return nil
}
