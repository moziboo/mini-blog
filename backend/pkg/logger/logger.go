package logger

import (
	"fmt"
	"log"
	"os"
)

// Logger represents a simple logging interface
type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
}

// SimpleLogger is a basic implementation of the Logger interface
type SimpleLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

// New creates a new SimpleLogger
func New() Logger {
	return &SimpleLogger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
		fatalLogger: log.New(os.Stderr, "FATAL: ", log.LstdFlags),
	}
}

// Info logs an informational message
func (l *SimpleLogger) Info(msg string, args ...interface{}) {
	if len(args) > 0 && len(args)%2 == 0 {
		l.infoLogger.Printf("%s %v", msg, formatKeyValues(args))
	} else {
		l.infoLogger.Println(msg)
	}
}

// Error logs an error message
func (l *SimpleLogger) Error(msg string, args ...interface{}) {
	if len(args) > 0 && len(args)%2 == 0 {
		l.errorLogger.Printf("%s %v", msg, formatKeyValues(args))
	} else {
		l.errorLogger.Println(msg)
	}
}

// Fatal logs a fatal message and exits
func (l *SimpleLogger) Fatal(msg string, args ...interface{}) {
	if len(args) > 0 && len(args)%2 == 0 {
		l.fatalLogger.Printf("%s %v", msg, formatKeyValues(args))
	} else {
		l.fatalLogger.Println(msg)
	}
	os.Exit(1)
}

// formatKeyValues formats a slice of alternating keys and values as a map-like string
func formatKeyValues(args []interface{}) string {
	if len(args) == 0 {
		return ""
	}

	result := "{"
	for i := 0; i < len(args); i += 2 {
		if i > 0 {
			result += ", "
		}
		result += args[i].(string) + ": " + formatValue(args[i+1])
	}
	result += "}"
	return result
}

// formatValue formats a value as a string
func formatValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case error:
		return v.Error()
	default:
		return fmt.Sprintf("%v", v)
	}
} 