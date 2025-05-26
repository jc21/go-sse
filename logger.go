package sse

import (
	"fmt"
	"io"
	"log"
)

// Logger Interface with all necessary functions to log.
type LogPrinter interface {
	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Error(errorClass string, err error)
}

// DefaultLogger ...
type DefaultLogger struct {
	log *log.Logger
}

// NewDefaultLogger ...
func NewDefaultLogger() DefaultLogger {
	return DefaultLogger{
		log: log.New(io.Discard, "", log.LstdFlags),
	}
}

// Debug ...
func (l *DefaultLogger) Debug(format string, args ...any) {
	l.log.Printf(fmt.Sprintf("Debug: %s", format), args...)
}

// Info ...
func (l *DefaultLogger) Info(format string, args ...any) {
	l.log.Printf(fmt.Sprintf("Info: %s", format), args...)
}

// Warn ...
func (l *DefaultLogger) Warn(format string, args ...any) {
	l.log.Printf(fmt.Sprintf("Warn: %s", format), args...)
}

// Error ...
func (l *DefaultLogger) Error(errorClass string, err error) {
	l.log.Printf("Error: %s - %s", errorClass, err.Error())
}
