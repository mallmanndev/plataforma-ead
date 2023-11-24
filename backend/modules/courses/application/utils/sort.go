package utils

import "sort"

func SortSlice[T comparable](resolutions []T, less func(int, int) bool) []T {
	sortedResolutions := make([]T, len(resolutions))
	copy(sortedResolutions, resolutions)
	sort.Slice(sortedResolutions, less)
	return sortedResolutions
}
