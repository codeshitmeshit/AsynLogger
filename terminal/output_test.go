package terminal

import (
	"testing"
)

func TestDebug(t *testing.T) {
	Debug("hello")
}

func TestWarn(t *testing.T) {
	Warn("hello")
}

func TestInfo(t *testing.T) {
	Info("hello")
}
