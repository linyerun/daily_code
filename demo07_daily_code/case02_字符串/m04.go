package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const name = "林叶润"
	//获取字符个数的函数
	cntName := utf8.RuneCountInString(name)
	fmt.Println("名字里面有多少个字：", cntName)
}
