package logexposer

import (
	"backend/pkg/env"
	"log"
	"os"
	"path/filepath"
)

func GetLogFiles() *[]string {
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
