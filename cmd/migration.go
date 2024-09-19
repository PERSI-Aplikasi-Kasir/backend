package cmd

import (
	"backend/database"
	"backend/database/migrations"
	"backend/pkg/env"
	"backend/pkg/logger"
	"fmt"

	"github.com/rs/zerolog/log"
)

func Migrate() {
	fmt.Println("Running migration...")

	logger.InitializeLogger(env.LogsPath + "migration.log")
	database.InitializeDB()
	db := database.GetDBInstance()
	err := migrations.Migrate(db)
	if err != nil {
		log.Error().Err(err).Msg("Error while migrating database")
		panic(err)
	}

	database.UnsyncDB()
	log.Info().Msg("✓ Migration process finished")
}
