package conf

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	Init("conf.toml")
	fmt.Println(Cfg.Aliyun.OssBucketName)
}
