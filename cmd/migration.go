package cmd

import (
	"backend/database"
	"backend/database/migrations"
	"fmt"

	"github.com/rs/zerolog/log"
)

func Migrate() {
	fmt.Println("Running migration...")

	database.InitializeDB()
	db := database.GetDBInstance()
	err := migrations.Migrate(db)
	if err != nil {
		log.Error().Err(err).Msg("Error while migrating database")
		panic(err)
	}

	database.UnsyncDB()

	fmt.Println("✓ Migration process finished")
}
