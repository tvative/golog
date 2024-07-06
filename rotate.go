package golog

import (
	"fmt"
	"os"
)

// logRotateByCount rotates the log file by count.
func (c *Config) logRotateByCount() {
	if c.isNeedLogRotateByCount {
		if c.logCount >= c.logRotateCount {
			if c.fileDescriptor != nil {
				err := c.fileDescriptor.Close()
				if err != nil {
					return
				}

				c.logCount = 0
			}

			destination := fmt.Sprintf("%s_%s.log", c.fileDestination, genRotatePostfix())
			descriptor, err := os.Create(destination)
			if err != nil {
				tempFile, _ := os.Create("log_error")
				_, err := tempFile.WriteString(fmt.Sprintf("unable to create the log rotation file: %v", err))
				if err != nil {
					return
				}

				return
			}

			c.fileDescriptor = descriptor
		}
	}

	c.logCount++
}
