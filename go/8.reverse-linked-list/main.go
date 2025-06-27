package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

func (n Node) String() string {
	return fmt.Sprintf("%d --> %s", n.data, n.next)
}

type LinkedList struct {
	head *Node
}

func (ll LinkedList) String() string {
	return fmt.Sprintf("head --> %s", ll.head)
}

type TestCases struct {
	Input    LinkedList
	Expected LinkedList
}

func arrayToLinkedList(arr []int) LinkedList {
	ll := LinkedList{}
	if len(arr) == 0 {
		return ll
	}

	ll.head = &Node{data: arr[0], next: nil}
	current := ll.head

	for i := 1; i < len(arr); i++ {
		newNode := &Node{data: arr[i], next: nil}
		current.next = newNode
		current = newNode
	}
	return ll
}

func compareLinkedLists(ll1 LinkedList, ll2 LinkedList) bool {
	l1, l2 := ll1.head, ll2.head

	for l1 != nil && l2 != nil {
		if l1.data != l2.data {
			return false
		}
		l1 = l1.next
		l2 = l2.next
	}

	return l1 == nil && l2 == nil
}

func main() {
	testCases := []TestCases{
		{Input: arrayToLinkedList([]int{1, 2, 3, 4, 5}), Expected: arrayToLinkedList([]int{5, 4, 3, 2, 1})},
		{Input: arrayToLinkedList([]int{3}), Expected: arrayToLinkedList([]int{3})},
		{Input: arrayToLinkedList([]int{}), Expected: arrayToLinkedList([]int{})},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println("\tinput:", testCase.Input)
		result := reverseLinkedListInPlace(testCase.Input)
		fmt.Println("\tExpected:", testCase.Expected)
		fmt.Println("\tActual:", result)

		if compareLinkedLists(result, testCase.Expected) {
			passed++
			fmt.Println("+ Testcase Passed!")
		} else {
			fmt.Println("- Testcase Failed!")
		}
		fmt.Println()
	}

	fmt.Println(passed, "out of", len(testCases), "testcases passed!")
}

func reverseLinkedList(linkedList LinkedList) LinkedList {
	result := LinkedList{}
	current := linkedList.head
	for current != nil {
		newHead := &Node{data: current.data, next: result.head}
		result.head = newHead
		current = current.next
	}
	return result
}

func reverseLinkedListInPlace(ll LinkedList) LinkedList {
	if ll.head == nil {
		return ll
	}
	current := ll.head.next
	ll.head.next = nil
	for current != nil {
		nextNode := current.next
		current.next = ll.head
		ll.head = current
		current = nextNode
	}

	return ll
}
