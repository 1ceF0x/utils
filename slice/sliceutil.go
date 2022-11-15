package slice

import (
	"sort"
	"strings"
)

// 切片是否包含元素
func SliceContains[T int | string](slice []T, item T) bool {
	for _, eachItem := range slice {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 切片去重
func SliceRemoveDuplicates[T int | string](slice []T) []T {
	if len(slice) < 2 {
		return slice
	}
	sort.SliceStable(slice, func(i, j int) bool { return slice[i] < slice[j] })
	uniqPointer := 0
	for i := 1; i < len(slice); i++ {
		if slice[uniqPointer] != slice[i] {
			uniqPointer++
			slice[uniqPointer] = slice[i]
		}
	}
	return slice[:uniqPointer+1]
}

func IsBlank(value string) bool {
	return strings.TrimSpace(value) == ""
}
