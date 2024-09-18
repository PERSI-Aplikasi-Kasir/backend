package cmd

import (
	"backend/database"
	"backend/database/migrations"
	"fmt"
)

func Migrate() {
	fmt.Println("Running migration...")

	database.InitializeDB()
	db := database.GetDBInstance()
	err := migrations.Migrate(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("✓ Migration process finished")
}
