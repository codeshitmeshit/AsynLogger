package worker

import (
	"github.com/codeshitmeshit/AsynLogger/constants"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	Log(constants.INFO, time.Now(), "11")
	Log(constants.WARN, time.Now(), "22")
	Log(constants.DEBUG, time.Now(), "33")
}
