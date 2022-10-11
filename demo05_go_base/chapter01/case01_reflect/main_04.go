//反射获取结构体的tag

package main

import (
	"fmt"
	"reflect"
)

type Element struct {
	Name string `json:"name" doc:"名字"`
	Age  int    `json:"age" doc:"年龄"`
}

func main() {
	// 传指针才能获取到
	printTag(&Element{Name: "aaa", Age: 20})
}

func printTag(str any) {
	e := reflect.TypeOf(str).Elem()
	for i := 0; i < e.NumField(); i++ {
		getJSON := e.Field(i).Tag.Get("json")
		getDOC := e.Field(i).Tag.Get("doc")
		fmt.Printf("json=%s;doc=%s\n", getJSON, getDOC)
	}
}
