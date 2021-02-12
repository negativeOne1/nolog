package nolog

import (
	"bytes"
	"strconv"
	"testing"
)

func assertStringEqual(t *testing.T, expect, got string) {
	if expect != got {
		t.Errorf("Expected: %s GOT: %s", strconv.Quote(expect), strconv.Quote(got))
	}
}

func TestBasic(t *testing.T) {
	buf := new(bytes.Buffer)
	SetWriter(buf)
	Info("info")

	expected := "[INFO] info\n"
	assertStringEqual(t, expected, buf.String())
}

func TestLevelFilter(t *testing.T) {
	buf := new(bytes.Buffer)
	SetWriter(buf)
	SetLevel(LevelError)
	Info("info")
	Error("error")

	expected := "[ERROR] error\n"
	assertStringEqual(t, expected, buf.String())
}
