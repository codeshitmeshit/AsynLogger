package conf

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type AliyunOss struct {
	OssEndpoint        string `toml:"oss_endpoint"`
	OssAccessKeyId     string `toml:"oss_accessKeyId"`
	OssAccessKeySecret string `toml:"oss_accessKeySecret"`
	OssBucketName      string `toml:"oss_bucketName"`
}

type AsynLogger struct {
	OpenAsyn int `toml:"open_asyn"`
	LogMode  int `toml:"log_mode"`
	OpenOss  int `toml:"open_oss"`
	OpenCos  int `toml:"open_cos"`

	Aliyun AliyunOss `toml:"aliyun"`
}

type Config struct {
	AsynLogger `toml:"asyn_logger"`
}

var (
	Cfg     = new(Config)
	cfgPath = ""
)

func Init(path string) {
	if path != "" {
		cfgPath = path
	}
	_, err := os.Stat(cfgPath)
	fmt.Println(cfgPath)

	if err != nil {
		log.Println(cfgPath, "config file not found, use default config")
	} else {
		_, err = toml.DecodeFile(cfgPath, Cfg)
		if err != nil {
			log.Fatal("decode toml failed: ", err)
		}
	}
}

func (Logger *AsynLogger) valueInit() {

}
