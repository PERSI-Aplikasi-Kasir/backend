package cmd

import (
	"backend/microservices/logexposer"
	"backend/pkg/env"
	"backend/pkg/logger"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func LogExposer() {
	fmt.Println("Running microservice: LogExposer")

	logger.InitializeLogger(env.LogsPath + "logexposer.log")
	logexposer.InitializeLogExposer()

	err := http.ListenAndServe(env.BEHost+":"+env.LoggerPort, nil)
	if err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("Microservice: LogExposer server failed to start")
		panic(err)
	}
}
