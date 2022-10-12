package main

import "fmt"

func main() {
	s := []byte("哈哈\n")
	fmt.Println(string(s[:len(s)-2]))
	fmt.Println("xxxxxxx")
}
