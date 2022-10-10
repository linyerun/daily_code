// golang实现这个模板方法不够优雅

package main

type Drinking interface {
	BoilWater()         //煮开水
	Brew()              //冲泡
	PourInCup()         //倒进杯中
	AddThings()         //添加佐料
	WantAddThing() bool //是否添加佐料
}

type makeDrinkingTemplate struct {
	d Drinking
}

func NewDrinking(drinking Drinking) *makeDrinkingTemplate {
	return &makeDrinkingTemplate{d: drinking}
}

func (t *makeDrinkingTemplate) Make() {
	t.d.BoilWater()
	t.d.Brew()
	t.d.PourInCup()
	if t.d.WantAddThing() {
		t.d.AddThings()
	}
}

type tea struct {
	wantAddThing bool
}

func NewTea(wantAddThing bool) *tea {
	return &tea{wantAddThing: wantAddThing}
}

func (t *tea) BoilWater() {
	println("将水煮到100摄氏度")
}

func (t *tea) Brew() {
	println("开始冲泡茶叶")
}

func (t *tea) PourInCup() {
	println("将茶叶倒入杯中")
}

func (t *tea) AddThings() {
	println("在茶叶中蜂蜜")
}

func (t *tea) WantAddThing() bool {
	return t.wantAddThing
}

func main() {
	NewDrinking(NewTea(false)).Make()
	NewDrinking(NewTea(true)).Make()
}
