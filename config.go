package golog

import (
	"os"
	"sync"
)

// Format is the format of the log.
type Format int

const (
	ReadableFormat Format = iota
	JSONFormat
)

// LogLevel is the level of the log.
type LogLevel string

const (
	NormalLog  LogLevel = "LOG" // NormalLog is the prefix for default log.
	SuccessLog          = "SUC" // SuccessLog is the prefix for success log.
	ErrorLog            = "ERR" // ErrorLog is the prefix for error log.
	WarningLog          = "WRR" // WarningLog is the prefix for warning log.
	DebugLog            = "DEB" // DebugLog is the prefix for debug log.
	InfoLog             = "INF" // InfoLog is the prefix for information log.
)

// Logs is the structure of the log.
type Logs struct {
	ID         string         // ID is the unique identifier of the log.
	Level      LogLevel       // Level is the level of the log.
	Time       string         // Time is the time when the log was created.
	Message    any            // Message is the content of the log.
	Additional map[string]any // Additional is the additional information of the log.
}

// Config is the configuration for the logger.
type Config struct {
	Destination    *os.File   // Destination is the file where the log will be written to.
	FileFormat     Format     // FileFormat is the format of the log file.
	TerminalFormat Format     // TerminalFormat is the format of the log printed to the terminal.
	mu             sync.Mutex // Ensures thread-safe writes to the log file.
	logChannel     chan Logs  // Channel to handle async log writing.
	doneChannel    chan bool  // Channel to signal log writing completion.
}
