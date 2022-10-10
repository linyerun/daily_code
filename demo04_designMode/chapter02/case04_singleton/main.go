// 单例模式：饿汉式、懒汉式，下面展示懒汉式

package main

import "sync"

type singleton struct {
}

var s *singleton
var once sync.Once

func NewSingleton() *singleton {
	once.Do(func() {
		s = new(singleton)
	})
	return s
}

func main() {
	s1 := NewSingleton()
	s2 := NewSingleton()
	if s1 == s2 {
		println("s1 == s2")
	} else {
		println("s1 != s2")
	}
}
