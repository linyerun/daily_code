//依赖倒置原则：通过下面的例子去感受，我说不清楚。两层都面向抽象层，不直接依赖，真牛

package main

import "fmt"

//抽象层

type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

//实现层：向上依赖抽象层

type ZhangSan struct{}

func (t *ZhangSan) Drive(car Car) {
	fmt.Print("张三：")
	car.Run()
}

type LiSi struct{}

func (t *LiSi) Drive(car Car) {
	fmt.Print("李四：")
	car.Run()
}

type Benz struct{}

func (t *Benz) Run() {
	fmt.Println("启动奔驰车")
}

type BWN struct{}

func (t *BWN) Run() {
	fmt.Println("启动宝马车")
}

func main() {
	//业务逻辑层：向下依赖抽象层
	new(ZhangSan).Drive(new(Benz)) //张三开奔驰
	new(LiSi).Drive(new(BWN))      //李四开宝马
}
