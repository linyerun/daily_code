package main

import (
	"fmt"
	"unicode/utf8"
)

//做了:=的测试，还有函数DecodeRuneInString的测试！

func main() {
	const name = "林叶润，加油！"
	for i, w := 0, 0; i < len(name); i += w {
		//第一个Rune的占字节数，值
		var val rune
		val, w = utf8.DecodeRuneInString(name[i:])
		fmt.Println("值：", string(val))
	}
	w := 0
	w, a := 1, 1
	println(w, a)
}
