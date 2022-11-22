package sliceutil

import (
	"fmt"
	"testing"
)

func TestSliceContains(t *testing.T) {
	s := []string{"aaa", "bbb", "ccc", "aaa"}
	fmt.Println(SliceContains(s, "aaa"))
}

func TestSliceRemoveDuplicates(t *testing.T) {
	s := []string{"aaa", "bbb", "ccc", "aaa"}
	fmt.Println(SliceRemoveDuplicates(s))
}
