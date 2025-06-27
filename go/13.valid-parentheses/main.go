package main

import (
	"fmt"
	"strings"
)

type RuneStack struct {
	runes []rune
}

func (r RuneStack) String() string {
	var sb strings.Builder
	for _, v := range r.runes {
		sb.WriteRune(v)
	}
	return sb.String()
}

func (s *RuneStack) Push(r rune) {
	s.runes = append(s.runes, r)
}

func (s *RuneStack) Pop() (rune, bool) {
	if len(s.runes) == 0 {
		return 0, false
	}
	lastIdx := len(s.runes) - 1
	r := s.runes[lastIdx]
	s.runes = s.runes[:lastIdx]
	return r, true
}

func (s *RuneStack) Peek() (rune, bool) {
	if len(s.runes) == 0 {
		return 0, false
	}
	return s.runes[len(s.runes)-1], true
}

func (s *RuneStack) IsEmpty() bool {
	return len(s.runes) == 0
}

func (s *RuneStack) Size() int {
	return len(s.runes)
}

type TestCases struct {
	Input    string
	Expected bool
}

func main() {
	testCases := []TestCases{
		{Input: "", Expected: true},
		{Input: "{", Expected: false},
		{Input: "(", Expected: false},
		{Input: "[", Expected: false},
		{Input: "}", Expected: false},
		{Input: ")", Expected: false},
		{Input: "]", Expected: false},
		{Input: "{}", Expected: true},
		{Input: "()", Expected: true},
		{Input: "[]", Expected: true},
		{Input: "{([]", Expected: false},
		{Input: "{([])}", Expected: true},
		{Input: "{()[]}", Expected: true},
		{Input: "{([])]", Expected: false},
		{Input: "{([)]}", Expected: false},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println("\tinput:", testCase.Input)
		result := areParenthesesValid(testCase.Input)
		fmt.Println("\tExpected:", testCase.Expected)
		fmt.Println("\tActual:", result)

		if result == testCase.Expected {
			passed++
			fmt.Println("+ Testcase Passed!")
		} else {
			fmt.Println("- Testcase Failed!")
		}
		fmt.Println()
	}

	fmt.Println(passed, "out of", len(testCases), "testcases passed!")
}

func areParenthesesValid(s string) bool {
	closeBracketMap := map[rune]rune{'}': '{', ')': '(', ']': '['}
	stack := RuneStack{}

	for _, ch := range s {
		expectedBracket, isClosingBracket := closeBracketMap[ch]
		if isClosingBracket {
			openBracket, exists := stack.Pop()
			if !exists || openBracket != expectedBracket {
				return false
			}
		} else {
			stack.Push(ch)
		}
	}
	return stack.IsEmpty()
}
