package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	//io.Writer是接口类型，它由两部分组成：动态类型、动态值
	//os.Stdout是*os.File类型
	var w io.Writer = os.Stdout
	t := reflect.TypeOf(w) //获取w的动态类型
	println(t.String())
	fmt.Println(t)        //本质上底层调用的也是t.String()，但是println(t)底层就不是调用String()方法了
	fmt.Printf("%T\n", w) //打印动态类型
}
