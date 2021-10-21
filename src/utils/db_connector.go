package utils

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnector interface {
	Open()
	Close()
	GetDB() *gorm.DB
}

func DBConnectorHandler() DBConnector {
	return &dbConnector{}
}

type dbConnector struct {
	db *gorm.DB
}

func (dbConnector *dbConnector) Open() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	dbConnector.db = db
}

func (dbConnector *dbConnector) Close() {
	postgresDB, err := dbConnector.db.DB()

	if err != nil {
		panic("failed to connect database")
	}

	postgresDB.Close()
}

func (dbConnector *dbConnector) GetDB() *gorm.DB {
	return dbConnector.db
}
