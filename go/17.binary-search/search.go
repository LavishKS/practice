package main

func binary_search_recursive(arr []int, left int, right int, element int) int {
	if right < left {
		return -1
	}

	mid := (left + right) / 2
	if arr[mid] == element {
		return mid
	} else if arr[mid] < element {
		left = mid + 1
	} else {
		right = mid - 1
	}

	return binary_search_recursive(arr, left, right, element)
}

func binary_search_iterative(arr []int, element int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == element {
			return mid
		} else if arr[mid] < element {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
