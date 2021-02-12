package nolog

import (
	"bytes"
	"encoding/json"
)

type JSONFormatter struct {
	timestampFormat string
	ignoreTime      bool
}

func NewBasicJSONFormatter(ignoreTime bool, timestampFormat string) *JSONFormatter {
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	return &JSONFormatter{
		timestampFormat,
		ignoreTime,
	}
}

func (f *JSONFormatter) Format(e Entry) ([]byte, error) {
	d := map[string]interface{}{
		FieldKeyLevel: e.Level.String(),
		FieldKeyMsg:   e.Message,
	}

	if !f.ignoreTime {
		d[FieldKeyTime] = e.Time
	}

	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(d)
	return buf.Bytes(), err
}
