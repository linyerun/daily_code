package main

import "container/heap"

func nthUglyNumber(n int) int {
	var h = &hp{}
	//var h *hp	//这样写不行，因为我们没有给地址开辟一个 4 字节的空间
	heap.Push(h, 1)
	set := make(map[int]struct{})
	set[1] = struct{}{}
	var ans int
	a := []int{2, 3, 5}
	for i := 0; i < n; i++ {
		ans = heap.Pop(h).(int)
		for j := 0; j < len(a); j++ {
			if _, ok := set[ans*a[j]]; !ok {
				set[ans*a[j]] = struct{}{}
				heap.Push(h, ans*a[j])
			}
		}
	}
	return ans
}

type hp []int

func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *hp) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *hp) Pop() interface{} {
	ret := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return ret
}
