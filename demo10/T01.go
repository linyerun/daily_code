package main

type AAA struct {
	name string
	age  int
}

type BBB struct {
	name string
	age  string
}

func main() {
	//golang的duck模型通过"接口+结构体"来体现
	//var a AAA  = new(BBB)
}
