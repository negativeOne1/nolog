package nolog

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	buf := new(bytes.Buffer)
	SetWriter(buf)
	SetFormatter(NewBasicTextFormatter(true, ""))

	Info("info")

	expected := "[INFO] info\n"
	assert.Equal(t, expected, buf.String())
}

func TestLevelFilter(t *testing.T) {
	buf := new(bytes.Buffer)
	SetWriter(buf)
	SetFormatter(NewBasicTextFormatter(true, ""))

	SetLevel(LevelError)

	Info("info")
	Error("error")

	expected := "[ERRO] error\n"
	assert.Equal(t, expected, buf.String())
}
