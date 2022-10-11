package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name     string `json:"name"`
	Director string `json:"director"`
}

func main() {
	m1 := new(Movie)
	m1.Name = "功夫"
	m1.Director = "周星驰"
	//JSON化
	marshal, err := json.Marshal(m1)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%v\n", string(marshal))
	//JSON解析
	var m2 Movie
	json.Unmarshal(marshal, &m2)
	fmt.Printf("%v\n", m2)
}
