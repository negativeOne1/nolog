package nolog

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJSONFormatting(t *testing.T) {
	tf := NewBasicJSONFormatter(true, "")
	e := Entry{
		Time:    time.Now(),
		Message: "foo",
		Level:   LevelInfo,
	}
	b, err := tf.Format(e)
	assert.Nil(t, err)
	assert.Equal(t, "{\"level\":\"INFO\",\"msg\":\"foo\"}\n", string(b))

}
