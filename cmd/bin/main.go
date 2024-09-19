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
	runLogExposer := flag.Bool("logexposer", false, "Menjalankan microservice: Log Exposer")
	runWhatsappClient := flag.Bool("whatsappclient", false, "Menjalankan microservice: Wahtsapp Client")

	flag.Parse()
	switch {
	case *runSeeder:
		cmd.Seed()

	case *runMigration:
		cmd.Migrate()

	case *runLogExposer:
		cmd.LogExposer()

	case *runWhatsappClient:
		cmd.WhatsappClient()

	default:
		cmd.App()
	}
}
