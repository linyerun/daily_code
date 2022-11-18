package main

import "container/heap"

func minCostConnectPoints(ps [][]int) int {
	var minHeap *MinHeap
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			heap.Push(
				minHeap,
				Pair{x: i, y: j, val: abs(ps[i][0], ps[j][0]) + abs(ps[i][1], ps[j][1])},
			)
		}
	}
	return 0
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

type Pair struct {
	x, y int
	val  int
}

type MinHeap []Pair

func (m *MinHeap) Len() int {
	return len(*m)
}

func (m *MinHeap) Less(i, j int) bool {
	return (*m)[i].val < (*m)[j].val
}

func (m *MinHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MinHeap) Push(x any) {
	*m = append(*m, x.(Pair))
}

func (m *MinHeap) Pop() any {
	defer func() { *m = (*m)[:m.Len()-1] }()
	return (*m)[m.Len()-1]
}
