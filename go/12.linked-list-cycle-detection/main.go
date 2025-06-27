package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head       *Node
	tail       *Node
	size       int
	cyclicNode *Node
}

func (l *LinkedList) push(node *Node, cyclicNode bool) {
	if cyclicNode {
		node.next = node
		l.cyclicNode = node
	}

	if l.head == nil {
		l.head = node
	} else {
		if !cyclicNode {
			node.next = l.tail.next
		}
		l.tail.next = node
	}

	l.size++
	l.tail = node
}

func (ll LinkedList) String() string {
	var sb strings.Builder
	sb.WriteString("head --> ")
	for i, size := ll.head, ll.size; size > 0 && i != nil; i, size = i.next, size-1 {
		sb.WriteString(fmt.Sprintf("%d --> ", i.data))
	}
	if ll.tail == nil || ll.tail.next == nil {
		sb.WriteString("<nil>")
	} else {
		sb.WriteString(" Cycle back to ")
		sb.WriteString(fmt.Sprintf("node: %d", ll.cyclicNode.data))
	}
	return sb.String()
}

type TestCases struct {
	Input LinkedList
}

func (t TestCases) isCyclicNode(result *Node) bool {
	return result == t.Input.cyclicNode
}

func arrayToLinkedList(arr []int, cycleAt int) LinkedList {
	ll := LinkedList{}

	for i, data := range arr {
		ll.push(&Node{data: data, next: nil}, i == cycleAt)
	}

	return ll
}

func main() {
	testCases := []TestCases{
		{Input: arrayToLinkedList([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4)},
		{Input: arrayToLinkedList([]int{1, 2, 3, 4, 5}, 2)},
		{Input: arrayToLinkedList([]int{1, 2, 3, 4, 5}, -1)},
		{Input: arrayToLinkedList([]int{3}, 0)},
		{Input: arrayToLinkedList([]int{}, 0)},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println("\tinput:", testCase.Input)
		result := detectCycle(testCase.Input)
		fmt.Println("\tExpected:", testCase.Input.cyclicNode)
		fmt.Println("\tActual:", result)

		if testCase.isCyclicNode(result) {
			passed++
			fmt.Println("+ Testcase Passed!")
		} else {
			fmt.Println("- Testcase Failed!")
		}
		fmt.Println()
	}

	fmt.Println(passed, "out of", len(testCases), "testcases passed!")
}

func detectCycle(ll LinkedList) *Node {
	return detectCycleFloydTH(ll)
}

func detectCycleUsingCache(ll LinkedList) *Node {
	seen := make(map[int]bool)
	for current := ll.head; current != nil; current = current.next {
		if seen[current.data] {
			return current
		}
		seen[current.data] = true
	}

	fmt.Println("No cycle detected")
	return nil
}

func detectCycleFloydTH(ll LinkedList) *Node {
	meetingNode := getCycleMeetingNode(ll)

	if meetingNode != nil {
		head := ll.head
		for head != meetingNode {
			head = head.next
			meetingNode = meetingNode.next
		}
	}

	return meetingNode
}

func getCycleMeetingNode(ll LinkedList) *Node {
	tortoise, hare := ll.head, ll.head

	for hare != nil {
		tortoise = tortoise.next
		if hare.next == nil || hare.next.next == nil {
			return nil
		}
		hare = hare.next.next
		if tortoise == hare {
			return hare
		}
	}

	return nil
}
