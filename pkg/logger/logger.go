package logger

import (
	"fmt"
	"github.com/codeshitmeshit/AsynLogger/pkg/conf"
	"github.com/codeshitmeshit/AsynLogger/pkg/terminal"
)

var (
	mode = 0
)

func init() {
	if conf.IsInit == false {
		conf.Init("")
	}
	mode = conf.Cfg.OutputMode
}

func Info(content interface{}) {
	if mode != 2 {
		terminal.Info(content)
	}
	if mode == 2 || mode == 3 {
		fmt.Println("文件")
	}
}

func Warn(content interface{}) {
	if mode != 2 {
		terminal.Warn(content)
	}
	if mode == 2 || mode == 3 {
		fmt.Println("文件")
	}
}

func Debug(content interface{}) {
	if mode != 2 {
		terminal.Debug(content)
	}
	if mode == 2 || mode == 3 {
		fmt.Println("文件")
	}
}
