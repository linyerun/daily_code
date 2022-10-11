package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func NewPerson(id int, name string, age int) *Person {
	return &Person{Id: id, Name: name, Age: age}
}

func (p Person) Call() {
	println("我是person对象的方法")
}

func main() {
	getPersonFieldMethod(*NewPerson(1, "哈哈", 21))
}

func getPersonFieldMethod(v any) {
	typeOf := reflect.TypeOf(v)
	valueOf := reflect.ValueOf(v)
	fmt.Printf("%v : %v\n", typeOf, valueOf)

	//获取结构体(实体)里面的字段
	for i := 0; i < typeOf.NumField(); i++ {
		t := typeOf.Field(i)
		v := valueOf.Field(i).Interface()
		fmt.Printf("%s=%v : %v\n", t.Name, t.Type, v)
	}

	//获取结构体方法
	for i := 0; i < typeOf.NumMethod(); i++ {
		m := typeOf.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

}
