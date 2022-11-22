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

	test := []string{"aaa", "bbb", "ccc", "ddd", "eee"}
	test1 := []int{1, 2, 3, 4, 5}
	fmt.Println(RandFromSlice(test))
	fmt.Println(RandFromSlice(test1))

}
