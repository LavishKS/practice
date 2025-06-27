package main

import (
	"fmt"
	"math"
)

type WallsAndGatesTestCase struct {
	Input    [][]int
	Expected [][]int
}

func (t WallsAndGatesTestCase) String() string {
	return fmt.Sprintf("\tInput    : %v\n\tExpected : %d", t.Input, t.Expected)
}

func (t WallsAndGatesTestCase) validateResult(result [][]int) bool {
	if len(t.Expected) != len(result) {
		return false
	}

	for i, r1 := range t.Expected {
		r2 := result[i]
		if len(r1) != len(r2) {
			return false
		}
		for j, v1 := range r1 {
			v2 := r2[j]
			if v1 != v2 {
				return false
			}
		}
	}

	return true
}

const (
	WALL       int = -1
	GATE       int = 0
	EMPTY_ROOM int = math.MaxInt32
)

func testWallsAndGatesMain() {
	testCases := []WallsAndGatesTestCase{
		{[][]int{{EMPTY_ROOM, WALL, GATE, EMPTY_ROOM}, {EMPTY_ROOM, EMPTY_ROOM, EMPTY_ROOM, WALL}, {EMPTY_ROOM, WALL, EMPTY_ROOM, WALL}, {GATE, WALL, EMPTY_ROOM, EMPTY_ROOM}}, [][]int{{3, WALL, GATE, 1}, {2, 2, 1, WALL}, {1, WALL, 2, WALL}, {GATE, WALL, 3, 4}}},
		{[][]int{{EMPTY_ROOM, WALL, GATE, EMPTY_ROOM}, {WALL, EMPTY_ROOM, EMPTY_ROOM, WALL}, {EMPTY_ROOM, WALL, EMPTY_ROOM, WALL}, {GATE, WALL, EMPTY_ROOM, EMPTY_ROOM}}, [][]int{{EMPTY_ROOM, WALL, GATE, 1}, {-1, 2, 1, WALL}, {1, WALL, 2, WALL}, {GATE, WALL, 3, 4}}},
		{[][]int{}, [][]int{}},
		{[][]int{{}, {}}, [][]int{{}, {}}},
	}

	passed := 0

	for i, testCase := range testCases {
		fmt.Println("Testcase:", i+1)
		fmt.Println(testCase)
		result := fillRoomStepsWithDFS(testCase.Input)
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


func traverseGrid(grid [][]int, step, i, j, x, y int) {
	if i < 0 || i == x || j < 0 || j == y || step > grid[i][j] {
		return
	}
	grid[i][j] = step
	step++
	traverseGrid(grid, step, i-1, j, x, y)
	traverseGrid(grid, step, i, j+1, x, y)
	traverseGrid(grid, step, i+1, j, x, y)
	traverseGrid(grid, step, i, j-1, x, y)
}

var dirs = [4]IndexPair{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func traverseGridV2(grid [][]int, step, i, j, x, y int) {
	grid[i][j] = step
	step++
	for _, d := range dirs {
		row, col := i+d.i, j+d.j
		if row >= 0 && row < x && col >= 0 && col < y && grid[row][col] > step {
			traverseGrid(grid, step, row, col, x, y)
		}
	}
}

func fillRoomStepsWithDFS(grid [][]int) [][]int {
	x := len(grid)
	if x < 1 {
		return grid
	}
	y := len(grid[0])
	if y < 1 {
		return grid
	}

	for i, row := range grid {
		for j, v := range row {
			if v == GATE {
				traverseGrid(grid, 0, i, j, x, y)
			}
		}
	}
	return grid
}

func fillRoomSteps(grid [][]int) [][]int {
	x := len(grid)
	if x < 1 {
		return grid
	}
	y := len(grid[0])
	if y < 1 {
		return grid
	}

	q := IndexPairQueue{}
	for i, row := range grid {
		for j, v := range row {
			if v == GATE {
				q.Enqueue(IndexPair{i, j})
			}
		}
	}

	for !q.isEmpty() {
		nextQ := IndexPairQueue{}

		for _, pair := range q {
			i, j := pair.i, pair.j
			dist := grid[i][j] + 1
			for _, d := range dirs {
				row, col := i+d.i, j+d.j
				if row >= 0 && row < x && col >= 0 && col < y && grid[row][col] == EMPTY_ROOM {
					grid[row][col] = dist
					nextQ.Enqueue(IndexPair{row, col})
				}
			}
		}

		q = nextQ
	}
	return grid
}
