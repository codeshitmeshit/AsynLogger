package util

import "runtime"

func GetDetail() (string, int) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "Unknown", 0
	}
	return file, line
}
