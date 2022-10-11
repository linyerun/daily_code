// 需要配合"const()"使用

package main

const (
	// iota初始值为0，然后不断+1，碰到公式按公式走，出现不同公式按照下一条公式开始
	a = iota
	b
	c
	d = iota * 10
	e
	f
)

func main() {
	println("a:", a)
	println("b:", b)
	println("c:", c)
	println("d:", d)
	println("e:", e)
	println("f:", f)
}
