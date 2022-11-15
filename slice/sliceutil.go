package slice

import "sort"

func SliceContains[T int | string](slice []T, item T) bool {
	for _, eachItem := range slice {
		if eachItem == item {
			return true
		}
	}
	return false
}

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