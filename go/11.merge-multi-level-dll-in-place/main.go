package main

import (
	"fmt"
)

type MultiLevelListNode struct {
	prev  *MultiLevelListNode
	data  int
	next  *MultiLevelListNode
	child *MultiLevelLinkedList
}

func (n MultiLevelListNode) String() string {
	return fmt.Sprintf("{p: %p, v: %d, n: %p} --> %s", n.prev, n.data, n.next, n.next)
}

type MultiLevelLinkedList struct {
	head *MultiLevelListNode
}

func (ll MultiLevelLinkedList) String() string {
	return fmt.Sprintf("head --> %s", ll.head)
}

func generateMultiLevelLinkedList(input MultiLevelArrayInput) *MultiLevelLinkedList {
	mlll := MultiLevelLinkedList{}
	vals := input.vals
	if len(vals) < 1 {
		return &mlll
	}
	mlll.head = &MultiLevelListNode{prev: nil, data: vals[0], next: nil, child: nil}
	children, exists := input.children[0]
	if exists {
		mlll.head.child = generateMultiLevelLinkedList(*children)
	}
	for i, current := 1, mlll.head; i < len(vals); i++ {
		current.next = &MultiLevelListNode{prev: current, data: vals[i], next: nil, child: nil}
		children, exists := input.children[i]
		if exists {
			current.next.child = generateMultiLevelLinkedList(*children)
		}
		current = current.next
	}
	return &mlll
}

type MultiLevelArrayInput struct {
	vals     []int
	children map[int]*MultiLevelArrayInput
}

type TestCases struct {
	Input    MultiLevelLinkedList
	Expected MultiLevelLinkedList
}

func compareLinkedLists(dll1 *MultiLevelLinkedList, dll2 *MultiLevelLinkedList) bool {
	if dll1 == nil || dll2 == nil {
		return dll1 == dll2
	}

	l1, l2 := dll1.head, dll2.head

	for l1 != nil && l2 != nil {
		if l1.data != l2.data || !compareLinkedLists(l1.child, l2.child) {
			return false
		}
		l1 = l1.next
		l2 = l2.next
	}

	return l1 == nil && l2 == nil
}

func main() {
	testCases := []TestCases{
		{Input: *generateMultiLevelLinkedList(MultiLevelArrayInput{vals: []int{1, 2, 3, 4}, children: map[int]*MultiLevelArrayInput{1: {vals: []int{5, 6}}}}), Expected: *generateMultiLevelLinkedList(MultiLevelArrayInput{vals: []int{1, 2, 5, 6, 3, 4}, children: nil})},
		{Input: *generateMultiLevelLinkedList(MultiLevelArrayInput{vals: []int{1, 2, 3}, children: map[int]*MultiLevelArrayInput{1: {vals: []int{4, 5, 6}, children: map[int]*MultiLevelArrayInput{1: {vals: []int{7}}}}}}), Expected: *generateMultiLevelLinkedList(MultiLevelArrayInput{vals: []int{1, 2, 4, 5, 7, 6, 3}})},
		{Input: *generateMultiLevelLinkedList(MultiLevelArrayInput{vals: []int{1, 7, 8, 9, 13}, children: map[int]*MultiLevelArrayInput{0: {vals: []int{2, 6}, children: map[int]*MultiLevelArrayInput{0: {vals: []int{3, 4, 5}}}}, 3: {vals: []int{10, 11}, children: map[int]*MultiLevelArrayInput{1: {vals: []int{12}}}}}}), Expected: *generateMultiLevelLinkedList(MultiLevelArrayInput{vals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}})},
		{Input: *generateMultiLevelLinkedList(MultiLevelArrayInput{vals: []int{3}}), Expected: *generateMultiLevelLinkedList(MultiLevelArrayInput{vals: []int{3}})},
		{Input: *generateMultiLevelLinkedList(MultiLevelArrayInput{vals: []int{}}), Expected: *generateMultiLevelLinkedList(MultiLevelArrayInput{vals: []int{}})},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		result := mergeMultiLevelDLL(testCase.Input)
		fmt.Println("\tExpected:", testCase.Expected)
		fmt.Println("\tActual:", result)

		if compareLinkedLists(&result, &testCase.Expected) {
			passed++
			fmt.Println("+ Testcase Passed!")
		} else {
			fmt.Println("- Testcase Failed!")
		}
		fmt.Println()
	}

	fmt.Println(passed, "out of", len(testCases), "testcases passed!")
}

func mergeMultiLevelDLL(mlll MultiLevelLinkedList) MultiLevelLinkedList {
	for current := mlll.head; current != nil; current = current.next {
		if current.child != nil {
			childTail := current.child.head
			for childTail.next != nil {
				childTail = childTail.next
			}
			childTail.next = current.next
			if current.next != nil {
				current.next.prev = childTail
			}
			current.child.head.prev = current
			current.next = current.child.head
			current.child = nil
		}
	}
	return mlll
}
