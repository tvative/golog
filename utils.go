package golog

import (
	"fmt"
	"strconv"
	"time"
)

const (
	defaultColor string = "\033[0;0m"  // defaultColor print log messages in default color.
	redColor     string = "\033[1;31m" // redColor print log messages in red color.
	greenColor   string = "\033[1;32m" // greenColor print log messages in green color.
	yellowColor  string = "\033[1;33m" // yellowColor print log messages in yellow color.
	blueColor    string = "\033[1;34m" // blueColor print log messages in blue color.
	magentaColor string = "\033[1;35m" // magentaColor print log messages in magenta color.
	cyanColor    string = "\033[1;36m" // cyanColor print log messages in cyan color.
)

func genID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func genTime() string {
	getTime := time.Now()
	generateLongTime := getTime.Format("2006-01-02 15:04:05")
	generatedTimeMillSeconds := getTime.Nanosecond() / 1e6
	generatedTimeNanoSeconds := getTime.Nanosecond()

	return generateLongTime + ":" +
		strconv.Itoa(generatedTimeMillSeconds) + ":" +
		strconv.Itoa(generatedTimeNanoSeconds)
}

func getColor(level LogLevel) string {
	var color string
	switch level {
	case NormalLog:
		color = defaultColor
	case SuccessLog:
		color = greenColor
	case ErrorLog:
		color = redColor
	case WarningLog:
		color = yellowColor
	case DebugLog:
		color = magentaColor
	case InfoLog:
		color = blueColor
	default:
		color = defaultColor
	}

	return fmt.Sprintf("%s%s%s", color, level, defaultColor)
}
