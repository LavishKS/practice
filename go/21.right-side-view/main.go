package main

import (
	"fmt"
	"reflect"
)

type TestCase struct {
	Input    BinaryTree
	Expected []int
}

func (t TestCase) String() string {
	// return fmt.Sprintf("\tInput    :\n%s\n\tExpected : %d\n", t.Input, t.Expected)
	return fmt.Sprintf("\tExpected : %d\n", t.Expected)
}

func (t TestCase) validateResult(result []int) bool {
	return reflect.DeepEqual(result, t.Expected)
}

func main() {
	testCases := []TestCase{
		{fromArrayToBT([]int{3, 6, 1, 9, 2, -1, 4, -1, 5, -1, -1, -1, -1, -1, -1, -1, -1, 8}), []int{3, 1, 4, 5, 8}},
		{fromArrayToBT([]int{1, 2, 3, 4, 5, -1, 6, -1, 7, -1, -1, -1, -1, -1, -1, -1, -1, 8}), []int{1, 3, 6, 7, 8}},
		{fromArrayToBT([]int{1, 3, 3, 5, 5, 5, 8, 9}), []int{1, 3, 8, 9}},
		{fromArrayToBT([]int{1}), []int{1}},
		{fromArrayToBT([]int{}), []int{}},
		{fromArrayToBT([]int{1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1}), []int{1, 1, 1, 1, 1, 1}},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := rightSideViewBFS(testCase.Input)
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

func rightSideViewBFS(b BinaryTree) []int {
	result := []int{}
	if b.root == nil {
		return result
	}

	currentLevel := []*BTNode{b.root}

	for len(currentLevel) > 0 {
		result = append(result, currentLevel[len(currentLevel)-1].val)
		nextLevel := []*BTNode{}
		for _, node := range currentLevel {
			if node.left != nil {
				nextLevel = append(nextLevel, node.left)
			}
			if node.right != nil {
				nextLevel = append(nextLevel, node.right)
			}
		}
		currentLevel = nextLevel
	}

	return result
}

func rightSideViewDFS(b BinaryTree) []int {
	result := []int{}
	if b.root == nil {
		return result
	}
	return rightTreeView(b.root, 0, result)
}

func rightTreeView(node *BTNode, level int, result []int) []int {
	if node == nil {
		return result
	}
	if level == len(result) {
		result = append(result, node.val)
	}

	level++
	result = rightTreeView(node.right, level, result)
	return rightTreeView(node.left, level, result)
}
