package database

import (
	"database/sql"
	"ewallet/internal/config"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Добавляем этот импорт
)

var db *sql.DB

func ConnectDb() (*sql.DB, error) {
	cfg := config.AppConfig

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("FAILED to connect to database: %v", err)
	}

	log.Println("connected")

	if err := migrateDB(); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	return db, nil
}

func migrateDB() error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS wallets (
		id VARCHAR(36) PRIMARY KEY,
		balance NUMERIC(15, 2)
	)`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS transactions (
		id VARCHAR(36) PRIMARY KEY,
		time TIMESTAMP,
		from_wallet VARCHAR(36),
		to_wallet VARCHAR(36),
		amount NUMERIC(15, 2)
	)`)
	if err != nil {
		return err
	}

	return nil
}
