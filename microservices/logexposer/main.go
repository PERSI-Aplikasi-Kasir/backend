package logexposer

import (
	"backend/pkg/env"
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func InitializeLogExposer() {
	logFiles := GetLogFiles()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(logFiles)
	})

	if logFiles != nil {
		for _, file := range *logFiles {
			http.HandleFunc("/"+file, FilePath(env.LogsPath+file).GetLogs)
		}
	}

	log.Info().Msgf("✓ Microservice: LogExposer is running on %s:%s", env.BEHost, env.LoggerPort)
}
