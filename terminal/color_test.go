package terminal

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	str := "hello world"
	fmt.Println(greenBg, str, reset)
	fmt.Println(whiteBg, str, reset)
	fmt.Println(yellowBg, str, reset)
	fmt.Println(redBg, str, reset)
	fmt.Println(blueBg, str, reset)
	fmt.Println(magentaBg, str, reset)
	fmt.Println(cyanBg, str, reset)
	word := "I love you"
	fmt.Println(green, word, reset)
	fmt.Println(white, word, reset)
	fmt.Println(yellow, word, reset)
	fmt.Println(red, word, reset)
	fmt.Println(blue, word, reset)
	fmt.Println(magenta, word, reset)
	fmt.Println(cyan, word, reset)
}
