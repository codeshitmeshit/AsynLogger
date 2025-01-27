package file

import (
	"fmt"
	"github.com/codeshitmeshit/AsynLogger/conf"
	"github.com/codeshitmeshit/AsynLogger/constants"
	"github.com/codeshitmeshit/AsynLogger/terminal"
	"testing"
	"time"
)

func TestCommon(t *testing.T) {
	fmt.Println(time.Now().Format(constants.FILE_TIME))
}

func TestWrite(t *testing.T) {
	terminal.Info("init")
	conf.Init(constants.CONF_PATH)

	meg := LogMessage{
		L: constants.INFO,
		T: time.Now(),
		V: "123",
	}
	meg.Write()
}
