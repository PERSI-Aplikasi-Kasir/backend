package cmd

import (
	"backend/pkg/env"
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

func Logger() {
	fmt.Println("Running microservice: logger")

	http.HandleFunc("/", getLogs)
	fmt.Println(env.BEHost)
	fmt.Println(env.LoggerPort)

	log.Info().Msgf("✓ Microservice: logger is running on %s:%s", env.BEHost, env.LoggerPort)

	if err := http.ListenAndServe(env.BEHost+":"+env.LoggerPort, nil); err != nil && err != http.ErrServerClosed {
		fmt.Println("err")
		fmt.Println(err)

		log.Fatal().Err(err).Msg("Microservice: logger server failed to start")
	}
}

type LogEntry struct {
	Level   string `json:"level"`
	Time    string `json:"time"`
	Caller  string `json:"caller"`
	Message string `json:"message"`
}

type PaginatedResponse struct {
	Data  []LogEntry `json:"data"`
	Page  int        `json:"page"`
	Limit int        `json:"limit"`
	Total int        `json:"total"`
}

func readLogs() ([]LogEntry, error) {
	var logs []LogEntry

	file, err := os.Open(env.LogsPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var logEntry LogEntry
		err := json.Unmarshal(scanner.Bytes(), &logEntry)
		if err != nil {
			continue
		}
		logs = append(logs, logEntry)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return logs, nil
}

func paginateLogs(logs []LogEntry, page, limit int) []LogEntry {
	start := len(logs) - (page * limit)
	if start < 0 {
		start = 0
	}

	end := start + limit
	if end > len(logs) {
		end = len(logs)
	}

	return logs[start:end]
}

func getLogs(w http.ResponseWriter, r *http.Request) {
	logs, err := readLogs()
	if err != nil {
		http.Error(w, "Error reading logs", http.StatusInternalServerError)
		return
	}

	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	paginatedLogs := paginateLogs(logs, page, limit)

	response := PaginatedResponse{
		Data:  paginatedLogs,
		Page:  page,
		Limit: limit,
		Total: len(logs),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
