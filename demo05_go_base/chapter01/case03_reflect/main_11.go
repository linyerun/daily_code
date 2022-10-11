//调用对象方法测试

package main

import "reflect"

type Hello struct {
}

func (h Hello) SayHelloToOthers(name string) {
	println("Hello", name)
}

/**
	1. reflect.Type的Method(i)返回reflect.Method类型实例，这个结构类型描述了这个方法的名称和类型
	2. reflect.Value的Method(i)返回一个reflect.Value，代表一个方法值(即一个已绑定接收者的方法)，
       使用Call方法可以调用Func类型的Value
*/

func main() {
	v := reflect.ValueOf(new(Hello))
	v.MethodByName("SayHelloToOthers").Call([]reflect.Value{reflect.ValueOf("Ernie")})
}
