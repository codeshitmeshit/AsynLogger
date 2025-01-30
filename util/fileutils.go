package util

import (
	"fmt"
	"os"
)

func EnsureDirExists(path string) error {
	// 检查目录是否存在
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		// 目录不存在，创建目录
		err = os.MkdirAll(path, 0755) // 0755 是设置权限：所有者可读写，其他用户可读
		if err != nil {
			return fmt.Errorf("创建目录失败: %v", err)
		}
		fmt.Println("目录创建成功:", path)
	} else if err != nil {
		// 发生其他错误
		return fmt.Errorf("检查目录失败: %v", err)
	} else if !info.IsDir() {
		// 如果路径存在，但它不是目录
		return fmt.Errorf("指定路径已经存在，但不是目录: %s", path)
	}

	return nil
}
