package main

import (
	"fmt"
	"strings"
)

type IntStack struct {
	data []int
}

func (s *IntStack) push(v int) {
	s.data = append(s.data, v)
}

func (s *IntStack) pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	lastIdx := len(s.data) - 1
	i := s.data[lastIdx]
	s.data = s.data[:lastIdx]
	return i, true
}

// func (s *IntStack) peek() (int, bool) {
// 	lastIdx := len(s.data) - 1
// 	if lastIdx < 0 {
// 		return 0, false
// 	}

// 	return s.data[lastIdx], true
// }

func (s *IntStack) isEmpty() bool {
	return len(s.data) == 0
}

// func (s *IntStack) size() int {
// 	return len(s.data)
// }

type StackQueue struct {
	in  IntStack
	out IntStack
}

func (q *StackQueue) enqueue(v int) {
	q.in.push(v)
}

func (q *StackQueue) dequeue() (int, bool) {
	if q.out.isEmpty() {
		q._moveStack()
	}

	return q.out.pop()
}

func (q *StackQueue) _moveStack() {
	for v, exists := q.in.pop(); exists; v, exists = q.in.pop() {
		q.out.push(v)
	}
}

// func (q *StackQueue) peek() (int, bool) {
// 	if q.out.isEmpty() {
// 		q._moveStack()
// 	}

// 	return q.out.peek()
// }

// func (q *StackQueue) isEmpty() bool {
// 	return q.size() == 0
// }

// func (q *StackQueue) size() int {
// 	return q.out.size() + q.in.size()
// }

type QueueOp int

const (
	Enqueue QueueOp = iota
	Dequeue
)

type QueueOps struct {
	op  QueueOp
	val int
}

type TestCase struct {
	Input    []QueueOps
	Expected string
}

func (t TestCase) validateResult(s string) bool {
	return s == t.Expected
}

func main() {
	testCases := []TestCase{
		{Input: []QueueOps{{op: Dequeue}}, Expected: "ERROR"},
		{Input: []QueueOps{{op: Enqueue, val: 1}, {op: Enqueue, val: 2}, {op: Dequeue}, {op: Dequeue}}, Expected: "12"},
		{Input: []QueueOps{{op: Enqueue, val: 1}, {op: Enqueue, val: 2}, {op: Dequeue}, {op: Dequeue}, {op: Dequeue}}, Expected: "ERROR"},
		{Input: []QueueOps{{op: Enqueue, val: 1}, {op: Enqueue, val: 2}, {op: Dequeue}, {op: Enqueue, val: 3}, {op: Enqueue, val: 4}, {op: Enqueue, val: 5}, {op: Enqueue, val: 6}, {op: Dequeue}, {op: Dequeue}, {op: Dequeue}}, Expected: "1234"},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println("\tinput:", testCase.Input)
		result := performQueueOps(testCase.Input)
		fmt.Println("\tExpected:", testCase.Expected)
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

func performQueueOps(ops []QueueOps) string {
	q := StackQueue{}
	var sb strings.Builder
	for _, op := range ops {
		if op.op == Enqueue {
			q.enqueue(op.val)
		} else {
			val, exists := q.dequeue()
			if !exists {
				return "ERROR"
			}
			sb.WriteString(fmt.Sprint(val))
		}
	}

	return sb.String()
}
