package main

import (
	"log"
	"time"

	"github.com/tvative/golog"
)

func main() {
	// Initialize logger instance
	logger := golog.Initialize()
	defer logger.Close()

	// Set log file destination
	logFilePath := "application.log"
	logger.SetDestination(logFilePath)

	// Set log formats
	logger.SetFileFormat(golog.JSONFormat)
	logger.SetTerminalFormat(golog.JSONFormat)

	// Log messages
	logger.Log(golog.NormalLog, map[string]any{"user": "Alice"}, "This is an info message", "This is another info message")
	logger.Log(golog.ErrorLog, map[string]any{"error": "file not found"}, "This is an error message")
	logger.Log(golog.SuccessLog, map[string]any{"transaction": "TX12345"}, "This is a success message")
	logger.Log(golog.WarningLog, map[string]any{"memory": "low"}, "This is a warning message")
	logger.Log(golog.DebugLog, map[string]any{"details": "debugging app"}, "This is a debug message")
	logger.Log(golog.InfoLog, nil, "%d This is another info message", 1, 10, 10.5, 'c', true, "aa")

	// Give some time for logs to be written asynchronously
	time.Sleep(1 * time.Second)

	log.Println("Logs have been written to the terminal and the file.")
}
