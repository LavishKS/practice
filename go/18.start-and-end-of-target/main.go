package main

import "fmt"

type Range struct {
	start int
	end   int
}

func (r *Range) UpdateBounds(t int) {
	if r.start > t {
		r.start = t
	} else if r.end < t {
		r.end = t
	}
}

func (r Range) String() string {
	return fmt.Sprintf("Range{start: %d, end: %d}", r.start, r.end)
}

func (r Range) Equal(o *Range) bool {
	if o == nil {
		return false
	}

	return r.start == o.start && r.end == o.end
}

type TestCase struct {
	Input    []int
	Target   int
	Expected *Range
}

func (t TestCase) String() string {
	return fmt.Sprintf("\tInput    : %v\n\tK        : %d\n\tExpected : %s\n", t.Input, t.Target, t.Expected)
}

func (t TestCase) validateResult(result *Range) bool {
	if result == nil {
		return t.Expected == nil
	}
	return result.Equal(t.Expected)
}

func main() {
	testCases := []TestCase{
		{[]int{1, 3, 3, 5, 5, 5, 8, 9}, 5, &Range{3, 5}},
		{[]int{1, 3, 3, 5, 5, 5, 8, 9}, 2, nil},
		{[]int{1, 2, 3, 4, 5, 6}, 4, &Range{3, 3}},
		{[]int{1, 2, 3, 4, 5, 6}, 9, nil},
		{[]int{}, 3, nil},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := search_element_range_optimized(testCase.Input, testCase.Target, 0, len(testCase.Input)-1, nil)
		fmt.Println("\tActual:", result)

		if testCase.validateResult(result) {
			passed++
			fmt.Println("+ Testcase Passed!")
		} else {
			fmt.Println("- Testcase Failed!")
		}
		fmt.Println()
	}

	fmt.Println(passed, "out of", len(testCases), "testcases passed!")
}

func search_element_range_o_n(arr []int, target int) *Range {
	// return binary_search_recursive(arr, 0, len(arr)-1, element)
	start := binary_search_iterative(arr, target)
	if start < 0 {
		return nil
	}
	end := start
	for start > 0 && arr[start-1] == target {
		start--
	}
	for end < len(arr)-1 && arr[end+1] == target {
		end++
	}
	return &Range{start, end}
}

func search_element_range_optimized(arr []int, target int, left int, right int, result *Range) *Range {
	if right < left {
		return result
	}

	mid := (left + right) / 2
	if arr[mid] == target {
		if result == nil {
			result = &Range{mid, mid}
			search_element_range_optimized(arr, target, left, mid-1, result)
			left = mid + 1
		} else if result.start > mid {
			result.start = mid
			right = mid - 1
		} else if result.end < mid {
			result.end = mid
			left = mid + 1
		} else {
			fmt.Println("Should never be logged.")
			return result
		}

	} else if arr[mid] < target {
		left = mid + 1
	} else {
		right = mid - 1
	}

	return search_element_range_optimized(arr, target, left, right, result)
}
