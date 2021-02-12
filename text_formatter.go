package nolog

import "fmt"

type Formatter interface {
	Format(Entry) ([]byte, error)
}

type TextFormatter struct{}

func (f *TextFormatter) Format(e Entry) ([]byte, error) {
	l := prefixes[e.Level]
	s := fmt.Sprintf("[%s] %s\n", l, e.Message)
	return []byte(s), nil
}
