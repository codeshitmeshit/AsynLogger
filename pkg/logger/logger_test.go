package logger

import "testing"

func TestInfo(t *testing.T) {
	Info("hello")
}

func TestWarn(t *testing.T) {
	Warn("hello")
}

func TestDebug(t *testing.T) {
	Debug("hello")
}
