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
	logger.SetTerminalOutput(false)

	// Set log formats
	logger.SetFileFormat(golog.JSONFormat)
	logger.SetTerminalFormat(golog.JSONFormat)

	// Set log rotation
	logger.SetLogRotateByCount(10)

	// Log messages
	logger.Log(golog.NormalLog, map[string]any{"user": "Alice"}, "This is an info message", "another content")
}

```

### Example

For a minimal example, see the [test/main.go](test/main.go) file. That file contains a simple example of how to use the
package.

## License

This project is licensed under the MIT License; see the [LICENSE](LICENSE) file for details.
