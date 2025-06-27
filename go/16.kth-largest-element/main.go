package main

import "fmt"

type TestCase struct {
	Input    []int
	K        int
	Expected int
}

func (t TestCase) String() string {
	return fmt.Sprintf("\tInput    : %v\n\tK        : %d\n\tExpected : %d\n", t.Input, t.K, t.Expected)
}

func (t TestCase) validateResult(result int) bool {
	return result == t.Expected
}

func main() {
	testCases := []TestCase{
		{[]int{1, 2, 3, 4}, 2, 3},
		{[]int{5, 3, 1, 6, 4, 2}, 2, 5},
		{[]int{2, 3, 1, 2, 4, 2}, 4, 2},
		{[]int{3}, 1, 3},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := findKthLargeNumber(testCase.Input, testCase.K)
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

func findKthLargeNumber(arr []int, k int) int {
	// quickSortV2(arr, 0, len(arr) - 1)
	// return arr[len(arr)-k]
	return hoareQuickSelectAlgortihm(arr, len(arr)-k, 0, len(arr)-1)
}
