package main

type AAA interface {
	A()
}

type BBB interface {
	AAA
	B()
}

// CCC TODO 接口中嵌入结构体
type CCC interface {
	HaHa
	TTT
}

type HaHa struct {
}

type TTT struct {
}

func (t *TTT) A() {

}

func (t *TTT) B() {

}

type LLL struct {
	AAA
	BBB
}

func main() {
	t := new(T1)
	var i I1 = t
	var ii = &i
	i.B()
	i.A()
	(*ii).B()
	t.I1.A()
	t.I1.B()
}

//接口嵌入接口
//接口嵌入结构体
//结构体嵌入结构体	(OK)
//结构体嵌入接口(不可以放入接口的指针)	(OK)

type T1 struct {
	I1
}

type I1 interface {
	B()
	A()
}
