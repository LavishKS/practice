/**
 * Given a Binary Tree, Find it's maximum depth.
 */
package main

import "fmt"

type TestCase struct {
	Input    BinaryTree
	Expected int
}

func (t TestCase) String() string {
	return fmt.Sprintf("\tInput    :\n%s\n\tExpected : %d\n", t.Input, t.Expected)
}

func (t TestCase) validateResult(result int) bool {
	return result == t.Expected
}

func main() {
	testCases := []TestCase{
		{fromArrayToBT([]int{1, 3, 3, 5, 5, 5, 8, 9}), 4},
		{fromArrayToBT([]int{1}), 1},
		{fromArrayToBT([]int{}), 0},
		{fromArrayToBT([]int{1, -1, 1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1}), 6},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := find_max_binary_tree_depth(testCase.Input)
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

func find_max_binary_tree_depth(b BinaryTree) int {
	// return calc_binary_tree_max_depth(b.root, 0)
	return calc_tree_depth_iterative(b.root)
}

func calc_binary_tree_max_depth(bTNode *BTNode, depth int) int {
	if bTNode == nil {
		return depth
	}

	depth += 1
	return max(
		calc_binary_tree_max_depth(bTNode.left, depth),
		calc_binary_tree_max_depth(bTNode.right, depth),
	)
}

func calc_tree_depth_iterative(bTNode *BTNode) int {
	maxheight := 0
	if bTNode == nil {
		return maxheight
	}

	stack := []*BTNode{bTNode}
	traversed := make(map[*BTNode]bool)
	traversed[nil] = true
	
	currentHeight := len(stack)
	for currentHeight > 0 {
		if currentHeight > maxheight {
			maxheight = currentHeight
		}

		top := stack[currentHeight-1]

		if !traversed[top.left] {
			stack = append(stack, top.left)
		} else if !traversed[top.right] {
			stack = append(stack, top.right)
		} else {
			traversed[top] = true
			stack = stack[:currentHeight-1]
		}

		currentHeight = len(stack)
	}

	return maxheight
}
