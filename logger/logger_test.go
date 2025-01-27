package logger

import "testing"

func TestInfo(t *testing.T) {
	MInfo("hello")
}

func TestWarn(t *testing.T) {
	Warn("hello")
}

func TestDebug(t *testing.T) {
	Debug("hello")
}
