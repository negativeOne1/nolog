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
	level     Threshold
	writer    io.Writer
	formatter Formatter
}

func New() *Logger {
	return &Logger{
		level:     LevelInfo,
		writer:    os.Stdout,
		formatter: new(TextFormatter),
	}
}

func SetWriter(writer io.Writer) {
	std.writer = writer
}

func SetLevel(level Threshold) {
	std.level = level
}

type Formatter interface {
	Format(args ...interface{}) ([]byte, error)
}

type TextFormatter struct{}

func (f *TextFormatter) Format(args ...interface{}) ([]byte, error) {
	ss := ""
	for _, a := range args[1].([]interface{}) {
		ss += a.(string)
	}
	s := fmt.Sprintf("[%s] %s\n", prefixes[args[0].(Threshold)], ss)
	return []byte(s), nil
}

func (l *Logger) writeLog(level Threshold, args ...interface{}) {
	if level >= l.level {
		b, err := l.formatter.Format(level, args)
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
