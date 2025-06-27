package main

import "fmt"

type Result struct {
	order []int
	size  uint
	index uint
}

func makeResult(size uint) Result {
	r := Result{[]int{}, 0, 0}
	if size > 0 {
		r.order = make([]int, size)
		r.size = size
	}
	return r

}

func (r *Result) push(val int) {
	if r.size <= r.index {
		return
	}
	r.order[r.index] = val
	r.index++
}

func (r Result) String() string {
	return fmt.Sprintf("Result: %v", r.order)
}

func calcId(i, j, y int) int {
	return i*y + j
}

func dfs(input [][]int, i, j, x, y int, seen map[int]bool, r Result) {
	id := calcId(i, j, y)
	if i < 0 || i >= x || j < 0 || j >= y || seen[id] {
		return
	}

	seen[id] = true
	r.push(input[i][j])

	dfs(input, i-1, j, x, y, seen, r)
	dfs(input, i, j+1, x, y, seen, r)
	dfs(input, i+1, j, x, y, seen, r)
	dfs(input, i, j-1, x, y, seen, r)
}

func performDFS(twoDArr [][]int) Result {
	x, y := len(twoDArr), len(twoDArr[0])
	result := makeResult(uint(x * y))
	seen := make(map[int]bool, 0)
	dfs(twoDArr, 0, 0, x, y, seen, result)
	return result
}

func twoDimArrayMain() {
	twoDArr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}
	// result := performDFS(twoDArr)
	result := performBFS(twoDArr)
	fmt.Println(result)
}

// func bfs(twoDArr [][]int, x, y int, seen map[int]map[int]bool, result Result) {
// 	q := IndexPairQueue{{0, 0}}
// 	for !q.isEmpty() {
// 		top, _ := q.Dequeue()
// 		i, j := top.i, top.j
// 		if i < 0 || i >= x || j < 0 || j >= y || seen[i][j] {
// 			continue
// 		}
// 		seen[i][j] = true
// 		result.push(twoDArr[i][j])
// 		q.Enqueue(IndexPair{i - 1, j}, IndexPair{i, j + 1}, IndexPair{i + 1, j}, IndexPair{i, j - 1})
// 	}
// }

func performBFS(twoDArr [][]int) Result {
	x, y := len(twoDArr), len(twoDArr[0])
	result := makeResult(uint(x * y))
	seen := make([][]bool, x)
	for i := range x {
		seen[i] = make([]bool, y)
	}

	q := IndexPairQueue{{2, 2}}
	for !q.isEmpty() {
		top, _ := q.Dequeue()
		i, j := top.i, top.j
		if i >= 0 && i < x && j >= 0 && j < y && !seen[i][j] {
			seen[i][j] = true
			result.push(twoDArr[i][j])
			q.Enqueue(IndexPair{i - 1, j}, IndexPair{i, j + 1}, IndexPair{i + 1, j}, IndexPair{i, j - 1})
		}
	}
	return result
}
