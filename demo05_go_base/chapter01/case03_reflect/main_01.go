// 当我们无法透视一个未知类型的布局时，我们就需要反射了

package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := print(1.2)
	println(s)
}

//简单实现一个类型字符串化
func print(v any) string {
	type str interface {
		String() string
	}
	//类型断言第二种
	switch v := v.(type) {
	case string:
		return v
	case str:
		return v.String()
	case int:
		return strconv.Itoa(v)
	case bool:
		if v {
			return "true"
		}
		return "false"
	case float64:
		return fmt.Sprintf("%v", v)
	default:
		return "???"
	}
}
