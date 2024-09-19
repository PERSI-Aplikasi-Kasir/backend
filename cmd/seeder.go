package cmd

import (
	"backend/database"
	"backend/database/seeder"
	"backend/pkg/env"
	"backend/pkg/logger"
	"fmt"

	"github.com/rs/zerolog/log"
)

func Seed() {
	fmt.Println("Running seeder...")

	logger.InitializeLogger(env.LogsPath + "seeder.log")
	database.InitializeDB()
	db := database.GetDBInstance()
	seeder.Seeder(db)

	database.UnsyncDB()
	log.Info().Msg("✓ Seed process finished")
}
