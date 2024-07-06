package golog

// Log logs a message with the specified level and information.
func (c *Config) Log(level LogLevel, additional map[string]any, msg ...any) {
	log := Logs{
		ID:         genID(),
		Level:      level,
		Time:       genTime(),
		Message:    msg,
		Additional: additional,
	}

	if c.isNeedTerminalOutput {
		c.dumpTerminal(log)
	}

	if c.isNeedFileOutput {
		c.dumpFile(log)
	}
}
