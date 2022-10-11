//reflect.Value也可以包含一个接口值

//reflect.Value也满足fmt.Stringer，但是除非Value包含的是一个字符串，否则String方法仅仅暴露类型

package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	// ---- 测试2 ----
	v = reflect.ValueOf("123")
	fmt.Println(v)
	fmt.Println(v.String())
}
