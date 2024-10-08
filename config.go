package golog

import (
	"os"
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
	fileDestination        string   // fileDestination is the path where the log file will be written to.
	fileDescriptor         *os.File // fileDescriptor is the file where the log will be written to.
	fileFormat             Format   // fileFormat is the format of the log file.
	terminalFormat         Format   // terminalFormat is the format of the log printed to the terminal.
	isNeedFileOutput       bool     // isNeedFileOutput is the flag to determine if the log needs to be written to the file.
	isNeedTerminalOutput   bool     // isNeedTerminalOutput is the flag to determine if the log needs to be printed to the terminal.
	logCount               int64    // logCount is the number of logs.
	logRotateCount         int64    // logRotateCount is the number of logs before rotating the log file.
	isNeedLogRotateByCount bool     // isNeedLogRotateByCount is the flag to determine if the log file needs to be rotated.
}
