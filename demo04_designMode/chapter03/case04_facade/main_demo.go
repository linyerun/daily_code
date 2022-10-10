// 外观模式：

package main

type SystemA struct {
}

func (t *SystemA) A() {
	println("A系统方法A")
}

func (t *SystemA) B() {
	println("A系统方法B")
}

func (t *SystemA) C() {
	println("A系统方法C")
}

type SystemB struct {
}

func (t *SystemB) A() {
	println("B系统方法A")
}

func (t *SystemB) B() {
	println("B系统方法B")
}

type SystemC struct {
}

func (t *SystemC) A() {
	println("C系统方法A")
}

func (t *SystemC) B() {
	println("C系统方法B")
}

type facade struct {
	a *SystemA
	b *SystemB
	c *SystemC
}

func NewFacade(a *SystemA, b *SystemB, c *SystemC) *facade {
	return &facade{a: a, b: b, c: c}
}

func (t *facade) MethodOne() {
	t.a.A()
	t.b.A()
	t.c.A()
}

func (t *facade) MethodTwo() {
	t.a.B()
	t.b.B()
	t.c.B()
}

func (t *facade) MethodThree() {
	t.a.C()
}

func main() {
	facade := NewFacade(new(SystemA), new(SystemB), new(SystemC))
	facade.MethodOne()
	facade.MethodTwo()
	facade.MethodThree()
}
