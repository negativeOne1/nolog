package nolog

import (
	"fmt"
	"time"
)

type Entry struct {
	Time    time.Time
	Message string
	Level   Level
}

func NewEntry(l Level, args ...interface{}) Entry {
	return Entry{
		Time:    time.Now(),
		Message: fmt.Sprint(args...),
		Level:   l,
	}
}
