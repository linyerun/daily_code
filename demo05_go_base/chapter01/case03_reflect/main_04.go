// reflect.ValueOf的逆操作Interface()：返回一个interface{}接口值，与reflect.Value包含同一个具体值

package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf(3)
	x := v.Interface()
	switch x := x.(type) {
	case int:
		fmt.Println(x)
	}
}
