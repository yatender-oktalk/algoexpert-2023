package binarysearch

func BinarySearch(arr []int, start_index int, end_index int, target int) int {
	// first check the middle element
	if start_index == end_index {
		if arr[start_index] == target {
			return start_index
		}
		return -1
	}
	mid := (start_index + end_index) / 2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] >= target {
		return BinarySearch(arr, start_index, mid, target)
	} else {
		if arr[mid] < target {
			return BinarySearch(arr, mid+1, end_index, target)
		}
	}
	return -1
}
