//访问结构体标签
//访问类型方法

package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Person) GetName(author string) string {
	return author
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func main() {
	t0902()
}

func t0901() {
	p := new(Person)
	ty := reflect.TypeOf(p).Elem()
	//vl := reflect.ValueOf(p).Elem()
	for i := 0; i < ty.NumField(); i++ {
		println(ty.Field(i).Tag.Get("json"))
	}
}

func t0902() {
	print09(new(Person))
	println("------------------")
	print09(Person{})
}

func print09(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name, strings.TrimPrefix(methType.String(), "func"))
	}
}
