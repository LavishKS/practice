package main

import (
	"fmt"
)

type InformEmployeesTestCase struct {
	N          int
	HeadId     int
	Managers   []int
	InformTime []int
	Expected   int
}

func (t InformEmployeesTestCase) String() string {
	return fmt.Sprintf("\tN          : %d\n\tHeadId     : %d\n\tManagers   : %v\n\tInformTime : %v\n\tExpected : %d", t.N, t.HeadId, t.Managers, t.InformTime, t.Expected)
}

func (t InformEmployeesTestCase) validateResult(result int) bool {
	return result == t.Expected
}

func testInformEmployeesMain() {
	testCases := []InformEmployeesTestCase{
		{8, 4, []int{2, 2, 4, 6, -1, 4, 4, 5}, []int{0, 0, 4, 0, 7, 3, 6, 0}, 13},
		{1, 0, []int{-1}, []int{0}, 0},
		{7, 6, []int{1, 2, 3, 4, 5, 6, -1}, []int{0, 6, 5, 4, 3, 2, 1}, 21},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := calcTimeToInform(testCase.N, testCase.HeadId, testCase.Managers, testCase.InformTime)
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

func fromManagersArrayToGraph(size int, managers []int) AdjacencyListGraph {
	graph := AdjacencyListGraph{}
	graph.init(size)

	for emp, manager := range managers {
		if manager > 0 {
			graph.addDirectionalEdge(manager, emp)
		}
	}

	return graph
}

func calcTimeToInform(n, headId int, managers, informTime []int) int {
	graph := fromManagersArrayToGraph(n, managers)
	return headId
}
