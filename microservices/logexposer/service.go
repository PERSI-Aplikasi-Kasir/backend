package logexposer

import (
	"backend/pkg/env"
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type LogExposer struct {
	path string
}

func FilePath(path string) *LogExposer {
	return &LogExposer{path: path}
}

func (l *LogExposer) GetLogs(w http.ResponseWriter, r *http.Request) {
	logs, err := readLogs(l.path)
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

func readLogs(logsPath string) ([]LogEntry, error) {
	var logs []LogEntry

	file, err := os.Open(logsPath)
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

func getLogFiles() *[]string {
	files, err := os.ReadDir(env.LogsPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	var logFiles []string
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".log" {
			logFiles = append(logFiles, file.Name())
		}
	}

	return &logFiles
}
