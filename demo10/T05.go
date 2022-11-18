package main

//type RecentCounter struct {
//	h []int
//}
//
//func Constructor() RecentCounter {
//	return RecentCounter{}
//}
//
//func (r *RecentCounter) Ping(t int) int {
//	val := t - 3000
//	for r.Len() != 0 && r.h[0] < val {
//		heap.Pop(r)
//	}
//	heap.Push(r, t)
//	return r.Len()
//}
//
//func (r *RecentCounter) Len() int {
//	return len(r.h)
//}
//
//func (r *RecentCounter) Less(i, j int) bool {
//	return r.h[i] < r.h[j]
//}
//
//func (r *RecentCounter) Swap(i, j int) {
//	r.h[i], r.h[j] = r.h[j], r.h[i]
//}
//
//func (r *RecentCounter) Push(x interface{}) {
//	r.h = append(r.h, x.(int))
//}
//
//func (r *RecentCounter) Pop() interface{} {
//	defer func() { r.h = r.h[:r.Len()-1] }()
//	return r.h[r.Len()-1]
//}

func main() {

}
