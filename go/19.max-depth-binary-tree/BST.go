package main

import (
	"fmt"
	"math"
	"strings"
)

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

func (b *BSTNode) insert(node *BSTNode) {
	if node.val > b.val {
		if b.right == nil {
			b.right = node
		} else {
			b.right.insert(node)
		}
	} else {
		if b.left == nil {
			b.left = node
		} else {
			b.left.insert(node)
		}
	}
}

func bstNodeToString(node *BSTNode) string {
	if node == nil {
		return "[nil]"
	}

	format := "[%d]"
	if node.val < 10 {
		format = "[ %d ]"
	} else if node.val < 100 {
		format = "[ %d]"

	}
	return fmt.Sprintf(format, node.val)
}

func binarySearchTreeToStringBuffer(node *BSTNode, level int, index int, buffer *[][]string) {
	if level == len(*buffer) {
		size := int(math.Pow(2, float64(level)))
		newRow := make([]string, size)
		for i := range newRow {
			newRow[i] = "     "
		}
		*buffer = append(*buffer, newRow)
	}

	(*buffer)[level][index] = bstNodeToString(node)

	if node == nil {
		return
	}

	level, index = level+1, index*2
	binarySearchTreeToStringBuffer(node.left, level, index, buffer)
	binarySearchTreeToStringBuffer(node.right, level, index+1, buffer)

}

type BinarySearchTree struct {
	root *BSTNode
}

func (b BinarySearchTree) deapthFirstSearchInOrder() any {
	result := []int{}

	stack := []*BSTNode{}
	stack = append(stack, b.root)

	seen := make(map[*BSTNode]bool)
	seen[b.root] = true
	seen[nil] = true

	for len(stack) != 0 {
		l := len(stack) - 1
		top := stack[l]

		if top == nil {
			stack = stack[:l]
		} else if !seen[top.left] {
			seen[top.left] = true
			stack = append(stack, top.left)
		} else {
			stack = stack[:l]
			result = append(result, top.val)
			stack = append(stack, top.right)
		}
	}

	return result
}

func (b BinarySearchTree) deapthFirstSearchPreOrder() any {
	result := []int{}
	stack := []*BSTNode{}
	stack = append(stack, b.root)

	for len(stack) != 0 {
		l := len(stack) - 1
		top := stack[l]
		stack = stack[:l]

		if top == nil {
			continue
		}
		result = append(result, top.val)
		stack = append(stack, top.right, top.left)
	}

	return result
}

func (b BinarySearchTree) deapthFirstSearchPostOrder() any {
	result := []int{}

	stack := []*BSTNode{}
	stack = append(stack, b.root)

	seen := make(map[*BSTNode]bool)
	seen[b.root] = true
	seen[nil] = true

	for len(stack) != 0 {
		l := len(stack) - 1
		top := stack[l]

		if top == nil {
			stack = stack[:l]
		} else if !seen[top.left] {
			seen[top.left] = true
			stack = append(stack, top.left)
		} else if !seen[top.right] {
			seen[top.right] = true
			stack = append(stack, top.right)
		} else {
			result = append(result, top.val)
			stack = stack[:l]
		}
	}

	return result
}

func (b BinarySearchTree) breadthFirstSearchOrder() []int {
	result := []int{}
	queue := []*BSTNode{}
	queue = append(queue, b.root)

	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]

		if top == nil {
			continue
		}
		result = append(result, top.val)
		queue = append(queue, top.left, top.right)
	}

	return result
}

func (b *BinarySearchTree) insert(v int) {
	node := &BSTNode{v, nil, nil}
	if b.root == nil {
		b.root = node
	} else {
		b.root.insert(node)
	}
}

func (b BinarySearchTree) String() string {
	buffer := [][]string{{"[nil]"}}
	binarySearchTreeToStringBuffer(b.root, 0, 0, &buffer)

	formatedBuffer := make([]string, len(buffer))
	spacing, incremental := "", "   "
	for i := len(buffer) - 1; i >= 0; i-- {
		var sb strings.Builder
		for _, v := range buffer[i] {
			sb.WriteString(spacing + v + spacing + " ")
		}
		spacing += incremental
		incremental += incremental
		formatedBuffer[i] = sb.String()
	}
	var result strings.Builder
	for _, v := range formatedBuffer {
		result.WriteString(v)
		result.WriteRune('\n')
	}
	return result.String()
}

func test_main() {
	tree := BinarySearchTree{}
	//         9
	tree.insert(9)
	//         9
	//    4
	tree.insert(4)
	//         9
	//    4
	//       6
	tree.insert(6)
	//         9
	//    4           20
	//       6
	tree.insert(20)
	//         9
	//    4           20
	//       6            170
	tree.insert(170)
	//         9
	//    4           20
	//       6    15      170
	tree.insert(15)
	//         9
	//    4           20
	// 1     6    15      170
	tree.insert(1)
	tree.insert(5)

	fmt.Println(tree)

	fmt.Println("BFS Order:", tree.breadthFirstSearchOrder())
	fmt.Println("DFS In-Order:", tree.deapthFirstSearchInOrder())
	fmt.Println("DFS Pre-Order:", tree.deapthFirstSearchPreOrder())
	fmt.Println("DFS Post-Order:", tree.deapthFirstSearchPostOrder())
}
