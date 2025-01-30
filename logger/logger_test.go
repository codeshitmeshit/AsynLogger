package logger

import (
	"testing"
	"time"
)

func TestInfo(t *testing.T) {
	MInfo("hello")
}

func TestWarn(t *testing.T) {
	Warn("hello")
	time.Sleep(2 * time.Second)
}

func TestDebug(t *testing.T) {
	MDebug("hello")
	time.Sleep(2 * time.Second)
}
