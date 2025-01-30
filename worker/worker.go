package worker

import (
	"github.com/codeshitmeshit/AsynLogger/file"
	"time"
)

var (
	logChannel = make(chan file.LogMessage, 100)
)

func init() {
	go func() {
		for {
			i := <-logChannel
			i.Write()
		}
	}()
}

func Log(l int, t time.Time, v string) {
	logChannel <- file.LogMessage{
		L: l,
		T: t,
		V: v,
	}
}
