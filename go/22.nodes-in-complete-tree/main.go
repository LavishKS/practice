package main

import (
	"fmt"
)

type TestCase struct {
	Input    BinaryTree
	Expected int
}

func (t TestCase) String() string {
	// return fmt.Sprintf("\tInput    :\n%s\n\tExpected : %d\n", t.Input, t.Expected)
	return fmt.Sprintf("\tExpected : %d\n", t.Expected)
}

func (t TestCase) validateResult(result int) bool {
	return result == t.Expected
}

func main() {
	testCases := []TestCase{
		{fromArrayToBT([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}), 15},
		{fromArrayToBT([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}), 12},
		{fromArrayToBT([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}), 11},
		{fromArrayToBT([]int{1, 1, 1, 1, 1, 1, 1, 1}), 8},
		{fromArrayToBT([]int{1}), 1},
		{fromArrayToBT([]int{}), 0},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := countNode(testCase.Input)
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

func countNode(b BinaryTree) int {
	return countNodesInCompleteTree(b.root)
}

func powInt(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func countNodesInCompleteTree(btNode *BTNode) int {
	if btNode == nil {
		return 0
	}

	height := getTreeHeight(btNode)
	lastLevelNodeMaxCount := powInt(2, height)

	left, right := 1, lastLevelNodeMaxCount

	for left < right {
		nodePosition := (left + right + 1) / 2
		node := getNodeAtPosition(btNode, lastLevelNodeMaxCount, height, nodePosition)
		if node == nil {
			right = nodePosition - 1
		} else {
			left = nodePosition
		}
	}

	return lastLevelNodeMaxCount + left - 1
}

func getNodeAtPosition(btNode *BTNode, lastLevelNodeMaxCount int, height int, nodePosition int) *BTNode {
	node := btNode
	start, end := 1, lastLevelNodeMaxCount
	for range height {
		currentPos := (end + start) / 2
		if nodePosition > currentPos {
			node = node.right
			start = currentPos + 1
		} else {
			node = node.left
			end = currentPos
		}
	}
	return node
}

func getTreeHeight(btNode *BTNode) int {
	height := 0
	for n := btNode; n.left != nil; n = n.left {
		height++
	}
	return height
}

func countNodeBruteForceBFS(bTNode *BTNode) int {
	if bTNode == nil {
		return 0
	}
	queue := []*BTNode{bTNode}

	for i := 0; i < len(queue); i++ {
		node := queue[i]
		if node.left == nil {
			break
		}
		queue = append(queue, node.left)
		if node.right == nil {
			break
		}
		queue = append(queue, node.right)
	}
	return len(queue)

}

func countNodeBruteForceDFS(node *BTNode, result int) int {
	if node == nil {
		return result
	}

	result = countNodeBruteForceDFS(node.left, result+1)
	return countNodeBruteForceDFS(node.right, result)
}
