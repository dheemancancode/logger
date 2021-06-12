package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Log levels for controlling the logging output.
type LogLevel int

const (
	LEVEL_TRACE LogLevel = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_PANIC
)

var logLevelNames = []string{"trace", "debug", "info", "warn", "error", "panic"}

// LogLevel : controls the global log Level used by the logger.
var level = LEVEL_INFO

// SetLogLevel : sets the global log Level used by the simple
// logger.
func SetLogLevel(l LogLevel) {
	fmt.Printf("LOGGER: Setting the log level to %s\n", logLevelNames[l])
	if l <= LEVEL_DEBUG {
		logger.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	} else {
		logger.SetFlags(0)
	}
	level = l
}

// SetLogLevelByName : attempts to sets the global log level
// using a string name.
func SetLogLevelByName(levelName string) error {
	for i, name := range logLevelNames {
		if strings.ToLower(levelName) == name {
			SetLogLevel(LogLevel(i))
			return nil
		}
	}
	return fmt.Errorf("unknown log level name %s: options are %v", levelName, logLevelNames)
}

// GetLogLevel : Gets the global log level
func GetLogLevel() LogLevel {
	return level
}

func GetLogLevelNames() []string {
	return logLevelNames
}

// logger references the used application logger.
var logger = log.New(os.Stdout, "", 0)

// Debug : logs a message at metric Level.
func Trace(format string, v ...interface{}) {
	if level <= LEVEL_TRACE {
		logger.Output(2, "[Trace] "+fmt.Sprintf(format, v...))
	}
}

// Debug : logs a message at debug Level.
func Debug(format string, v ...interface{}) {
	if level <= LEVEL_DEBUG {
		logger.Output(2, "[Debug] "+fmt.Sprintf(format, v...))
	}
}

// Info : logs a message at info Level.
func Info(format string, v ...interface{}) {
	if level <= LEVEL_INFO {
		logger.Output(2, "[Info] "+fmt.Sprintf(format, v...))
	}
}

// Warn : logs a message at warning Level.
func Warn(format string, v ...interface{}) {
	if level <= LEVEL_WARN {
		logger.Output(2, "[Warning] "+fmt.Sprintf(format, v...))
	}
}

// Error : logs a message at error Level.
func Error(format string, v ...interface{}) {
	if level <= LEVEL_ERROR {
		logger.Output(2, "[Error] "+fmt.Sprintf(format, v...))
	}
}

// Panic : logs a message and panics.
func Panic(format string, v ...interface{}) {
	if level <= LEVEL_PANIC {
		logger.Output(2, "[Panic] "+fmt.Sprintf(format, v...))
		panic(fmt.Sprintf(format, v...))
	}
}
