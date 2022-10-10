// 工厂方法模式：简单工厂模式+开闭原则
// 我的理解：工厂方法模式是将之前需要在业务逻辑层创建对象的大量逻辑抽离出来实现复用

package main

//抽象层

type Fruit interface {
	Show()
}

type FruitFactory interface {
	CreateFruit() Fruit
}

//实现层

//水果实现层

type apple struct {
}

func (t *apple) Show() {
	println("我是苹果")
}

type banana struct {
}

func (t *banana) Show() {
	println("我是香蕉")
}

type pear struct {
}

func (t pear) Show() {
	println("我是鸭梨")
}

//工厂实现层

type AppleFactory struct {
}

func (t *AppleFactory) CreateFruit() Fruit {
	return new(apple)
}

type BananaFactory struct {
}

func (t *BananaFactory) CreateFruit() Fruit {
	return new(banana)
}

type PearFactory struct {
}

func (t *PearFactory) CreateFruit() Fruit {
	return new(pear)
}

//业务逻辑层
func main() {
	new(AppleFactory).CreateFruit().Show()
	new(BananaFactory).CreateFruit().Show()
	new(PearFactory).CreateFruit().Show()
}
