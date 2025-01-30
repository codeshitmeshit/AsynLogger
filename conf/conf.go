package conf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	IsInit bool = false
)

type AliyunOss struct {
	OssEndpoint        string `toml:"oss_endpoint"`
	OssAccessKeyId     string `toml:"oss_accessKeyId"`
	OssAccessKeySecret string `toml:"oss_accessKeySecret"`
	OssBucketName      string `toml:"oss_bucketName"`
}

type AsynLogger struct {
	OutputMode      int    `toml:"output_mode"`
	LogMode         int    `toml:"log_mode"`
	OpenOss         int    `toml:"open_oss"`
	OpenCos         int    `toml:"open_cos"`
	FileStorageMode int    `toml:"file_storage_mode"`
	FileSize        int    `toml:"file_size"`
	FilePath        string `toml:"file_path"`
}

type Config struct {
	AsynLogger `toml:"asyn_logger"`
	Aliyun     AliyunOss `toml:"aliyun_oss"`
}

var (
	Cfg     = new(Config)
	cfgPath = ""
)

func Init(path string) {
	if path != "" {
		cfgPath = path
	} else {
		cfgPath = "conf.toml"
	}
	_, err := os.Stat(cfgPath)

	if err != nil {
		//log.Println(cfgPath, "config file not found, use default config")
	} else {
		_, err = toml.DecodeFile(cfgPath, Cfg)
		if err != nil {
			log.Fatal("decode toml failed: ", err)
		}
	}
	IsInit = true
}
func (Logger *AsynLogger) valueInit() {

}
