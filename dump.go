package golog

import (
	"encoding/json"
	"fmt"
	"strings"
)

func retReadable(log Logs, isTerminal bool) string {
	var logLevel string
	if isTerminal {
		logLevel = getColor(log.Level)
	} else {
		logLevel = fmt.Sprint(log.Level)
	}

	var msg string
	switch v := log.Message.(type) {
	case []string:
		msg = strings.Join(v, " ")
	case string:
		msg = v
	case []int:
		intSTRs := make([]string, len(v))
		for i, n := range v {
			intSTRs[i] = fmt.Sprintf("%d", n)
		}

		msg = strings.Join(intSTRs, " ")
	case []interface{}:
		anySTRs := make([]string, len(v))
		for i, elem := range v {
			anySTRs[i] = fmt.Sprintf("%v", elem)
		}

		msg = strings.Join(anySTRs, " ")
	default:
		msg = fmt.Sprintf("%v", v)
	}

	if log.Additional != nil {
		jsonAdditional, _ := json.Marshal(log.Additional)
		return fmt.Sprintf("%-33s [ %-3s ] %s %v", log.Time, logLevel, msg, string(jsonAdditional))
	} else {
		return fmt.Sprintf("%-33s [ %-3s ] %s", log.Time, logLevel, msg)
	}
}

func retJSON(log Logs) string {
	var msg string
	switch v := log.Message.(type) {
	case []string:
		msg = strings.Join(v, " ")
	case string:
		msg = v
	case []int:
		intSTRs := make([]string, len(v))
		for i, n := range v {
			intSTRs[i] = fmt.Sprintf("%d", n)
		}

		msg = strings.Join(intSTRs, " ")
	case []interface{}:
		anySTRs := make([]string, len(v))
		for i, elem := range v {
			anySTRs[i] = fmt.Sprintf("%v", elem)
		}

		msg = strings.Join(anySTRs, " ")
	default:
		msg = fmt.Sprintf("%v", v)
	}

	log.Message = msg
	jsonLog, err := json.Marshal(log)
	if err != nil {
		return ""
	}

	return string(jsonLog)
}

func (c *Config) dumpTerminal(log Logs) {
	switch c.terminalFormat {
	case JSONFormat:
		fmt.Println(retJSON(log))
	case ReadableFormat:
		fmt.Println(retReadable(log, true))
	default:
		fmt.Println(retReadable(log, true))
	}
}

func (c *Config) dumpFile(log Logs) {
	c.logRotateByCount()

	var content string
	switch c.fileFormat {
	case JSONFormat:
		content = retJSON(log)
	case ReadableFormat:
		content = retReadable(log, false)
	default:
		content = retReadable(log, false)
	}

	if c.fileDescriptor != nil {
		_, err := c.fileDescriptor.WriteString(content + "\n")
		if err != nil {
			return
		}
	}
}
