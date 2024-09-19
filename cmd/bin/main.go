package main

import (
	"backend/cmd"
	"backend/pkg/env"
	"backend/pkg/logger"
	"flag"
)

func main() {
	env.InitializeEnv()
	logger.InitializeLogger(env.LogsPath)

	runSeeder := flag.Bool("seed", false, "Menjalankan seed")
	runMigration := flag.Bool("migrate", false, "Menjalankan migration")
	runLogger := flag.Bool("logger", false, "Menjalankan microservice: logger")
	flag.Parse()

	switch {
	case *runSeeder:
		cmd.Seed()

	case *runMigration:
		cmd.Migrate()

	case *runLogger:
		cmd.Logger()

	default:
		cmd.App()
	}
}
