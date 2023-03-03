package sliceutil

import (
	"sort"
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

// 切片移除指定元素
func SliceRemoveElement[T int | string](slice []T, elem T) []T {
	i := 0
	for _, v := range slice {
		if v != elem {
			slice[i] = v
			i++
		}
	}
	return slice[:i]
}
