package main

import (
	"fmt"
	"io"
	"log"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

type Logger struct {
	*log.Logger
	level LogLevel
}

func NewLogger(writer io.Writer, level LogLevel) *Logger {
	return &Logger{
		Logger: log.New(writer, "", log.LstdFlags|log.Lshortfile),
		level:  level,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	if l.level <= DEBUG {
		l.Output(2, fmt.Sprintf("[DEBUG] %s", fmt.Sprint(v...)))
	}
}

func (l *Logger) Info(v ...interface{}) {
	if l.level <= INFO {
		l.Output(2, fmt.Sprintf("[INFO] %s", fmt.Sprint(v...)))
	}
}

func (l *Logger) Warn(v ...interface{}) {
	if l.level <= WARN {
		l.Output(2, fmt.Sprintf("[WARN] %s", fmt.Sprint(v...)))
	}
}

func (l *Logger) Error(v ...interface{}) {
	if l.level <= ERROR {
		l.Output(2, fmt.Sprintf("[ERROR] %s", fmt.Sprint(v...)))
	}
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level <= DEBUG {
		l.Output(2, fmt.Sprintf("[DEBUG] %s", fmt.Sprintf(format, v...)))
	}
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level <= INFO {
		l.Output(2, fmt.Sprintf("[INFO] %s", fmt.Sprintf(format, v...)))
	}
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.level <= WARN {
		l.Output(2, fmt.Sprintf("[WARN] %s", fmt.Sprintf(format, v...)))
	}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level <= ERROR {
		l.Output(2, fmt.Sprintf("[ERROR] %s", fmt.Sprintf(format, v...)))
	}
}
