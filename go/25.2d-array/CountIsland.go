package main

import (
	"fmt"
)

type CountIslandTestCase struct {
	Input    [][]int
	Expected int
}

func (t CountIslandTestCase) String() string {
	return fmt.Sprintf("\tInput    : %v\n\tExpected : %d", t.Input, t.Expected)
}

func (t CountIslandTestCase) validateResult(result int) bool {
	return result == t.Expected
}

func testCountIslandMain() {
	testCases := []CountIslandTestCase{
		{[][]int{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 1}, {0, 0, 0, 1, 1}}, 2},
		{[][]int{{0, 1, 0, 1, 0}, {1, 0, 1, 0, 1}, {0, 1, 1, 1, 0}, {1, 0, 1, 0, 1}}, 7},
		{[][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}, 1},
		{[][]int{}, 0},
		{[][]int{{}, {}}, 0},
		{[][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}, 0},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := countIslandsInPlace(testCase.Input)
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

func countIslandsInPlace(grid [][]int) int {
	result, x := 0, len(grid)
	if x < 1 {
		return result
	}
	y := len(grid[0])
	if y < 1 {
		return result
	}

	for i := range x {
		for j := range y {
			if grid[i][j] == 1 {
				markIslandInPlace(grid, i, j, x, y)
				result++
			}
		}
	}

	return result
}

func markIslandInPlace(grid [][]int, i, j, x, y int) {
	q := IndexPairQueue{}
	q.Enqueue(IndexPair{i, j})
	for !q.isEmpty() {
		top, _ := q.Dequeue()
		i, j := top.i, top.j
		if i >= 0 && i < x && j >= 0 && j < y && grid[i][j] == 1 {
			grid[i][j] = 0
			q.Enqueue(IndexPair{i - 1, j}, IndexPair{i, j + 1}, IndexPair{i + 1, j}, IndexPair{i, j - 1})
		}
	}
}

func countIslands(grid [][]int) int {
	result, x := 0, len(grid)
	if x < 1 {
		return result
	}
	seen := make([][]bool, x)
	y := len(grid[0])
	if y < 1 {
		return result
	}
	for i := range x {
		seen[i] = make([]bool, y)
	}

	for i := range x {
		for j := range y {
			if !seen[i][j] && grid[i][j] == 1 {
				markIsland(grid, i, j, x, y, seen)
				result++
			}
		}
	}

	return result
}

func markIsland(grid [][]int, i, j, x, y int, seen [][]bool) {
	q := IndexPairQueue{}
	q.Enqueue(IndexPair{i, j})
	for !q.isEmpty() {
		top, _ := q.Dequeue()
		i, j := top.i, top.j
		if i >= 0 && i < x && j >= 0 && j < y && !seen[i][j] && grid[i][j] == 1 {
			seen[i][j] = true
			q.Enqueue(IndexPair{i - 1, j}, IndexPair{i, j + 1}, IndexPair{i + 1, j}, IndexPair{i, j - 1})
		}
	}
}
