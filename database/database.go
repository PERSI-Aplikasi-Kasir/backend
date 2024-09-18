package database

import (
	"backend/pkg/env"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var databaseInstance *gorm.DB

func InitializeDB() {
	fmt.Println("===== Initialize Database =====")
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	databaseInstance = db
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
		return nil, err
	}

	return db, nil
}
