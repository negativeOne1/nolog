package nolog

import (
	"bytes"
	"testing"
)

func TestBasic(t *testing.T) {
	buf := new(bytes.Buffer)
	SetWriter(buf)
	Info("info")

	expected := "[INFO] [info]\n"
	if buf.String() != expected {
		t.Errorf("Expected: %s GOT: %s", expected, buf.String())
	}
}

func TestLevelFilter(t *testing.T) {
	buf := new(bytes.Buffer)
	SetWriter(buf)
	SetLevel(LevelError)
	Info("info")
	Error("error")

	expected := "[ERROR] [error]\n"
	if buf.String() != expected {
		t.Errorf("Expected: %s GOT: %s", expected, buf.String())
	}
}
