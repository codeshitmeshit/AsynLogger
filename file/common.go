package file

import (
	"github.com/codeshitmeshit/AsynLogger/conf"
	"github.com/codeshitmeshit/AsynLogger/constants"
	"github.com/codeshitmeshit/AsynLogger/terminal"
	"github.com/codeshitmeshit/AsynLogger/util"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var (
	fileMode    int
	fileSize    int
	path        string
	currentFile string
	mtx         sync.Mutex
	bytesChan   chan []byte
	fd          *os.File
)

func writeIn() {
	defer func(fd *os.File) {
		_ = fd.Close()
	}(fd)
	for {
		by := <-bytesChan
		if fd != nil {
			_, err := fd.Write(by)
			if err != nil {
				terminal.Debug(err)
			}
		}
	}
}

func init() {
	bytesChan = make(chan []byte, 100)
	fileMode = conf.Cfg.AsynLogger.FileStorageMode
	fileSize = conf.Cfg.AsynLogger.FileSize
	path = conf.Cfg.AsynLogger.FilePath
	if path == "" {
		path = "../.logs"
	}
	if path[len(path)-1] != '/' {
		path += "/"
	}
	err := util.EnsureDirExists(path)
	if err != nil {
		terminal.Debug(err)
	}
	terminal.Info(path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	// 遍历目录下的文件
	for _, file := range files {
		// 检查是否是文件（不包括子目录）
		if !file.IsDir() {
			name := file.Name()
			if checkFileName(name) == name {
				if currentFile == "" {
					currentFile = name
				} else {
					a, _ := strconv.ParseInt(currentFile, 10, 64)
					b, _ := strconv.ParseInt(name, 10, 64)
					if b > a {
						currentFile = name
					}
				}
			}
		}
	}
	if currentFile != "" {
		fd, err = os.OpenFile(path+currentFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}
	go writeIn()
}

func logWrite(log LogMessage) {
	NewFile()
	text := ""
	m := log.L
	if m == 1 {
		text += "[INFO] "
	} else if m == 2 {
		text += "[WARN] "
	} else if m == 3 {
		text += "[ERROR] "
	}
	text += log.T.Format(constants.GO_TIME) + ": "
	text += log.V
	text += "\n"
	bytesChan <- []byte(text)
}

func NewFile() {

	mtx.Lock()
	defer mtx.Unlock()

	now := time.Now().Format(constants.FILE_TIME) + ".log"
	var err error
	if currentFile == "" || now[:8] != currentFile[:8] {
		currentFile = now
		if fd != nil {
			_ = fd.Close()
			fd = nil
		}
		fd, err = os.OpenFile(path+currentFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			terminal.Debug(err)
			return
		}
	}

	if fd == nil {
		fd, err = os.OpenFile(path+currentFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			terminal.Debug(err)
			return
		}
	}
}

type LogMessage struct {
	L int
	T time.Time
	V string
}

func (i LogMessage) Write() {
	logWrite(i)
}

func checkFileName(fileName string) string {
	regexPattern := `^\d{14}\.log$`
	r, _ := regexp.Compile(regexPattern)
	if r.MatchString(fileName) {
		return fileName
	}
	return fileName
}
