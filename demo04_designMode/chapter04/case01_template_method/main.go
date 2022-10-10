// 模板方法模式：父类定好抽象方法执行流程，子类负责实现模板方法

package main

//饮料制作流程

type Beverage interface {
	BoilWater()         //煮开水
	Brew()              //冲泡
	PourInCup()         //倒进杯中
	AddThings()         //添加佐料
	WantAddThing() bool //是否添加佐料
}

type Template struct {
	Beverage
}

func (t *Template) MakeBeverage() {
	if t.Beverage == nil {
		return
	}
	t.BoilWater()
	t.Brew()
	t.PourInCup()
	if t.WantAddThing() {
		t.AddThings()
	}
}

//制作咖啡

type makeCoffee struct {
	Template
	wantAddThings bool
}

func NewMakeCoffee(wantAddThings bool) *makeCoffee {
	m := new(makeCoffee)
	m.Beverage = m
	m.wantAddThings = wantAddThings
	return m
}

func (t *makeCoffee) BoilWater() {
	println("将水煮到100摄氏度")
}

func (t *makeCoffee) Brew() {
	println("开始冲泡咖啡")
}

func (t *makeCoffee) PourInCup() {
	println("将咖啡倒入杯中")
}

func (t *makeCoffee) AddThings() {
	println("在咖啡中加入牛奶和白糖")
}

// WantAddThing 这里一开始写错了，导致运行有误
func (t *makeCoffee) WantAddThing() bool {
	return t.wantAddThings
}

func main() {
	coffee := NewMakeCoffee(true)
	coffee.MakeBeverage()
}
