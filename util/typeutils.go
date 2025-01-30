package util

import "reflect"

func IsString(v interface{}) bool {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.String {
		return true
	}
	return false
}
