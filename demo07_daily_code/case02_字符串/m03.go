package main

import "fmt"

func main() {
	const s = "我是中国人"
	for idx, elem := range s {
		fmt.Println("起始下标:", idx, ";元素:", string(elem))
	}
}
