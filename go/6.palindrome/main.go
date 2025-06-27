package main

import (
	"fmt"
	"regexp"
	"strings"
)

type TestCases struct {
	Str      string
	Expected bool
}

func main() {
	testCases := []TestCases{
		{Str: "aabaa", Expected: true},
		{Str: "aabbaa", Expected: true},
		{Str: "abc", Expected: false},
		{Str: "a", Expected: true},
		{Str: "", Expected: true},
		{Str: "A man, a plan, a canal: Panama", Expected: true},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println("\tinput:", testCase.Str)
		result := isValidPalindrome(testCase.Str)
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

func stripString(str string) string {
	// Remove all non-alphanumeric characters and convert to lowercase
	reg := regexp.MustCompile("[^a-zA-Z0-9]")
	return strings.ToLower(reg.ReplaceAllString(str, ""))
}

func isValidPalindrome(str string) bool {
	str = stripString(str)
	return isValidPalindromeV3(str)
}

func reverseString(str string) string {
	var reverseStr strings.Builder
	reverseStr.Grow(len(str))

	for i := len(str) - 1; i >= 0; i-- {
		reverseStr.WriteByte(str[i])
	}

	return reverseStr.String()
}

func isValidPalindromeV3(str string) bool {
	reverseStr := reverseString(str)
	for i := 0; i < len(str); i++ {
		if str[i] != reverseStr[i] {
			return false
		}
	}

	return true
}

func isValidPalindromeV2(str string) bool {
	length := len(str)
	left, right := (length-1)/2, length/2
	fmt.Println("  Plain Str:", str)
	fmt.Println("  Length:", length, "Left:", left, "Right:", right)
	for right < length {
		if str[left] != str[right] {
			return false
		}
		left--
		right++
	}
	return true
}

func isValidPalindromeV1(str string) bool {
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}
