package main

import (
	"fmt"
)

type RottingOrangeTestCase struct {
	Input    [][]int
	Expected int
}

func (t RottingOrangeTestCase) String() string {
	return fmt.Sprintf("\tInput    : %v\n\tExpected : %d", t.Input, t.Expected)
}

func (t RottingOrangeTestCase) validateResult(result int) bool {
	return result == t.Expected
}

func testRottingOrangeMain() {
	testCases := []RottingOrangeTestCase{
		{[][]int{{2, 1, 1, 0, 0}, {1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 1, 0, 0, 1}}, 7},
		{[][]int{{2, 1, 1, 0, 0}, {1, 1, 1, 0, 2}, {0, 1, 1, 1, 1}, {0, 1, 0, 0, 1}}, 4},
		{[][]int{{1, 1, 0, 0, 0}, {2, 1, 0, 0, 0}, {0, 0, 0, 1, 2}, {0, 1, 0, 0, 1}}, -1},
		{[][]int{{2, 0, 0, 0, 0, 0, 0}, {1, 0, 1, 1, 1, 1, 0}, {1, 0, 1, 0, 0, 1, 0}, {1, 0, 1, 1, 0, 1, 0}, {1, 0, 0, 0, 0, 1, 0}, {1, 1, 1, 1, 1, 1, 0}, {0, 0, 0, 0, 0, 0, 0}}, 20},
		{[][]int{}, 0},
		{[][]int{{}, {}}, 0},
		{[][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}, 0},
		{[][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}, -1},
		{[][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}}, 0},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := countMinutesTillAllOrangesRotV2(testCase.Input)
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

func countMinutesTillAllOrangesRotV2(grid [][]int) int {
	result, x := 0, len(grid)
	if x < 1 {
		return result
	}
	y := len(grid[0])
	if y < 1 {
		return result
	}

	freshOranges, rottenOranges := 0, IndexPairQueue{}

	for i, row := range grid {
		for j, v := range row {
			switch v {
			case 1:
				freshOranges++
			case 2:
				rottenOranges.Enqueue(IndexPair{i, j})
			}
		}
	}

	dirs := []IndexPair{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for freshOranges > 0 {
		rottingOranges := IndexPairQueue{}
		for _, rottenOrangeIndex := range rottenOranges {
			i, j := rottenOrangeIndex.i, rottenOrangeIndex.j
			for _, dir := range dirs {
				row, col := i+dir.i, j+dir.j
				if row >= 0 && row < x && col >= 0 && col < y && grid[row][col] == 1 {
					grid[row][col] = 2
					freshOranges--
					rottingOranges.Enqueue(IndexPair{row, col})
				}
			}
		}
		if rottingOranges.isEmpty() {
			break
		}
		result++
		rottenOranges = rottingOranges
	}

	if freshOranges > 0 {
		return -1
	}

	return result
}

func countMinutesTillAllOrangesRot(grid [][]int) int {
	result, x := 0, len(grid)
	if x < 1 {
		return result
	}
	y := len(grid[0])
	if y < 1 {
		return result
	}

	freshOranges, rottenOranges := map[IndexPair]struct{}{}, IndexPairQueue{}

	for i, row := range grid {
		for j, v := range row {
			switch v {
			case 1:
				freshOranges[IndexPair{i, j}] = struct{}{}
			case 2:
				rottenOranges.Enqueue(IndexPair{i, j})
			}
		}
	}

	dirs := []IndexPair{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for {
		rottingOranges := IndexPairQueue{}
		for !rottenOranges.isEmpty() && len(freshOranges) > 0 {
			top, _ := rottenOranges.Dequeue()
			i, j := top.i, top.j
			for _, dir := range dirs {
				indexPair := IndexPair{i + dir.i, j + dir.j}
				if _, exists := freshOranges[indexPair]; exists {
					delete(freshOranges, indexPair)
					rottingOranges.Enqueue(indexPair)
				}
			}
		}
		if rottingOranges.isEmpty() {
			break
		}
		result++
		rottenOranges = rottingOranges
	}

	if len(freshOranges) > 0 {
		return -1
	}

	return result
}
