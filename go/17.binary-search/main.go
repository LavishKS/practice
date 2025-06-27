package main

import "fmt"

type TestCase struct {
	Input    []int
	Element  int
	Expected int
}

func (t TestCase) String() string {
	return fmt.Sprintf("\tInput    : %v\n\tK        : %d\n\tExpected : %d\n", t.Input, t.Element, t.Expected)
}

func (t TestCase) validateResult(result int) bool {
	return result == t.Expected
}

func main() {
	testCases := []TestCase{
		{[]int{1, 2, 3, 4}, 2, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 5, 4},
		{[]int{1, 2, 3, 5, 7}, 7, 4},
		{[]int{1, 2, 3, 5, 7}, 8, -1},
		{[]int{1, 2, 3, 5, 7}, 0, -1},
		{[]int{3}, 3, 0},
		{[]int{3}, 2, -1},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := search_element(testCase.Input, testCase.Element)
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

func search_element(arr []int, element int) int {
	// return binary_search_recursive(arr, 0, len(arr)-1, element)
	return binary_search_iterative(arr, element)
}
