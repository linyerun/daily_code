package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("请分别输入A、B、C：")
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if !ok(a) || !ok(b) || !ok(c) {
		fmt.Println("无效值")
		return
	}
	if a == 0 {
		if b == 0 {
			if c == 0 {
				fmt.Println("有无数根")
				return
			}
			fmt.Println("无根")
			return
		}
		fmt.Println("有一个实根")
		return
	}
	t := b*b - 4*a*c
	if t > 0 {
		fmt.Println("有两个不相等的实根")
	} else if t == 0 {
		fmt.Println("有两个相等的实根")
	} else {
		fmt.Println("有两个虚根")
	}
}

func ok(num int) bool {
	return num >= -200 && num <= 200
}
