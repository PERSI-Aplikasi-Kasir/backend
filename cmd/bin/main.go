package main

import (
	"backend/cmd"
	"backend/pkg/env"
	"backend/pkg/logger"
	"flag"
)

func main() {
	env.InitializeEnv()
	logger.InitializeLogger(env.LogsPath + "app.log")

	runSeeder := flag.Bool("seed", false, "Menjalankan seed")
	runMigration := flag.Bool("migrate", false, "Menjalankan migration")
	runLogExposer := flag.Bool("logexposer", false, "Menjalankan microservice: logger")
	flag.Parse()

	switch {
	case *runSeeder:
		cmd.Seed()

	case *runMigration:
		cmd.Migrate()

	case *runLogExposer:
		cmd.LogExpose()

	default:
		cmd.App()
	}
}
