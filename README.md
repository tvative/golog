<div align="center">
  <h1>Go Log</h1>
  <p>A Simple and Minimalist Logger for Go Language</p>
</div>

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/github.com/tvative/golog)

`golog` is a minimalist package for logging in Go. It provides a simple and easy-to-use interface for logging in Go.

## Features

- **Simple**: `golog` is a simple and minimalist package for logging in Go.
- **Log to File**: `golog` supports logging to a file.
- **Log to Terminal**: `golog` supports logging to the terminal.
- **JSON Format**: `golog` supports JSON format for logs.
- **Readable Format**: `golog` supports a readable format for logs.
- **Log Rotation**: `golog` supports log rotation by count.
- **Log Levels**: `golog` supports different log levels.

## Usage

### Installation

To use `golog` in your Go project, you need to install it using Go modules. You can add it to your project with the
following command:

```bash
go get -u github.com/tvative/golog
```

### Basic Usage

To use `golog`, you need to create a new instance of the logger and then use the `Log` method to log messages. Here is a
simple example of how to use `golog`:

```go
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
	logger.Log(golog.NormalLog, map[string]any{"user": "Alice"}, "User login attempt", "Login attempt successful for user: Alice")
	logger.Log(golog.ErrorLog, map[string]any{"error": "file not found", "path": "/var/www/html/index.html"}, "File not found error", "Attempt to access non-existent file: /var/www/html/index.html")
	logger.Log(golog.SuccessLog, map[string]any{"transaction": "TX12345", "amount": 150.75}, "Transaction successful", "Transaction ID TX12345 completed successfully with amount $150.75")
	logger.Log(golog.WarningLog, map[string]any{"memory": "low", "server_id": "srv-01"}, "Low memory warning", "Server srv-01 is experiencing low memory conditions")
	logger.Log(golog.DebugLog, map[string]any{"details": "request processing", "request_id": "req-67890"}, "Debugging request", "Processing request ID req-67890 for debugging purposes")
}
```

**File Output:**

```json
{"ID":"1720245990408132641","Level":"LOG","Time":"2024-07-06 11:36:30:408:408136394","Message":"User login attempt Login attempt successful for user: Alice","Additional":{"user":"Alice"}}
{"ID":"1720245990408465947","Level":"ERR","Time":"2024-07-06 11:36:30:408:408468440","Message":"File not found error Attempt to access non-existent file: /var/www/html/index.html","Additional":{"error":"file not found","path":"/var/www/html/index.html"}}
{"ID":"1720245990408550091","Level":"SUC","Time":"2024-07-06 11:36:30:408:408551195","Message":"Transaction successful Transaction ID TX12345 completed successfully with amount $150.75","Additional":{"amount":150.75,"transaction":"TX12345"}}
{"ID":"1720245990408598067","Level":"WRR","Time":"2024-07-06 11:36:30:408:408599492","Message":"Low memory warning Server srv-01 is experiencing low memory conditions","Additional":{"memory":"low","server_id":"srv-01"}}
{"ID":"1720245990408645649","Level":"DEB","Time":"2024-07-06 11:36:30:408:408646453","Message":"Debugging request Processing request ID req-67890 for debugging purposes","Additional":{"details":"request processing","request_id":"req-67890"}}
```

**Terminal Output:**

```json
2024-07-06 11:36:30:408:408136394 [ LOG ] User login attempt Login attempt successful for user: Alice {"user":"Alice"}
2024-07-06 11:36:30:408:408468440 [ ERR ] File not found error Attempt to access non-existent file: /var/www/html/index.html {"error":"file not found","path":"/var/www/html/index.html"}
2024-07-06 11:36:30:408:408551195 [ SUC ] Transaction successful Transaction ID TX12345 completed successfully with amount $150.75 {"amount":150.75,"transaction":"TX12345"}
2024-07-06 11:36:30:408:408599492 [ WRR ] Low memory warning Server srv-01 is experiencing low memory conditions {"memory":"low","server_id":"srv-01"}
2024-07-06 11:36:30:408:408646453 [ DEB ] Debugging request Processing request ID req-67890 for debugging purposes {"details":"request processing","request_id":"req-67890"}
```

### Example

For a minimal example, see the [test/main.go](test/main.go) file. That file contains a simple example of how to use the
package.

## License

This project is licensed under the MIT License; see the [LICENSE](LICENSE) file for details.
