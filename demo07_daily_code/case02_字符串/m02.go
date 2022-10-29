package main

import "fmt"

func main() {
	s := "abc"
	t := []byte(s)
	for i := 0; i < len(t); i++ {
		t[i] += 1
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d]=%c,t[%d]=%c\n", i, s[i], i, t[i])
	}
}
