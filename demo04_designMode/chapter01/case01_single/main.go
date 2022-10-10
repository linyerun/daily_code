//单一职责原则: 类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个。

package main

import "fmt"

type ClothesShop struct{}

func (t *ClothesShop) OnShop() {
	fmt.Println("休闲的装扮")
}

type ClothesWork struct{}

func (t *ClothesWork) OnWork() {
	fmt.Println("工作的装扮")
}

func main() {
	//case1
	cs := new(ClothesShop)
	cs.OnShop()
	//case2
	cw := new(ClothesWork)
	cw.OnWork()
}
