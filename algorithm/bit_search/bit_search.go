package bit_search

import "cmp"

func BinarySearch[T cmp.Ordered](array []T, target T) int {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := left + (right-left)>>1
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
