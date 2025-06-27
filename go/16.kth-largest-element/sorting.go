package main

import "fmt"

func sorting_main() {
	arrs := [][]int{
		{6, 5, 3, 13, 1, 8, 7, 2, 4},
		{99, 44, 6, 2, 1, 5, 63, 87, 287, 4, 0},
		{3, 7, 8, 5, 2, 1, 9, 5, 4},
	}
	for _, arr := range arrs {
		fmt.Println("Before sorting Input:", arr)
		// bubbleSort(arr)
		// selectionSort(arr)
		// insertionSort(arr)
		// arr = mergeSort(arr, 0, len(arr)-1)
		quickSort(arr, 0, len(arr)-1)
		fmt.Println("After sort:", arr)
	}
}

func hoareQuickSelectAlgortihm(arr []int, k int, start int, end int) int {
	/**
	 * Finds the "k"th smallest element in an unordered array "arr"
	 */

	if end <= start {
		return -1
	}
	pivot := partition(arr, start, end)

	if pivot == k {
		return arr[pivot]
	}

	if pivot < k {
		start = pivot + 1
	} else {
		end = pivot - 1
	}
	return hoareQuickSelectAlgortihm(arr, k, start, end)
}

func quickSort(arr []int, start int, end int) {
	if end <= start {
		return
	}
	pivot, pVal := end, arr[end]
	for i := start; i < pivot; {
		if arr[i] > pVal {
			arr[pivot] = arr[i]
			pivot--
			arr[i] = arr[pivot]
		} else {
			i++
		}
	}
	arr[pivot] = pVal
	quickSort(arr, start, pivot-1)
	quickSort(arr, pivot+1, end)
}

func quickSortV2(arr []int, start int, end int) {
	if end <= start {
		return
	}

	pivot := partition(arr, start, end)
	quickSortV2(arr, start, pivot-1)
	quickSortV2(arr, pivot+1, end)
}

func partition(arr []int, start int, end int) int {
	pVal := arr[end]
	partitionIndex := start
	for j := start; j < end; j++ {
		if arr[j] < pVal {
			arr[partitionIndex], arr[j] = arr[j], arr[partitionIndex]
			partitionIndex++
		}
	}
	arr[end], arr[partitionIndex] = arr[partitionIndex], arr[end]

	return partitionIndex
}

func bubbleSort(arr []int) {
	for i := range arr {
		for j, k := 0, 1; k < len(arr)-i; j, k = j+1, k+1 {
			if arr[k] < arr[j] {
				temp := arr[j]
				arr[j] = arr[k]
				arr[k] = temp
			}
		}
	}
}

func selectionSort(arr []int) {
	for i, v := range arr {
		smallest := i
		for j := i + 1; j < len(arr); j++ {
			if arr[smallest] > arr[j] {
				smallest = j
			}
		}
		if smallest != i {
			arr[i] = arr[smallest]
			arr[smallest] = v
		}
	}
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		j, k, v := i-1, i, arr[i]
		for ; j >= 0 && arr[j] > v; j, k = j-1, k-1 {
			arr[k] = arr[j]
		}
		arr[k] = v
	}
}

func mergeSort(arr []int, start int, end int) []int {
	if start == end {
		return []int{arr[start]}
	}

	mid := (start + end) / 2
	return merge(mergeSort(arr, start, mid), mergeSort(arr, mid+1, end))
}

func merge(first []int, second []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(first) && j < len(second) {
		smallest := first[i]
		if second[j] < first[i] {
			smallest = second[j]
			j++
		} else {
			i++
		}
		result = append(result, smallest)
	}
	result = append(result, first[i:]...)
	result = append(result, second[j:]...)
	return result
}
