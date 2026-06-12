package logger

import (
	"log"
)

type LoggerLevel int

const (
	DebugLevel LoggerLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func (l LoggerLevel) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	default:
		return "unknown"
	}
}

func ParseLoggerLevel(level string) LoggerLevel {
	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "error":
		return ErrorLevel
	default:
		log.Printf("Invalid logger level '%s', defaulting to 'info'", level)
		return InfoLevel // Default to info if invalid level is provided
	}
}

func NewLogger(level string) *Logger {
	return New(ParseLoggerLevel(level))
}

type Logger struct {
	Level LoggerLevel
	log   *log.Logger
}

func New(level LoggerLevel) *Logger {
	return &Logger{Level: LoggerLevel(level), log: log.Default()}
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.Level <= DebugLevel {
		l.log.Printf("[DEBUG] "+msg, args...)
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	if l.Level <= InfoLevel {
		l.log.Printf("[INFO] "+msg, args...)
	}
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	if l.Level <= WarnLevel {
		l.log.Printf("[WARN] "+msg, args...)
	}
}

func (l *Logger) Error(msg string, args ...interface{}) {
	if l.Level <= ErrorLevel {
		l.log.Printf("[ERROR] "+msg, args...)
	}
}
