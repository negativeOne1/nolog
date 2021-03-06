package nolog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTextFormatting(t *testing.T) {
	tf := NewBasicTextFormatter(true, "")
	e := Entry{
		Time:    time.Now(),
		Message: "foo",
		Level:   LevelInfo,
	}
	b, err := tf.Format(e)
	assert.Nil(t, err)
	assert.Equal(t, "[INFO] foo\n", string(b))
}
