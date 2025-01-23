package conf

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
)

type AsynLogger struct {
	openAsyn int  `toml:"open_asyn"`
	logMode  int  `toml:"log_mode"`
	openOss  bool `toml:"open_oss"`
	openCos  bool `toml:"open_cos"`

	ossEndpoint        string `toml:"oss_endpoint"`
	ossAccessKeyId     string `toml:"oss_accessKeyId"`
	ossAccessKeySecret string `toml:"oss_accessKeySecret"`
	ossBucketName      string `toml:"oss_bucketName"`
}

var (
	Cfg     = new(AsynLogger)
	cfgPath = flag.String("config", "conf.toml", "config file path")
)

func Init() {
	_, err := os.Stat(*cfgPath)
	fmt.Println(*cfgPath)
	// 打开文件
	file, err := os.Open(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 读取文件内容
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// 打印文件内容
	fmt.Println(string(content))
	if err != nil {
		log.Println(*cfgPath, "config file not found, use default config")
	} else {
		_, err = toml.DecodeFile(*cfgPath, Cfg)
		if err != nil {
			log.Fatal("decode toml failed: ", err)
		}
	}
}

func (Logger *AsynLogger) valueInit() {

}
