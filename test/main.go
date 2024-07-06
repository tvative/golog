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
	logger.SetTerminalOutput(true)

	// Set log formats
	logger.SetFileFormat(golog.JSONFormat)
	logger.SetTerminalFormat(golog.ReadableFormat)

	// Set log rotation
	logger.SetLogRotateByCount(10)

	// Log messages
	for i := 0; i < 20; i++ {
		logger.Log(golog.NormalLog, map[string]any{"user": "Alice"}, "User login attempt", "Login attempt successful for user: Alice")
		logger.Log(golog.ErrorLog, map[string]any{"error": "file not found", "path": "/var/www/html/index.html"}, "File not found error", "Attempt to access non-existent file: /var/www/html/index.html")
		logger.Log(golog.SuccessLog, map[string]any{"transaction": "TX12345", "amount": 150.75}, "Transaction successful", "Transaction ID TX12345 completed successfully with amount $150.75")
		logger.Log(golog.WarningLog, map[string]any{"memory": "low", "server_id": "srv-01"}, "Low memory warning", "Server srv-01 is experiencing low memory conditions")
		logger.Log(golog.DebugLog, map[string]any{"details": "request processing", "request_id": "req-67890"}, "Debugging request", "Processing request ID req-67890 for debugging purposes")
		logger.Log(golog.InfoLog, nil, "150 concurrent users currently online")
	}
}
