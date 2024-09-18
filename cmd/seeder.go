package cmd

import (
	"backend/database"
	"backend/database/seeder"
	"fmt"
)

func Seed() {
	fmt.Println("Running seeder...")

	database.InitializeDB()
	db := database.GetDBInstance()
	seeder.Seeder(db)

	fmt.Println("✓ Seed process finished")
}
