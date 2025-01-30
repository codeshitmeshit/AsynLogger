package terminal

import (
	"fmt"
	"github.com/codeshitmeshit/AsynLogger/constants"
	"time"
)

func Debug(content interface{}) {
	text := fmt.Sprintf("%v%v%v", redBg, content, reset)
	text = red + time.Now().Format(constants.GO_TIME) + ": " + reset + text
	fmt.Println(text)
}

func Warn(content interface{}) {
	timeStr := time.Now().Format(constants.GO_TIME)
	text := fmt.Sprintf("%v%v: %v%v", yellow, timeStr, content, reset)
	fmt.Println(text)
}

func Info(content interface{}) {
	timeStr := time.Now().Format(constants.GO_TIME)
	text := fmt.Sprintf("%v%v: %v%v", blue, timeStr, content, reset)
	fmt.Println(text)
}
