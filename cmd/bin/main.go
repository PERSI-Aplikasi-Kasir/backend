package main

import (
	"backend/cmd"
	"backend/pkg/env"
	"flag"
)

func main() {
	env.InitializeEnv()

	runSeeder := flag.Bool("seed", false, "Menjalankan seed")
	runMigration := flag.Bool("migrate", false, "Menjalankan migration")
	flag.Parse()

	switch {
	case *runSeeder:
		cmd.Seed()

	case *runMigration:
		cmd.Migrate()

	default:
		cmd.App()
	}
}
