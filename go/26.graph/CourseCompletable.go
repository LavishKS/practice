package main

import (
	"fmt"
)

type CourseSchedularTestCase struct {
	N          int
	Courses    [][]int
	Expected   bool
}

func (t CourseSchedularTestCase) String() string {
	return fmt.Sprintf("\tN          : %d\n\tHeadId     : %v\n\tExpected : %d", t.N, t.Courses, t.Expected)
}

func (t CourseSchedularTestCase) validateResult(result bool) bool {
	return result == t.Expected
}

func testCompleteCourseMain() {
	testCases := []CourseSchedularTestCase{
		{6, [][]int{{1, 0}, {2, 1}, {2, 5}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}, true},
		{6, [][]int{{1, 0}, {2, 1}, {5, 2}, {0, 3}, {4, 3}, {3, 5}, {4, 5}}, false},
		{8, [][]int{{0, 3}, {1, 0}, {2, 1}, {4, 5}, {7, 4}, {5, 7}}, false},
		{0, [][]int{}, true},
		{3, [][]int{{}, {}}, true},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := isCourseCompletable(testCase.N, testCase.Courses)
		fmt.Println("\tActual   :", result)

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
