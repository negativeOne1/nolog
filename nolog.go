package nolog

import (
	"fmt"
	"io"
	"os"
)

var std = New()

type Threshold int

const (
	LevelDebug Threshold = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var prefixes map[Threshold]string = map[Threshold]string{
	LevelDebug: "DEBUG",
	LevelInfo:  "INFO",
	LevelWarn:  "WARN",
	LevelError: "ERROR",
	LevelFatal: "FATAL",
}

type Logger struct {
	level  Threshold
	writer io.Writer
}

func New() *Logger {
	return &Logger{
		level:  LevelInfo,
		writer: os.Stdout,
	}
}

func SetWriter(writer io.Writer) {
	std.writer = writer
}

func SetLevel(level Threshold) {
	std.level = level
}

func (l *Logger) writeLog(level Threshold, args ...interface{}) {
	if level >= l.level {
		s := fmt.Sprintf("[%s] %s\n", prefixes[level], args)
		l.writer.Write([]byte(s))
	}
}

func Debug(args ...interface{}) {
	std.writeLog(LevelDebug, args...)
}

func Info(args ...interface{}) {
	std.writeLog(LevelInfo, args...)
}

func Warn(args ...interface{}) {
	std.writeLog(LevelWarn, args...)
}

func Error(args ...interface{}) {
	std.writeLog(LevelError, args...)
}

func Fatal(args ...interface{}) {
	std.writeLog(LevelError, args...)
	os.Exit(1)
}
