package database

import (
	"fmt"
	"log"
	"sesi7-gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST = "localhost"
	DB_USER = "root"
	DB_PASS = "root"
	DB_NAME = "db_go_sql"
	DB_PORT = 5432
)

func StartDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)
	fmt.Println("dsn:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Default().Println("Connection db success")

	err = migration(db)
	if err != nil {
		panic(err)
	}
	return db
}

func migration(db *gorm.DB) error {
	if err := db.AutoMigrate(models.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(models.Product{}); err != nil {
		return err
	}

	return nil
}
