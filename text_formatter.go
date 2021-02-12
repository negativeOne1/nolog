package nolog

import "fmt"

type Formatter interface {
	Format(Entry) ([]byte, error)
}

type TextFormatter struct {
	formatString string
}

func NewBasicTextFormatter() *TextFormatter {
	return &TextFormatter{
		formatString: "[%s] %s\n",
	}
}

func (f *TextFormatter) Format(e Entry) ([]byte, error) {
	l := prefixes[e.Level]
	s := fmt.Sprintf(f.formatString, l, e.Message)
	return []byte(s), nil
}
