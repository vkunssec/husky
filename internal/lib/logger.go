package lib

import (
	"fmt"
	"os"
)

type LogLevel int

const (
	LogLevelSilent LogLevel = iota
	LogLevelError
	LogLevelInfo
	LogLevelDebug
)

var currentLogLevel = LogLevelInfo

func SetLogLevel(level LogLevel) {
	currentLogLevel = level
}

func LogDebug(format string, args ...interface{}) {
	if currentLogLevel >= LogLevelDebug {
		fmt.Printf("DEBUG: "+format+"\n", args...)
	}
}

func LogInfo(format string, args ...interface{}) {
	if currentLogLevel >= LogLevelInfo {
		fmt.Printf("INFO: "+format+"\n", args...)
	}
}

func LogError(format string, args ...interface{}) {
	if currentLogLevel >= LogLevelError {
		fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", args...)
	}
}
