//简单工厂模式：还是存在依赖，不完美，不符合开闭原则

package main

//抽象层

type Fruit interface {
	Show()
}

//实现层

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

// SimplyFactory 工厂
type SimplyFactory struct {
}

func (t *SimplyFactory) CreateFruit(name string) (fruit Fruit) {
	switch name {
	case "apple":
		fruit = new(apple)
	case "banana":
		fruit = new(banana)
	case "pear":
		fruit = new(pear)
	}
	return
}

//业务逻辑层

func main() {
	sf := new(SimplyFactory)
	sf.CreateFruit("apple").Show()
	sf.CreateFruit("banana").Show()
	sf.CreateFruit("pear").Show()
}
