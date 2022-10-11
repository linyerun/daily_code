//reflect.ValueOf()的Kind()方法：类型分成几种：基础类型(int,bool...)、聚合类型(Array、struct)、引用类型(指针、Slice、Map、Slice...)、Invalid类型

package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func main() {
	var x int64 = 1
	var d = time.Nanosecond
	fmt.Println(Any(x))
	fmt.Println(Any(d))
	fmt.Println(Any([]int64{x}))
	fmt.Println(Any([]time.Duration{d}))
}

func Any(value any) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.Invalid:
		return "invalid"
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default: //reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
