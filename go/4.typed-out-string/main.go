package main

import "fmt"

func driveString(str string) string {
	var stack []rune

	for _, ch := range str {
		if ch != '#' {
			stack = append(stack, ch)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}

func areTypedEqualBruteForce(str1 string, str2 string) bool {
	bstr1 := driveString(str1)
	bstr2 := driveString(str2)
	if len(bstr1) != len(bstr2) {
		return false
	}

	for i := 0; i < len(bstr1); i++ {
		if bstr1[i] != bstr2[i] {
			return false
		}
	}

	return true
}

// skipBackspaces processes the string in reverse starting from index `i`,
// skipping characters affected by backspaces ('#') and returns the new index.
func skipBackspaces(str string, i int) int {
	skipCount := 0
	for i > -1 && (skipCount > 0 || str[i] == '#') {
		if str[i] == '#' {
			skipCount++
		} else {
			skipCount--
		}
		i--
	}
	return i
}
func areBackspaceStringsEqual(str1 string, str2 string) bool {
	// Initialize i1 to the last index of str1
	i1 := len(str1) - 1
	// Initialize i2 to the last index of str2
	i2 := len(str2) - 1

	for {
		i1 = skipBackspaces(str1, i1)
		i2 = skipBackspaces(str2, i2)

		if i1 < 0 || i2 < 0 || str1[i1] != str2[i2] {
			break
		}

		i1--
		i2--
	}

	return i1 < 0 && i2 < 0
}

func areTypedEqual(str1 string, str2 string) bool {
	return areBackspaceStringsEqual(str1, str2)
}

type TestCases struct {
	Str1     string
	Str2     string
	Expected bool
}

func main() {
	testCases := []TestCases{
		{Str1: "cb#d", Str2: "cd", Expected: true},
		{Str1: "cb#d", Str2: "c#d", Expected: false},
		{Str1: "cb#d", Str2: "##cd", Expected: true},
		{Str1: "ab#c", Str2: "az#c", Expected: true},
		{Str1: "ab#c", Str2: "abc", Expected: false},
		{Str1: "ab##", Str2: "", Expected: true},
		{Str1: "ab##", Str2: "ab", Expected: false},
		{Str1: "a###b", Str2: "b", Expected: true},
		{Str1: "a###b", Str2: "a", Expected: false},
		{Str1: "x#y#z#", Str2: "a#", Expected: true},
		{Str1: "ab#c", Str2: "Ab#c", Expected: false},
		{Str1: "abcasdf#c##jkl####", Str2: "abca", Expected: true},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println("\tinput1:", testCase.Str1)
		fmt.Println("\tinput2:", testCase.Str2)
		result := areTypedEqual(testCase.Str1, testCase.Str2)
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

	fmt.Println(passed, "testcases passed out of", len(testCases))
}
