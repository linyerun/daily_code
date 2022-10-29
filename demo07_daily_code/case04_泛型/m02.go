package main

import (
	"fmt"
	"reflect"
)

func GetAllKeys[K comparable, V any](m map[K]V) (ans []K) {
	for key := range m {
		ans = append(ans, key)
	}
	return
}

func main() {
	m := map[int]string{
		1: "啊啊",
		2: "噢噢",
		3: "呃呃",
	}
	//当调用 MapKeys 的时候， 我们不需要为 K 和 V 指定类型 - 编译器会进行自动推断
	var keys01 = GetAllKeys(m)
	fmt.Println("Keys01:", keys01, "type:", reflect.TypeOf(keys01).String())
	keys02 := GetAllKeys[int, string](m)
	fmt.Println("Keys02:", keys02, "type:", reflect.TypeOf(keys02).String())
}
