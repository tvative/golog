package golog

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Initialize initializes the logger.
func Initialize() *Config {
	c := &Config{
		Destination:    nil,
		FileFormat:     ReadableFormat,
		TerminalFormat: ReadableFormat,
		logChannel:     make(chan Logs, 1),
		doneChannel:    make(chan bool),
	}

	go c.dumpFile()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChannel
		fmt.Println("Received shutdown signal, waiting for log processing to complete...")
		close(c.doneChannel)        // Signal the dumpFile goroutine to finish
		time.Sleep(1 * time.Second) // Give some time to process remaining logs
		close(c.logChannel)         // Close the logChannel
	}()

	return c
}

// SetDestination sets the destination of the log file.
func (c *Config) SetDestination(path string) {
	descriptor, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	c.Destination = descriptor
}

// SetFileFormat sets the format of the log file.
func (c *Config) SetFileFormat(format Format) {
	c.FileFormat = format
}

// SetTerminalFormat sets the format of the log printed to the terminal.
func (c *Config) SetTerminalFormat(format Format) {
	c.TerminalFormat = format
}

// Close closes the logger.
func (c *Config) Close() {
	close(c.logChannel)
	<-c.doneChannel
	if c.Destination != nil {
		err := c.Destination.Close()
		if err != nil {
			return
		}
	}
}
