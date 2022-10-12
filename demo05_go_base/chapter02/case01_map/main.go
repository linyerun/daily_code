//测试map的value为指针的时候的更新操作,没问题
//map无法调用cap内置函数

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) String() string {
	return fmt.Sprintf("User{name=%s,Age=%d}", u.Name, u.Age)
}

func main() {
	var m = make(map[string]*User)
	u := &User{Name: "aa", Age: 20}
	m["aa"] = u
	for _, user := range m {
		fmt.Println(user)
	}
	u.Name = "bb"
	u.Age = 21
	for _, user := range m {
		fmt.Println(user)
	}
}
