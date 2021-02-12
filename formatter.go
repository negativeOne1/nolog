package nolog

import "time"

const (
	defaultTimestampFormat = time.RFC3339
	FieldKeyMsg            = "msg"
	FieldKeyLevel          = "level"
	FieldKeyTime           = "time"
)

type Formatter interface {
	Format(Entry) ([]byte, error)
}
