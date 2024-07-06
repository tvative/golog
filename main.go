package golog

import (
	"fmt"
	"os"
)

// Initialize initializes the logger.
func Initialize() *Config {
	c := &Config{
		fileDescriptor:       nil,
		fileFormat:           ReadableFormat,
		terminalFormat:       ReadableFormat,
		isNeedFileOutput:     false,
		isNeedTerminalOutput: true,
		logCount:             0,
		logRotateCount:       1000,
	}

	return c
}

// SetFileFormat sets the format of the log file.
func (c *Config) SetFileFormat(format Format) {
	c.fileFormat = format
}

// SetTerminalFormat sets the format of the log printed to the terminal.
func (c *Config) SetTerminalFormat(format Format) {
	c.terminalFormat = format
}

// SetTerminalOutput sets the flag to determine if the log needs to be printed to the terminal.
func (c *Config) SetTerminalOutput(status bool) {
	c.isNeedTerminalOutput = status
}

// SetFileOutput sets the flag to determine if the log needs to be written to the file.
func (c *Config) SetFileOutput(destination string) {
	if destination == "" {
		c.isNeedFileOutput = false
		return
	}

	genDestination := fmt.Sprintf("%s_%s.log", destination, genRotatePostfix())
	descriptor, err := os.Create(genDestination)
	if err != nil {
		panic(fmt.Sprintf("unable to create the log file: %v", err))
		return
	}

	c.fileDestination = destination
	c.fileDescriptor = descriptor
	c.isNeedFileOutput = true
}

// SetLogRotateByCount sets the number of logs before rotating the log file.
func (c *Config) SetLogRotateByCount(count int64) {
	c.logRotateCount = count
	c.isNeedLogRotateByCount = true
}
