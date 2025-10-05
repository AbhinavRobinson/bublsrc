package main

import (
	"fmt"
	"io"
	"log"
)

// LogLevel represents the logging level
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// LoggerService handles all logging operations
type LoggerService struct {
	*log.Logger
	level LogLevel
}

// NewLoggerService creates a new logger service
func NewLoggerService(writer io.Writer, level LogLevel) *LoggerService {
	return &LoggerService{
		Logger: log.New(writer, "", log.LstdFlags|log.Lshortfile),
		level:  level,
	}
}

// Debug logs a debug message
func (l *LoggerService) Debug(v ...interface{}) {
	if l.level <= DEBUG {
		l.Output(2, fmt.Sprintf("[DEBUG] %s", fmt.Sprint(v...)))
	}
}

// Info logs an info message
func (l *LoggerService) Info(v ...interface{}) {
	if l.level <= INFO {
		l.Output(2, fmt.Sprintf("[INFO] %s", fmt.Sprint(v...)))
	}
}

// Warn logs a warning message
func (l *LoggerService) Warn(v ...interface{}) {
	if l.level <= WARN {
		l.Output(2, fmt.Sprintf("[WARN] %s", fmt.Sprint(v...)))
	}
}

// Error logs an error message
func (l *LoggerService) Error(v ...interface{}) {
	if l.level <= ERROR {
		l.Output(2, fmt.Sprintf("[ERROR] %s", fmt.Sprint(v...)))
	}
}

// Debugf logs a formatted debug message
func (l *LoggerService) Debugf(format string, v ...interface{}) {
	if l.level <= DEBUG {
		l.Output(2, fmt.Sprintf("[DEBUG] %s", fmt.Sprintf(format, v...)))
	}
}

// Infof logs a formatted info message
func (l *LoggerService) Infof(format string, v ...interface{}) {
	if l.level <= INFO {
		l.Output(2, fmt.Sprintf("[INFO] %s", fmt.Sprintf(format, v...)))
	}
}

// Warnf logs a formatted warning message
func (l *LoggerService) Warnf(format string, v ...interface{}) {
	if l.level <= WARN {
		l.Output(2, fmt.Sprintf("[WARN] %s", fmt.Sprintf(format, v...)))
	}
}

// Errorf logs a formatted error message
func (l *LoggerService) Errorf(format string, v ...interface{}) {
	if l.level <= ERROR {
		l.Output(2, fmt.Sprintf("[ERROR] %s", fmt.Sprintf(format, v...)))
	}
}
