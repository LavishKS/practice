package main

import (
	"fmt"
)

type TestCases struct {
	Str      string
	Expected bool
}

func main() {
	testCases := []TestCases{
		{Str: "raceacar", Expected: true},
		{Str: "abccdba", Expected: true},
		{Str: "abcdefdba", Expected: false},
		{Str: "aabaa", Expected: true},
		{Str: "aabbaa", Expected: true},
		{Str: "", Expected: true},
		{Str: "a", Expected: true},
		{Str: "ab", Expected: true},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println("\tinput:", testCase.Str)
		result := isAlmostPalindrome(testCase.Str)
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

func isAlmostPalindrome(s string) bool {
	return isAlmostPalindromeV1(s)
}

func isValidPalindrome(str string, l int, r int) bool {
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func isAlmostPalindromeV1(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isValidPalindrome(s, l, r-1) || isValidPalindrome(s, l+1, r)
		}
		l++
		r--
	}

	return true
}
