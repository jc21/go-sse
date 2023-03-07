package sse

import (
	"fmt"
	"io"
	"log"
)

// Logger Interface with all necessary functions to log.
type LogPrinter interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(errorClass string, err error)
}

type DefaultLogger struct {
	log *log.Logger
}

func NewDefaultLogger() DefaultLogger {
	return DefaultLogger{
		log: log.New(io.Discard, "", log.LstdFlags),
	}
}

func (l *DefaultLogger) Debug(format string, args ...interface{}) {
	l.log.Printf(fmt.Sprintf("Debug: %s", format), args...)
}

func (l *DefaultLogger) Info(format string, args ...interface{}) {
	l.log.Printf(fmt.Sprintf("Info: %s", format), args...)
}

func (l *DefaultLogger) Warn(format string, args ...interface{}) {
	l.log.Printf(fmt.Sprintf("Warn: %s", format), args...)
}

func (l *DefaultLogger) Error(errorClass string, err error) {
	l.log.Printf(fmt.Sprintf("Error: %s - %s", errorClass, err.Error()))
}
