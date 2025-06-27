package main

import "fmt"

type TestCases struct {
	Str      string
	Expected int
}

func main() {
	testCases := []TestCases{
		{Str: "abccabb", Expected: 3},
		{Str: "cccccc", Expected: 1},
		{Str: "", Expected: 0},
		{Str: "abcbda", Expected: 4},
		{Str: "abcbdca", Expected: 4},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println("\tinput:", testCase.Str)
		result := nonRepeatingSubstring(testCase.Str)
		fmt.Println("\tExpected:", testCase.Expected)
		fmt.Println("\tActual:", result)

		if result == testCase.Expected {
			passed++
			fmt.Println("+ Testcase Pased!")
		} else {
			fmt.Println("- Testcase Failed!")
		}
		fmt.Println()
	}

	fmt.Println(passed, "out of", len(testCases), "testcases passed!")
}

func nonRepeatingSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	return nonRepeatingSubstringSlidingWindow(s)
}

func nonRepeatingSubstringSlidingWindow(s string) int {
	result := 0
	left := 0
	seen := map[rune]int{}
	for right, ch := range s {
		index, hasSeen := seen[ch]
		if hasSeen && index >= left {
			left = index + 1
		}
		seen[ch] = right
		result = max(right-left+1, result)
	}
	return result
}

func nonRepeatingSubstringBruteForce(s string) int {
	result := 0
	for i, ch := range s {
		seen := map[rune]struct{}{}
		seen[ch] = struct{}{}
		for _, och := range s[i+1:] {
			_, exists := seen[och]
			if exists {
				break
			}
			seen[och] = struct{}{}
		}

		if len(seen) > result {
			result = len(seen)
		}
	}

	return result
}
