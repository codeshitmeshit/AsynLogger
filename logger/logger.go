package logger

import (
	"encoding/json"
	"github.com/codeshitmeshit/AsynLogger/conf"
	"github.com/codeshitmeshit/AsynLogger/constants"
	"github.com/codeshitmeshit/AsynLogger/terminal"
	"github.com/codeshitmeshit/AsynLogger/util"
	"github.com/codeshitmeshit/AsynLogger/worker"
	"strconv"
	"time"
)

var (
	mode = 0
)

func init() {
	if conf.IsInit == false {
		conf.Init("../conf/conf.toml")
	}
	mode = conf.Cfg.OutputMode
	if mode == 2 {
		Info("OutputMode is: " + strconv.Itoa(mode))
	} else {
		worker.Log(constants.INFO, time.Now(), "OutputMode is: "+strconv.Itoa(mode))
	}
}

func Info(content interface{}) {
	c := getString(content)
	if mode != 1 {
		terminal.Info(c)
	}
	if mode == 1 || mode == 2 {
		worker.Log(constants.INFO, time.Now(), getString(c))
	}
}

func Warn(content interface{}) {
	c := getString(content)
	if mode != 1 {
		terminal.Warn(c)
	}
	if mode == 1 || mode == 2 {
		worker.Log(constants.WARN, time.Now(), getString(c))
	}
}

func Debug(content interface{}) {
	c := getString(content)
	if mode != 1 {
		terminal.Debug(c)
	}
	if mode == 1 || mode == 2 {
		worker.Log(constants.DEBUG, time.Now(), c)
	}
}

func MInfo(content interface{}) {
	c := mGetString(content)
	if mode != 1 {
		terminal.Info(c)
	}
	if mode == 1 || mode == 2 {
		worker.Log(constants.INFO, time.Now(), getString(c))
	}
}

func MWarn(content interface{}) {
	c := mGetString(content)
	if mode != 1 {
		terminal.Warn(c)
	}
	if mode == 1 || mode == 2 {
		worker.Log(constants.WARN, time.Now(), getString(c))
	}
}

func MDebug(content interface{}) {
	c := mGetString(content)
	if mode != 1 {
		terminal.Debug(c)
	}
	if mode == 1 || mode == 2 {
		worker.Log(constants.DEBUG, time.Now(), getString(c))
	}
}

func getString(v interface{}) string {
	if util.IsString(v) {
		return v.(string)
	}
	jsonStr, err := json.Marshal(v)
	if err != nil {
		terminal.Debug("Json Str Parse Error")
	}
	return string(jsonStr)
}

func mGetString(v interface{}) string {
	c, l := util.GetDetail()
	c = "[" + c
	c += ":" + strconv.Itoa(l)
	c += "] "
	c += getString(v)
	return c
}
