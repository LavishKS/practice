/**
 * Given a string only containing round brackets '(' and ')'
 * and lowercase characters, remove the least amount of brackets so the string is valid.
 *
 * A string is considered valid if it is empty or if there are brackets, they all close.
 *
 * Ex:
 *  1. "a)bc(d)" => "abc(d)"            : 1
 *  2. "(ab(c)a" => "ab(c)a"/"(abc)a"   : 1
 *  3. "(("      => ""                  : 2
 *
 *
 * Constraints:
 *       Algorithm Returns: Valid String
 *       No spaces in the string, only lowercase characters, '(' and ')'.
 *       A string without any brackets is a valid string.
 */

package main

import (
	"fmt"
	"strings"
)

type IntStack struct {
	indexData []int
	charData  []rune
}

func (s *IntStack) Push(i int, r rune) {
	s.indexData = append(s.indexData, i)
	s.charData = append(s.charData, r)
}

func (s *IntStack) Pop() (int, rune, bool) {
	if len(s.indexData) == 0 {
		return 0, 0, false
	}
	lastIdx := len(s.indexData) - 1
	i := s.indexData[lastIdx]
	s.indexData = s.indexData[:lastIdx]
	r := s.charData[lastIdx]
	s.charData = s.charData[:lastIdx]
	return i, r, true
}

func (s *IntStack) Peek() (int, rune, bool) {
	lastIdx := len(s.indexData) - 1
	if lastIdx < 0 {
		return 0, 0, false
	}

	return s.indexData[lastIdx], s.charData[lastIdx], true
}

func (s *IntStack) IsEmpty() bool {
	return len(s.indexData) == 0
}

func (s *IntStack) Size() int {
	return len(s.indexData)
}

type TestCase struct {
	Input    string
	Expected map[string]bool
}

func (t TestCase) validateResult(s string) bool {
	return t.Expected[s]
}

func main() {
	testCases := []TestCase{
		{Input: "", Expected: map[string]bool{"": true}},
		{Input: "a)bc(d)", Expected: map[string]bool{"abc(d)": true}},
		{Input: "(ab(c)d", Expected: map[string]bool{"ab(c)d": true, "(abc)d": true}},
		{Input: "((", Expected: map[string]bool{"": true}},
		{Input: "))((", Expected: map[string]bool{"": true}},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println("\tinput:", testCase.Input)
		result := makeStringValidV2(testCase.Input)
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

func makeStringValidV1(s string) string {
	indexStack := IntStack{}

	for i, ch := range s {
		if ch == '(' {
			indexStack.Push(i, ch)
		} else if ch == ')' {
			_, topRune, exists := indexStack.Peek()
			if exists && topRune == '(' {
				indexStack.Pop()
			} else {
				indexStack.Push(i, ch)
			}
		}
	}

	var sb strings.Builder
	index := 0
	for i, ch := range s {
		if index < len(indexStack.indexData) && indexStack.indexData[index] == i {
			index++
			continue
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}

func makeStringValidV2(s string) string {
	const openBracket = '('
	const closeBracket = ')'
	const bogus = '_'
	runes := []rune(s)
	indexStack := IntStack{}
	for i, ch := range runes {
		if ch == openBracket {
			indexStack.Push(i, bogus)
		} else if ch == closeBracket {
			index, _, exists := indexStack.Peek()
			if exists && runes[index] == openBracket {
				indexStack.Pop()
			} else {
				indexStack.Push(i, bogus)
			}
		}
	}

	var sb strings.Builder
	index := 0

	for i, ch := range runes {
		if index < len(indexStack.indexData) && indexStack.indexData[index] == i {
			index++
			continue
		}
		sb.WriteRune(ch)
	}

	return sb.String()
}
