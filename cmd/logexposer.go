package cmd

import (
	"backend/microservices/logexposer"
	"backend/pkg/env"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func LogExposer() {
	fmt.Println("Running microservice: LogExposer")

	logexposer.InitializeLogExposer()

	if err := http.ListenAndServe(env.BEHost+":"+env.LoggerPort, nil); err != nil && err != http.ErrServerClosed {
		fmt.Println("err")
		fmt.Println(err)

		log.Fatal().Err(err).Msg("Microservice: LogExposer server failed to start")
	}
}
