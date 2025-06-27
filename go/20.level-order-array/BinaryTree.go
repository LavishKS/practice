package main

import (
	"fmt"
	"math"
	"strings"
)

type BTNode struct {
	val   int
	left  *BTNode
	right *BTNode
}

func btNodeToString(node *BTNode) string {
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

func binaryTreeToStringBuffer(node *BTNode, level int, index int, buffer *[][]string) {
	if level == len(*buffer) {
		size := int(math.Pow(2, float64(level)))
		newRow := make([]string, size)
		for i := range newRow {
			newRow[i] = "     "
		}
		*buffer = append(*buffer, newRow)
	}

	(*buffer)[level][index] = btNodeToString(node)

	if node == nil {
		return
	}

	level, index = level+1, index*2
	binaryTreeToStringBuffer(node.left, level, index, buffer)
	binaryTreeToStringBuffer(node.right, level, index+1, buffer)

}

type BinaryTree struct {
	root *BTNode
}

func (b BinaryTree) String() string {
	buffer := [][]string{{"[nil]"}}
	binaryTreeToStringBuffer(b.root, 0, 0, &buffer)

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

func fromArrayToBT(tArr []int) BinaryTree {
	result := BinaryTree{}

	if len(tArr) < 1 {
		return result
	}

	result.root = &BTNode{val: tArr[0], left: nil, right: nil}
	nodes := make([]*BTNode, len(tArr))
	nodes[0] = result.root

	for i, v := range tArr[1:] {
		if v < 0 {
			continue
		}
		node := &BTNode{val: v, left: nil, right: nil}
		nodes[i+1] = node
		parentIndex := i / 2
		if i%2 == 0 {
			nodes[parentIndex].left = node
		} else {
			nodes[parentIndex].right = node
		}
	}

	result.root = nodes[0]

	return result
}
