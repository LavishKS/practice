package main

import (
	"fmt"
	"math"
)

type TestCase struct {
	Input    BinaryTree
	Expected bool
}

func (t TestCase) String() string {
	// return fmt.Sprintf("\tInput    :\n%s\n\tExpected : %d\n", t.Input, t.Expected)
	return fmt.Sprintf("\tExpected : %v\n", t.Expected)
}

func (t TestCase) validateResult(result bool) bool {
	return result == t.Expected
}

func main() {
	testCases := []TestCase{
		{fromArrayToBT([]int{}), true},
		{fromArrayToBT([]int{1}), true},
		{fromArrayToBT([]int{12, 8, 18, 5, 10, 14, 25}), true},
		{fromArrayToBT([]int{16, 8, 22, 9, -1, 19, 25}), false},
		{fromArrayToBT([]int{13, 6, 17, 2, -1, 10, 22}), false},
		{fromArrayToBT([]int{15, 12, 17, 10, 16, 16, 18}), false},
		{fromArrayToBT([]int{15, 12, 18, 10, 14, 13, 20}), false},
		{fromArrayToBT([]int{12, 15, 17}), false},
		{fromArrayToBT([]int{20, 15, 17}), false},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := validateBST(testCase.Input)
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

func validateBST(b BinaryTree) bool {
	return validateBSTDFS(b.root, math.MinInt32, math.MaxInt32)
}

func validateBSTDFS(node *BTNode, minVal int, maxVal int) bool {
	if node == nil {
		return true
	}

	if node.val >= maxVal || node.val <= minVal {
		return false
	}

	return validateBSTDFS(node.left, minVal, node.val) && validateBSTDFS(node.right, node.val, maxVal)
}
