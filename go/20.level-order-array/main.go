package main

import (
	"fmt"
	"reflect"
)

type TestCase struct {
	Input    BinaryTree
	Expected [][]int
}

func (t TestCase) String() string {
	// return fmt.Sprintf("\tInput    :\n%s\n\tExpected : %d\n", t.Input, t.Expected)
	return fmt.Sprintf("\tExpected : %d\n", t.Expected)
}

func (t TestCase) validateResult(result [][]int) bool {
	return reflect.DeepEqual(result, t.Expected)
}

func main() {
	testCases := []TestCase{
		{fromArrayToBT([]int{3, 6, 1, 9, 2, -1, 4, -1, 5, -1, -1, -1, -1, -1, -1, -1, -1, 8}), [][]int{{3}, {6, 1}, {9, 2, 4}, {5}, {8}}},
		{fromArrayToBT([]int{1, 3, 3, 5, 5, 5, 8, 9}), [][]int{{1}, {3, 3}, {5, 5, 5, 8}, {9}}},
		{fromArrayToBT([]int{1}), [][]int{{1}}},
		{fromArrayToBT([]int{}), [][]int{}},
		{fromArrayToBT([]int{1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1}), [][]int{{1}, {1}, {1}, {1}, {1}, {1}}},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := levelOrderTraversalTwoQueues(testCase.Input)
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

type NodeLevelPair struct {
	node  *BTNode
	level int
}

func levelOrderTraversal(t BinaryTree) [][]int {
	result := [][]int{}
	if t.root == nil {
		return result
	}

	queue := []NodeLevelPair{{t.root, 0}}

	// Could be optimized using indexes
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		node, level := top.node, top.level
		if node != nil {
			if level == len(result) {
				result = append(result, []int{})
			}
			result[level] = append(result[level], node.val)
			level += 1
			queue = append(queue, NodeLevelPair{node.left, level}, NodeLevelPair{node.right, level})
		}
	}

	return result
}

func levelOrderTraversalTwoQueues(t BinaryTree) [][]int {
	result := [][]int{}
	if t.root == nil {
		return result
	}

	currentLevel := []*BTNode{t.root}

	for levelCount := len(currentLevel); levelCount > 0; levelCount = len(currentLevel) {
		i := len(result)
		result = append(result, make([]int, levelCount, levelCount))
		nextLevel := []*BTNode{}
		for j, node := range currentLevel {
			result[i][j] = node.val
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
