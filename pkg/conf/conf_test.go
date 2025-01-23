package conf

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	Init()
	str := Cfg.ossBucketName
	fmt.Println(str)
}
