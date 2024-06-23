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

	c.dumpTerminal(log)
	c.logChannel <- log
}
