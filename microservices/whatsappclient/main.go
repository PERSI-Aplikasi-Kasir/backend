package whatsappclient

import (
	"backend/microservices/whatsappclient/config"
	"backend/microservices/whatsappclient/database"
	"backend/microservices/whatsappclient/router"
	"backend/pkg/env"
	"backend/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/rs/zerolog/log"
)

func InitializeWhatsappClient() {
	database.InitializeWAClientDB()
	config.InitializeClient()
	router.InitializeRouter()
	router.InitializeRoutes()

	server := &http.Server{
		Addr:    env.BEHost + ":" + env.WAClientPort,
		Handler: router.GetRouterInstance(),
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Whatsapp Client service failed to start")
		}
	}()

	log.Info().Msgf("✓ Microservice: Whatsapp Client is running on %s:%s", env.BEHost, env.WAClientPort)
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
	config.UnsyncClient()

	log.Info().Msg("✓ Server shutted down")
}
