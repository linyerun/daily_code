package main

type A struct {
	*A
}

func main() {
	ff := new(f)
	ff.Func(100)
	//t := new(TestF)	//确实调用不了
	new(HaHa).Func(100) //特例：通过嵌套可以调用
}

type f func(int)

func (t *f) Func(num int) {
	(*t)(num)
}

type TestF f

type HaHa struct {
	f
}
