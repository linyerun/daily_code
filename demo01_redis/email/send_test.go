package email

import "testing"

func TestSend(t *testing.T) {
	if err := Send("xxx", []byte("123"), "linyerundgut@126.com"); err != nil {
		panic(err)
	}
}
