package randutil

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	length := 100
	choices := "asdh443242!@#~!#@$%%^_+_-(*&^"
	fmt.Println(RandFromChoices(choices, length))
	fmt.Println(RandString(length))
	fmt.Println(RandStringLower(length))
	fmt.Println(RandStringUpper(length))
	fmt.Println(RandNumber(length))
	fmt.Println(RandNumberString(length))
	fmt.Println(RandNumberStringLower(length))
	fmt.Println(RandNumberStringUpper(length))
	fmt.Println(RandUserAgent())

}
