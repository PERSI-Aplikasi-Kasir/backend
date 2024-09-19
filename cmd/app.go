package cmd

import (
	"backend/database"
	"backend/internal/config/router"
	"backend/internal/integration/mailer"
	"backend/pkg/env"
	"backend/pkg/logger"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/rs/zerolog/log"
)

func App() {
	fmt.Println("Running app...")

	logger.InitializeLogger("logs/app.log")

	database.InitializeDB()
	router.InitializeRouter()
	router.InitializeRoutes()
	mailer.InitializeMailer()

	server := &http.Server{
		Addr:    env.BEHost + ":" + env.BEPort,
		Handler: router.GetRouterInstance(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server failed to start")
		}
	}()

	log.Info().Msg("✓ App is running")

	shutdownListener(server)
}

func shutdownListener(server *http.Server) {
	quit := make(chan os.Signal, 1)

	shutdownSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
	if runtime.GOOS == "windows" {
		shutdownSignals = []os.Signal{os.Interrupt}
	}

	signal.Notify(quit, shutdownSignals...)
	<-quit
	log.Info().Msg("Server is shutting down...")

	database.UnsyncDB()
	router.UnsyncRouter(server)
	logger.UnsyncLogger()

	log.Info().Msg("✓ Server shutted down")
}
