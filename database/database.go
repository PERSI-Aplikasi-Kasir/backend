package database

import (
	"backend/pkg/env"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var databaseInstance *gorm.DB
var sqlDB *sql.DB

func InitializeDB() {
	fmt.Println("===== Initialize Database =====")
	db, err := connectDB()
	if err != nil {
		log.Error().Err(err).Msg("Error while connecting to database")
		panic(err)
	}

	databaseInstance = db

	sqlDB, err = db.DB()
	if err != nil {
		log.Error().Err(err).Msg("Error while getting database instance")
		panic(err)
	}

	fmt.Printf("connected to: %s\n", env.DBName)
}

func GetDBInstance() *gorm.DB {
	return databaseInstance
}

func connectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DBHost,
		env.DBPort,
		env.DBUser,
		env.DBPassword,
		env.DBName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("Error while connecting to database")
		return nil, err
	}

	return db, nil
}

func UnsyncDB() {
	if sqlDB != nil {
		if err := sqlDB.Close(); err != nil {
			log.Error().Err(err).Msg("Error while closing database connection")
			return
		}
	}

	databaseInstance = nil
	fmt.Println("✓ Database connection closed")
}
