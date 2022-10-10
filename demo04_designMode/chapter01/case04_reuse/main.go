//合成复用：能使用组合实现的，就别使用继承

package main

import "fmt"

//小猫：能吃

type Cat struct {
}

func (t *Cat) Eat() {
	fmt.Println("小猫能吃")
}

//我想拓展小猫的功能
//1. 使用继承

type CatLyr struct {
	Cat
}

func (t *CatLyr) Sleep() {
	fmt.Println("小猫能睡")
}

//2. 使用组合

type CatErnie struct {
	c *Cat
}

func (t *CatErnie) Eat() {
	t.c.Eat()
}

func (t *CatErnie) Sleep() {
	fmt.Println("小猫能睡")
}

//3. 更加好的组合

type CatHaHa struct {
}

func (t *CatHaHa) Eat(c *Cat) {
	c.Eat()
}

func (t *CatHaHa) Sleep() {
	fmt.Println("小猫能睡")
}

//演示
func main() {
	//1. 原始的
	new(Cat).Eat()

	fmt.Println("-----------------------")
	catLyr := new(CatLyr)
	catLyr.Eat()
	catLyr.Sleep()

	fmt.Println("-----------------------")
	catErnie := &CatErnie{&Cat{}}
	catErnie.Eat()
	catErnie.Sleep()

	fmt.Println("-----------------------")
	catHaHa := new(CatHaHa)
	catHaHa.Eat(new(Cat))
	catHaHa.Sleep()
}
