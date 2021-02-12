package nolog

import (
	"io"
	"os"
)

var std = New()

type Level int

func (l Level) String() string {
	return prefixes[l]
}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var prefixes map[Level]string = map[Level]string{
	LevelDebug: "DEBUG",
	LevelInfo:  "INFO",
	LevelWarn:  "WARN",
	LevelError: "ERROR",
	LevelFatal: "FATAL",
}

type Logger struct {
	level     Level
	writer    io.Writer
	formatter Formatter
}

func New() *Logger {
	return &Logger{
		level:     LevelInfo,
		writer:    os.Stdout,
		formatter: NewBasicTextFormatter(false, ""),
	}
}

func SetWriter(writer io.Writer) {
	std.writer = writer
}

func SetLevel(level Level) {
	std.level = level
}

func SetFormatter(formatter Formatter) {
	std.formatter = formatter
}

func (l *Logger) writeLog(level Level, args ...interface{}) {
	if level >= l.level {
		b, err := l.formatter.Format(NewEntry(level, args...))
		if err != nil {
			panic(err)
		}
		l.writer.Write(b)
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
