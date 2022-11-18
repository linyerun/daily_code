package main

type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (r *RecentCounter) Ping(t int) int {
	val := t - 3000
	for len(r.q) != 0 && r.q[0] < val {
		r.q = r.q[1:]
	}
	r.q = append(r.q, t)
	return len(r.q)
}

func main() {

}
