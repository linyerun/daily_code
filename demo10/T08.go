package main

import "sort"

func findNumberOfLIS(a []int) (ans int) {
	b := append(sort.IntSlice{}, a...)
	b.Sort()
	var bt BIT = make([]int, len(a)+1)
	var arr []int
	var maxVal int
	for _, x := range a {
		idx := b.Search(x + 1) //查一下x映射到的下标(关于+1, dddd)
		bt.add(idx)
		res := bt.query(idx)
		arr = append(arr, res)
		maxVal = max(maxVal, res)
	}
	for _, x := range arr {
		if x == maxVal {
			ans++
		}
	}
	return
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

type BIT []int

func (t BIT) query(x int) (res int) {
	for x > 0 {
		res += t[x]
		x &= x - 1
	}
	return
}

func (t BIT) add(x int) {
	for x < len(t) {
		t[x]++
		x += x & -x
	}
}
