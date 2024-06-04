package database

import (
	"ewallet/config"
	"ewallet/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDb() (*gorm.DB, error) {
	cfg := config.AppConfig

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Printf("DB_USER: %s, DB_PASSWORD: %s, DB_NAME: %s", cfg.DBUser, cfg.DBPassword, cfg.DBName)
		log.Fatal("FAILED to connect to database. \n", err)
	}

	// Сообщение о том, что мы подключились
	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	// Воспользуемся функцией автоматической миграции, чтобы создать таблицы из нашей модели
	log.Println("running migrations")
	db.AutoMigrate(&models.Wallet{})
	db.AutoMigrate(&models.Transaction{})

	return db, nil
}
