package main

import "fmt"

func main() {
	s := "11111"
	fmt.Printf("%s\n", s[:11]) //panic
}
