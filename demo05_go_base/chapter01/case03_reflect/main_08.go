//修改值：在更新变量前使用CanAddr来检查并不能保证正确。
//		  CanSet方法才能正确报告一个reflect.Value是否可寻址且可更改。
//		  因为未导出字段不允许被修改

package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	//t0801()
	//t0802()
	//t0803()
	t0804()
}

func t0801() {
	x := 2
	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	d := c.Elem() //通过指针间接获取一个可寻址的reflect.Value

	fmt.Println("CanAddr:", a.CanAddr())
	fmt.Println("CanAddr:", b.CanAddr())
	fmt.Println("CanAddr:", c.CanAddr())
	fmt.Println("CanAddr:", d.CanAddr())
}

func t0802() {
	x := 10
	a := reflect.ValueOf(&x).Elem()
	a.SetInt(100) //使用Set相关的进行设置
	println("x:", x)
	a.Set(reflect.ValueOf(2)) //一定要是同类型的，不然就报错
	println("x:", x)
	ptr := a.Addr().Interface().(*int)
	*ptr = 3
	println("x:", x)
}

func t0803() {
	stdout := reflect.ValueOf(os.Stdout).Elem()
	fmt.Println(stdout.Type())
	//读取出之前不能读取出来的东西
	nameFd := stdout.FieldByName("name")
	println(nameFd.String())
	println(stdout.FieldByName("appendMode").Bool())
	fmt.Println(stdout.FieldByName("appendMode").CanSet()) //不可设置
}

func t0804() {
	a := new(AAA)
	println(reflect.ValueOf(a).Elem().CanAddr())
	println(reflect.ValueOf(a).Elem().FieldByName("a").CanSet())
	s := [...]int{1, 2, 3}
	println(reflect.ValueOf(&s[0]).Elem().CanAddr())
	is := []int{1, 2, 3}
	//println(reflect.ValueOf(is).Elem().CanAddr()) //报错了
	println(reflect.ValueOf(&is[0]).Elem().CanAddr())
}

type AAA struct {
	a string // 这个字段CanSet返回：false
}

func t0805() {
	//https://www.cnblogs.com/cheyunhua/p/15423910.html
	//细品吧，虽然知道自己不会回来看了
}
