package nolog

type TextFormatter struct {
	timestampFormat string
	ignoreTime      bool
}

func NewBasicTextFormatter(ignoreTime bool, timestampFormat string) *TextFormatter {
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	return &TextFormatter{
		timestampFormat,
		ignoreTime,
	}
}

func (f *TextFormatter) Format(e Entry) ([]byte, error) {
	s := "[" + e.Level.String() + "]"
	if !f.ignoreTime {
		s += " " + e.Time.Format(f.timestampFormat)
	}
	s += " " + e.Message + "\n"
	return []byte(s), nil
}
