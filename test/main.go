package main

import (
	"github.com/tvative/golog"
)

func main() {
	// Initialize logger instance
	logger := golog.Initialize()

	// Set log file destination
	logFilePath := "application"
	logger.SetFileOutput(logFilePath)
	logger.SetTerminalOutput(false)

	// Set log formats
	logger.SetFileFormat(golog.JSONFormat)
	logger.SetTerminalFormat(golog.JSONFormat)

	// Set log rotation
	logger.SetLogRotateByCount(10)

	// Set log outputs

	// Log messages
	for i := 0; i < 20; i++ {
		logger.Log(golog.NormalLog, map[string]any{"user": "Alice"}, "This is an info message", "This is another info message")
		logger.Log(golog.ErrorLog, map[string]any{"error": "file not found"}, "This is an error message")
		logger.Log(golog.SuccessLog, map[string]any{"transaction": "TX12345"}, "This is a success message")
		logger.Log(golog.WarningLog, map[string]any{"memory": "low"}, "This is a warning message")
		logger.Log(golog.DebugLog, map[string]any{"details": "debugging app"}, "This is a debug message")
		logger.Log(golog.InfoLog, nil, "%d This is another info message", 1, 10, 10.5, 'c', true, "aa")
	}
}
