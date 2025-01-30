package util

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	a int
	b string
}

func TestUtils(t *testing.T) {
	fmt.Println(IsString("123"))
	fmt.Println(IsString(1))
	fmt.Println(IsString(TestStruct{1, "!23"}))
}
