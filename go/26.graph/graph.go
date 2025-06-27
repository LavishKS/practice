package main

import (
	"fmt"
	"strings"
)

type AdjacencyListGraph struct {
	list [][]int
}

func (g *AdjacencyListGraph) init(vertaxCount int) {
	g.list = make([][]int, vertaxCount)
	for i := range vertaxCount {
		g.list[i] = make([]int, 0)
	}
}

func (g *AdjacencyListGraph) addBiDirectionalEdge(v1, v2 int) {
	if v1 >= len(g.list) || v2 >= len(g.list) {
		panic("Invalid edge.")
	}
	g.list[v1] = append(g.list[v1], v2)
	g.list[v2] = append(g.list[v2], v1)
}

func (g *AdjacencyListGraph) addDirectionalEdge(v1, v2 int) {
	if v1 < 0 || v1 >= len(g.list) || v2 < 0 || v2 >= len(g.list) {
		panic("Invalid edge.")
	}
	g.list[v1] = append(g.list[v1], v2)
}

func (g AdjacencyListGraph) getBFSOrder(start int) []int {
	q := Queue{start}
	result := []int{}
	seen := make(map[int]bool)

	for !q.isEmpty() {
		top, _ := q.Dequeue()
		seen[top] = true
		result = append(result, top)
		for _, ch := range g.list[top] {
			if !seen[ch] {
				q.Enqueue(ch)
			}
		}
	}

	return result
}

func (g AdjacencyListGraph) getDFSOrder(start int) []int {
	seen := make(map[int]bool)
	var dfs func(node int, result []int) []int
	dfs = func(node int, result []int) []int {
		if node >= 0 && node < len(g.list) {
			seen[node] = true
			result = append(result, node)
			for _, ch := range g.list[node] {
				if !seen[ch] {
					result = dfs(ch, result)
				}
			}
		}
		return result
	}
	return dfs(start, []int{})
}

func (g AdjacencyListGraph) String() string {
	var sb strings.Builder
	sb.WriteString("List Graph:\n[\n")
	for i, v := range g.list {
		sb.WriteString(fmt.Sprintf("%d: %v\n", i, v))
	}
	sb.WriteString("]")
	return sb.String()
}

type AdjacencyMatrixGraph struct {
	matrix [][]uint8
}

func (g *AdjacencyMatrixGraph) init(vertaxCount int) {
	g.matrix = make([][]uint8, vertaxCount)
	for i := range vertaxCount {
		g.matrix[i] = make([]uint8, vertaxCount)
	}
}

func (g *AdjacencyMatrixGraph) addBiDirectionalEdge(v1, v2 int) {
	if v1 >= len(g.matrix) || v2 >= len(g.matrix) {
		panic("Invalid edge.")
	}

	g.matrix[v1][v2] = 1
	g.matrix[v2][v1] = 1
}

func (g AdjacencyMatrixGraph) String() string {
	var sb strings.Builder
	sb.WriteString("Matrix Graph:\n[\n")
	for _, row := range g.matrix {
		sb.WriteString("[ ")
		for _, v := range row {
			sb.WriteString(fmt.Sprintf("%d, ", v))
		}
		sb.WriteString("]\n")
	}
	sb.WriteString("]")
	return sb.String()
}

func fromListToMatrix(lg AdjacencyListGraph) AdjacencyMatrixGraph {
	mg := AdjacencyMatrixGraph{}
	mg.init(len(lg.list))
	for v1, l := range lg.list {
		for _, v2 := range l {
			mg.addBiDirectionalEdge(v1, v2)
		}
	}

	return mg
}

func graphIntroMain() {
	graph := AdjacencyListGraph{}
	graph.init(6)
	graph.addBiDirectionalEdge(0, 3)
	graph.addBiDirectionalEdge(1, 3)
	graph.addBiDirectionalEdge(2, 3)
	graph.addBiDirectionalEdge(4, 3)
	graph.addBiDirectionalEdge(4, 5)
	fmt.Println(graph)

	mgraph := fromListToMatrix(graph)
	fmt.Println(mgraph)
}

func graphBFSDFSMain() {
	graph := AdjacencyListGraph{}
	graph.init(9)
	graph.addBiDirectionalEdge(0, 1)
	graph.addBiDirectionalEdge(0, 3)
	graph.addBiDirectionalEdge(4, 3)
	graph.addBiDirectionalEdge(5, 3)
	graph.addBiDirectionalEdge(2, 3)
	graph.addBiDirectionalEdge(2, 8)
	graph.addBiDirectionalEdge(4, 6)
	graph.addBiDirectionalEdge(7, 6)

	fmt.Println(graph)

	fmt.Println("BFS:", graph.getBFSOrder(0))
	fmt.Println("DFS:", graph.getDFSOrder(0))
}
