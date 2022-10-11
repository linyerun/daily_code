package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64
	f = 12.3
	getNameValue(f)
}

func getNameValue(v any) {
	typeOf := reflect.TypeOf(v)
	valueOf := reflect.ValueOf(v)
	fmt.Printf("%v : %v \n", typeOf, valueOf)
}
