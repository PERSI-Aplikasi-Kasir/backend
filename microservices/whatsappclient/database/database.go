package database

import (
	"backend/pkg/env"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/rs/zerolog/log"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

var databaseInstance *sqlstore.Container

func InitializeWAClientDB() {
	fmt.Println("===== Initialize Database =====")
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.DBUser,
		url.QueryEscape(env.DBPassword),
		env.DBHost,
		env.DBPort,
		env.WAClientDBName,
	)

	container, err := sqlstore.New("pgx", dsn, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to initialize whatsapp client database")
		panic(fmt.Errorf("failed to initialize sqlstore: %w", err))
	}

	databaseInstance = container
	fmt.Printf("connected to: %s\n", env.WAClientDBName)
}

func GetDBInstance() *sqlstore.Container {
	return databaseInstance
}

func UnsyncDB() {
	if databaseInstance != nil {
		err := databaseInstance.Close()
		if err != nil {
			log.Error().Err(err).Msg("Error while closing database connection")
		} else {
			fmt.Println("===== Database connection closed gracefully =====")
		}
	}

	databaseInstance = nil
	fmt.Println("✓ Database connection closed")
}

func ResyncDB() {
	UnsyncDB()
	InitializeWAClientDB()
}
